package libsq

import (
	"context"
	"fmt"

	"github.com/neilotoole/sq/libsq/core/tablefq"

	"github.com/samber/lo"

	"github.com/neilotoole/sq/libsq/source"

	"github.com/neilotoole/sq/libsq/core/record"

	"github.com/neilotoole/sq/libsq/core/options"

	"github.com/neilotoole/sq/libsq/ast/render"

	"github.com/neilotoole/sq/libsq/core/lg"

	"github.com/neilotoole/sq/libsq/core/lg/lga"

	"github.com/neilotoole/sq/libsq/ast"
	"github.com/neilotoole/sq/libsq/core/errz"
	"github.com/neilotoole/sq/libsq/core/sqlmodel"
	"github.com/neilotoole/sq/libsq/core/sqlz"
	"github.com/neilotoole/sq/libsq/driver"
	"golang.org/x/sync/errgroup"
)

// pipeline is used to execute a SLQ query,
// and write the resulting records to a RecordWriter.
type pipeline struct {
	// query is the SLQ query
	query string

	// qc is the context in which the query is executed.
	qc *QueryContext

	// rc is the Context for rendering SQL.
	// This field is set during pipeline.prepare. It can't be set before
	// then because the target DB to use is calculated during pipeline.prepare,
	// based on the input query and other context.
	rc *render.Context

	// tasks contains tasks that must be completed before targetSQL
	// is executed against targetPool. Typically tasks is used to
	// set up the joindb before it is queried.
	tasks []tasker

	// targetSQL is the ultimate SQL query to be executed against targetPool.
	targetSQL string

	// targetPool is the destination for the ultimate SQL query to
	// be executed against.
	targetPool driver.Pool
}

// newPipeline parses query, returning a pipeline prepared for
// execution via pipeline.execute.
func newPipeline(ctx context.Context, qc *QueryContext, query string) (*pipeline, error) {
	log := lg.FromContext(ctx)

	a, err := ast.Parse(log, query)
	if err != nil {
		return nil, err
	}

	qModel, err := buildQueryModel(qc, a)
	if err != nil {
		return nil, err
	}

	p := &pipeline{
		qc:    qc,
		query: query,
	}

	if err = p.prepare(ctx, qModel); err != nil {
		return nil, err
	}

	return p, nil
}

// execute executes the pipeline, writing results to recw.
func (p *pipeline) execute(ctx context.Context, recw RecordWriter) error {
	lg.FromContext(ctx).Debug(
		"Execute SQL query",
		lga.Src, p.targetPool.Source(),
		lga.SQL, p.targetSQL,
	)

	if err := p.executeTasks(ctx); err != nil {
		return err
	}

	return QuerySQL(ctx, p.targetPool, nil, recw, p.targetSQL)
}

