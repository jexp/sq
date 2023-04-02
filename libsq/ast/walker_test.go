package ast

import (
	"testing"

	"github.com/neilotoole/slogt"
	"golang.org/x/exp/slog"

	"github.com/stretchr/testify/assert"
)

func TestWalker(t *testing.T) {
	log := slogt.New(t)

	// `@mydb1 | .user, .address | join(.uid == .uid) | .uid, .username, .country`
	p := getSLQParser(fixtJoinQuery1)
	query := p.Query()
	ast, err := buildAST(log, query)

	assert.Nil(t, err)
	assert.NotNil(t, ast)

	walker := NewWalker(log, ast)
	count := 0

	visitor := func(log *slog.Logger, w *Walker, node Node) error {
		count++
		return w.visitChildren(node)
	}

	walker.AddVisitor(typeJoinNode, visitor)
	err = walker.Walk()
	assert.Nil(t, err)
	assert.Equal(t, 1, count)

	// test multiple visitors on the same node type
	walker = NewWalker(log, ast)
	countA := 0
	visitorA := func(log *slog.Logger, w *Walker, node Node) error {
		countA++
		return w.visitChildren(node)
	}
	countB := 0
	visitorB := func(log *slog.Logger, w *Walker, node Node) error {
		countB++
		return w.visitChildren(node)
	}

	walker.AddVisitor(typeTblSelectorNode, visitorA)
	walker.AddVisitor(typeColSelectorNode, visitorB)
	err = walker.Walk()
	assert.Nil(t, err)
	assert.Equal(t, 2, countA)
	assert.Equal(t, 3, countB)
}
