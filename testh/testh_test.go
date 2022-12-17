package testh_test

import (
	"github.com/neilotoole/sq/testh/tutil"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/ryboe/q"

	"github.com/neilotoole/sq/drivers/csv"
	"github.com/neilotoole/sq/libsq/core/sqlz"
	"github.com/neilotoole/sq/libsq/core/stringz"
	"github.com/neilotoole/sq/libsq/source"
	"github.com/neilotoole/sq/testh"
	"github.com/neilotoole/sq/testh/proj"
	"github.com/neilotoole/sq/testh/sakila"
)

func TestVal(t *testing.T) {
	want := "hello"
	var got any

	if tutil.Val(nil) != nil {
		t.FailNow()
	}

	var v0 any
	if tutil.Val(v0) != nil {
		t.FailNow()
	}

	var v1 = want
	var v1a any = want
	var v2 = &v1
	var v3 any = &v1
	var v4 = &v2
	var v5 = &v4

	vals := []any{v1, v1a, v2, v3, v4, v5}
	for _, val := range vals {
		got = tutil.Val(val)

		if got != want {
			t.Errorf("expected %T(%v) but got %T(%v)", want, want, got, got)
		}
	}

	slice := []string{"a", "b"}
	require.Equal(t, slice, tutil.Val(slice))
	require.Equal(t, slice, tutil.Val(&slice))

	b := true
	require.Equal(t, b, tutil.Val(b))
	require.Equal(t, b, tutil.Val(&b))

	type structT struct {
		f string
	}

	st1 := structT{f: "hello"}
	require.Equal(t, st1, tutil.Val(st1))
	require.Equal(t, st1, tutil.Val(&st1))

	var c chan int
	require.Nil(t, tutil.Val(c))
	c = make(chan int, 10)
	require.Equal(t, c, tutil.Val(c))
	require.Equal(t, c, tutil.Val(&c))
}

func TestCopyRecords(t *testing.T) {
	var v1, v2, v3, v4, v5, v6 = int64(1), float64(1.1), true, "hello", []byte("hello"), time.Unix(0, 0)

	testCases := map[string][]sqlz.Record{
		"nil":   nil,
		"empty": {},
		"vals": {
			{nil, &v1, &v2, &v3, &v4, &v5, &v6},
			// {nil, &v1, &v2, &v3, &v4, &v5, &v6},
		},
	}

	for name, recs := range testCases {
		name, recs := name, recs

		t.Run(name, func(t *testing.T) {
			recs2 := testh.CopyRecords(recs)
			require.True(t, len(recs) == len(recs2))

			if recs == nil {
				require.True(t, recs2 == nil)
				return
			}

			for i := range recs {
				require.True(t, len(recs[i]) == len(recs2[i]))
				for j := range recs[i] {
					if recs[i][j] == nil {
						require.True(t, recs2[i][j] == nil)
						continue
					}

					require.False(t, recs[i][j] == recs2[i][j],
						"pointer values should not be equal: %#v --> %#v", recs[i][j], recs2[i][j])

					val1, val2 := tutil.Val(recs[i][j]), tutil.Val(recs2[i][j])
					require.Equal(t, val1, val2,
						"dereferenced values should be equal: %#v --> %#v", val1, val2)
				}
			}
		})
	}
}

func TestRecordsFromTbl(t *testing.T) {
	recMeta1, recs1 := testh.RecordsFromTbl(t, sakila.SL3, sakila.TblActor)
	require.Equal(t, sakila.TblActorColKinds(), recMeta1.Kinds())

	recs1[0][0] = t.Name()

	recMeta2, recs2 := testh.RecordsFromTbl(t, sakila.SL3, sakila.TblActor)
	require.False(t, &recMeta1 == &recMeta2, "should be distinct copies")
	require.False(t, &recs1 == &recs2, "should be distinct copies")
	require.NotEqual(t, recs1[0][0], recs2[0][0], "recs2 should not have the mutated value from recs1")
}

func TestHelper_Files(t *testing.T) {
	fpath := "drivers/csv/testdata/person.csv"
	wantBytes := proj.ReadFile(fpath)

	src := &source.Source{
		Handle:   "@test_" + stringz.Uniq8(),
		Type:     csv.TypeCSV,
		Location: proj.Abs(fpath),
	}

	th := testh.New(t)
	fs := th.Files()

	typ, err := fs.Type(th.Context, src.Location)
	require.NoError(t, err)
	require.Equal(t, src.Type, typ)

	g, _ := errgroup.WithContext(th.Context)

	for i := 0; i < 1000; i++ {
		g.Go(func() error {
			r, err := fs.Open(src)
			require.NoError(t, err)

			defer func() { require.NoError(t, r.Close()) }()

			b, err := ioutil.ReadAll(r)
			require.NoError(t, err)

			require.Equal(t, wantBytes, b)
			return nil
		})
	}

	err = g.Wait()
	require.NoError(t, err)
}

func TestTName(t *testing.T) {
	testCases := []struct {
		a    []any
		want string
	}{
		{a: []any{}, want: "empty"},
		{a: []any{"test", 1}, want: "test_1"},
		{a: []any{"/file/path/name"}, want: "_file_path_name"},
	}

	for _, tc := range testCases {
		got := tutil.Name(tc.a...)
		require.Equal(t, tc.want, got)
	}

}

// Keep the q lib around
var _ = q.Q