// executeTasks executes any tasks in pipeline.tasks.
// These tasks may exist if preparatory work must be performed
// before pipeline.targetSQL can be executed.
func (p *pipeline) executeTasks(ctx context.Context) error {
	switch len(p.tasks) {
	case 0:
		return nil
	case 1:
		return p.tasks[0].executeTask(ctx)
	default:
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.SetLimit(driver.OptTuningErrgroupLimit.Get(options.FromContext(ctx)))

	for _, task := range p.tasks {
		task := task

		g.Go(func() error {
			select {
			case <-gCtx.Done():
				return gCtx.Err()
			default:
			}
			return task.executeTask(gCtx)
		})
	}

	return g.Wait()
}

// prepareNoTable is invoked when the queryModel doesn't have a table.
// That is to say, the query doesn't have a "FROM table" clause. It is
// this function's responsibility to figure out what source to use, and
// to set the relevant pipeline fields.
func (p *pipeline) prepareNoTable(ctx context.Context, qm *queryModel) error {
	log := lg.FromContext(ctx)
	log.Debug("No table in query; will look for source to use...")

	var (
		src    *source.Source
		err    error
		handle = ast.NewInspector(qm.AST).FindFirstHandle()
	)

	if handle == "" {
		if src = p.qc.Collection.Active(); src == nil {
			log.Debug("No active source, will use scratchdb.")
			p.targetPool, err = p.qc.ScratchPoolOpener.OpenScratch(ctx, "scratch")
			if err != nil {
				return err
			}

			p.rc = &render.Context{
				Renderer: p.targetPool.SQLDriver().Renderer(),
				Args:     p.qc.Args,
				Dialect:  p.targetPool.SQLDriver().Dialect(),
			}
			return nil
		}

		log.Debug("Using active source.", lga.Src, src)
	} else if src, err = p.qc.Collection.Get(handle); err != nil {
		return err
	}

	// At this point, src is non-nil.
	if p.targetPool, err = p.qc.PoolOpener.Open(ctx, src); err != nil {
		return err
	}

	p.rc = &render.Context{
		Renderer: p.targetPool.SQLDriver().Renderer(),
		Args:     p.qc.Args,
		Dialect:  p.targetPool.SQLDriver().Dialect(),
	}

	return nil
}

// prepareFromTable builds the "FROM table" fragment.
//
// When this function returns, pipeline.rc will be set.
func (p *pipeline) prepareFromTable(ctx context.Context, tblSel *ast.TblSelectorNode) (fromClause string,
	fromPool driver.Pool, err error,
) {
	handle := tblSel.Handle()
	if handle == "" {
		handle = p.qc.Collection.ActiveHandle()
		if handle == "" {
			return "", nil, errz.New("query does not specify source, and no active source")
		}
	}

	src, err := p.qc.Collection.Get(handle)
	if err != nil {
		return "", nil, err
	}

	fromPool, err = p.qc.PoolOpener.Open(ctx, src)
	if err != nil {
		return "", nil, err
	}

	rndr := fromPool.SQLDriver().Renderer()
	p.rc = &render.Context{
		Renderer: rndr,
		Args:     p.qc.Args,
		Dialect:  fromPool.SQLDriver().Dialect(),
	}

	fromClause, err = rndr.FromTable(p.rc, tblSel)
	if err != nil {
		return "", nil, err
	}

	return fromClause, fromPool, nil
}

// joinClause models the SQL "JOIN" construct.
type joinClause struct {
	leftTbl *ast.TblSelectorNode
	joins   []*ast.JoinNode
}

// tables returns a new slice containing all referenced tables.
func (jc *joinClause) tables() []*ast.TblSelectorNode {
	tbls := make([]*ast.TblSelectorNode, len(jc.joins)+1)
	tbls[0] = jc.leftTbl
	for i := range jc.joins {
		tbls[i+1] = jc.joins[i].Table()
	}

	return tbls
}

// handles returns the set of (non-empty) handles from the tables,
// without any duplicates.
func (jc *joinClause) handles() []string {
	handles := make([]string, len(jc.joins)+1)
	handles[0] = jc.leftTbl.Handle()
	for i := 0; i < len(jc.joins); i++ {
		handles[i+1] = jc.joins[i].Table().Handle()
	}

	handles = lo.Uniq(handles)
	handles = lo.Without(handles, "")
	return handles
}

// isSingleSource returns true if the joins refer to the same handle.
func (jc *joinClause) isSingleSource() bool {
	leftHandle := jc.leftTbl.Handle()

	for _, join := range jc.joins {
		joinHandle := join.Table().Handle()
		if joinHandle == "" {
			continue
		}

		if joinHandle != leftHandle {
			return false
		}
	}

	return true
}

// prepareFromJoin builds the "JOIN" clause.
//
// When this function returns, pipeline.rc will be set.
func (p *pipeline) prepareFromJoin(ctx context.Context, jc *joinClause) (fromClause string,
	fromConn driver.Pool, err error,
) {
	if jc.isSingleSource() {
		return p.joinSingleSource(ctx, jc)
	}

	return p.joinCrossSource(ctx, jc)
}

// joinSingleSource sets up a join against a single source.
//
// On return, pipeline.rc will be set.
func (p *pipeline) joinSingleSource(ctx context.Context, jc *joinClause) (fromClause string,
	fromPool driver.Pool, err error,
) {
	src, err := p.qc.Collection.Get(jc.leftTbl.Handle())
	if err != nil {
		return "", nil, err
	}

	fromPool, err = p.qc.PoolOpener.Open(ctx, src)
	if err != nil {
		return "", nil, err
	}

	rndr := fromPool.SQLDriver().Renderer()
	p.rc = &render.Context{
		Renderer: rndr,
		Args:     p.qc.Args,
		Dialect:  fromPool.SQLDriver().Dialect(),
	}

	fromClause, err = rndr.Join(p.rc, jc.leftTbl, jc.joins)
	if err != nil {
		return "", nil, err
	}

	return fromClause, fromPool, nil
}

// joinCrossSource returns a FROM clause that forms part of
// the SQL SELECT statement against fromDB.
//
// On return, pipeline.rc will be set.
func (p *pipeline) joinCrossSource(ctx context.Context, jc *joinClause) (fromClause string,
	fromDB driver.Pool, err error,
) {
	handles := jc.handles()
	srcs := make([]*source.Source, 0, len(handles))
	for _, handle := range handles {
		var src *source.Source
		if src, err = p.qc.Collection.Get(handle); err != nil {
			return "", nil, err
		}
		srcs = append(srcs, src)
	}

	// Open the join db
	joinPool, err := p.qc.JoinPoolOpener.OpenJoin(ctx, srcs...)
	if err != nil {
		return "", nil, err
	}

	rndr := joinPool.SQLDriver().Renderer()
	p.rc = &render.Context{
		Renderer: rndr,
		Args:     p.qc.Args,
		Dialect:  joinPool.SQLDriver().Dialect(),
	}

	leftHandle := jc.leftTbl.Handle()
	// TODO: verify not empty

	tbls := jc.tables()
	for _, tbl := range tbls {
		tbl := tbl
		handle := tbl.Handle()
		if handle == "" {
			handle = leftHandle
		}
		var src *source.Source
		if src, err = p.qc.Collection.Get(handle); err != nil {
			return "", nil, err
		}
		var db driver.Pool
		if db, err = p.qc.PoolOpener.Open(ctx, src); err != nil {
			return "", nil, err
		}

		task := &joinCopyTask{
			fromPool: db,
			fromTbl:  tbl.Table(),
			toPool:   joinPool,
			toTbl:    tbl.TblAliasOrName(),
		}

		tbl.SyncTblNameAlias()

		p.tasks = append(p.tasks, task)
	}

	fromClause, err = rndr.Join(p.rc, jc.leftTbl, jc.joins)
	if err != nil {
		return "", nil, err
	}

	return fromClause, joinPool, nil
}

// tasker is the interface for executing a DB task.
type tasker interface {
	// executeTask executes a task against the DB.
	executeTask(ctx context.Context) error
}

// joinCopyTask is a specification of a table data copy task to be performed
// for a cross-source join. That is, the data in fromDB.fromTblName will
// be copied to a table in toPool. If colNames is
// empty, all cols in fromTbl are to be copied.
type joinCopyTask struct {
	fromPool driver.Pool
	fromTbl  tablefq.T
	toPool   driver.Pool
	toTbl    tablefq.T
}

func (jt *joinCopyTask) executeTask(ctx context.Context) error {
	return execCopyTable(ctx, jt.fromPool, jt.fromTbl, jt.toPool, jt.toTbl)
}

// execCopyTable performs the work of copying fromDB.fromTbl to destPool.destTbl.
func execCopyTable(ctx context.Context, fromDB driver.Pool, fromTbl tablefq.T,
	destPool driver.Pool, destTbl tablefq.T,
) error {
	log := lg.FromContext(ctx)

	createTblHook := func(ctx context.Context, originRecMeta record.Meta, destPool driver.Pool,
		tx sqlz.DB,
	) error {
		destColNames := originRecMeta.Names()
		destColKinds := originRecMeta.Kinds()
		destTblDef := sqlmodel.NewTableDef(destTbl.Table, destColNames, destColKinds)

		err := destPool.SQLDriver().CreateTable(ctx, tx, destTblDef)
		if err != nil {
			return errz.Wrapf(err, "failed to create dest table %s.%s", destPool.Source().Handle, destTbl)
		}

		return nil
	}

	inserter := NewDBWriter(
		destPool,
		destTbl.Table,
		driver.OptTuningRecChanSize.Get(destPool.Source().Options),
		createTblHook,
	)

	query := "SELECT * FROM " + fromTbl.Render(fromDB.SQLDriver().Dialect().Enquote)
	err := QuerySQL(ctx, fromDB, nil, inserter, query)
	if err != nil {
		return errz.Wrapf(err, "insert %s.%s failed", destPool.Source().Handle, destTbl)
	}

	affected, err := inserter.Wait() // Wait for the writer to finish processing
	if err != nil {
		return errz.Wrapf(err, "insert %s.%s failed", destPool.Source().Handle, destTbl)
	}
	log.Debug("Copied rows to dest", lga.Count, affected,
		lga.From, fmt.Sprintf("%s.%s", fromDB.Source().Handle, fromTbl),
		lga.To, fmt.Sprintf("%s.%s", destPool.Source().Handle, destTbl))
	return nil
}
