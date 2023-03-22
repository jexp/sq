package drivers_test

import (
	"strings"
	"testing"

	"github.com/neilotoole/sq/drivers/mysql"

	"github.com/neilotoole/sq/libsq/source"

	"github.com/neilotoole/sq/libsq"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
	"github.com/neilotoole/sq/testh"
	"github.com/neilotoole/sq/testh/sakila"
)

//nolint:exhaustive,lll
func TestSLQ2SQLNew(t *testing.T) {
	testCases := []struct {
		name     string
		in       string
		wantErr  bool
		want     string
		override map[source.Type]string
	}{
		{
			name:     "select/cols",
			in:       `@sakila | .actor | .first_name, .last_name`,
			want:     `SELECT "first_name", "last_name" FROM "actor"`,
			override: map[source.Type]string{mysql.Type: "SELECT `first_name`, `last_name` FROM `actor`"},
		},
		{
			name:     "select/cols-whitespace",
			in:       `@sakila | .actor | ."first name"`,
			want:     `SELECT "first name" FROM "actor"`,
			override: map[source.Type]string{mysql.Type: "SELECT `first name` FROM `actor`"},
		},
		{
			name:     "select/cols-whitespace-2",
			in:       `@sakila | .actor | .actor_id, ."first name", ."last name"`,
			want:     `SELECT "actor_id", "first name", "last name" FROM "actor"`,
			override: map[source.Type]string{mysql.Type: "SELECT `actor_id`, `first name`, `last name` FROM `actor`"},
		},
		{
			name:     "select/count-whitespace-col",
			in:       `@sakila | .actor | count(."first name")`,
			want:     `SELECT COUNT("first name") FROM "actor"`,
			override: map[source.Type]string{mysql.Type: "SELECT COUNT(`first name`) FROM `actor`"},
		},
		{
			name:     "select/table-whitespace",
			in:       `@sakila | ."film actor"`,
			want:     `SELECT * FROM "film actor"`,
			override: map[source.Type]string{mysql.Type: "SELECT * FROM `film actor`"},
		},
		{
			name:     "select/cols-aliases",
			in:       `@sakila | .actor | .first_name:given_name, .last_name:family_name`,
			want:     `SELECT "first_name" AS "given_name", "last_name" AS "family_name" FROM "actor"`,
			override: map[source.Type]string{mysql.Type: "SELECT `first_name` AS `given_name`, `last_name` AS `family_name` FROM `actor`"},
		},
		{
			name:     "select/count-star",
			in:       `@sakila | .actor | count(*)`,
			want:     `SELECT COUNT(*) FROM "actor"`,
			override: map[source.Type]string{mysql.Type: "SELECT COUNT(*) FROM `actor`"},
		},
		{
			name:     "select/count",
			in:       `@sakila | .actor | count()`,
			want:     `SELECT COUNT(*) FROM "actor"`,
			override: map[source.Type]string{mysql.Type: "SELECT COUNT(*) FROM `actor`"},
		},
		{
			name:     "select/handle-table/cols",
			in:       `@sakila.actor | .first_name, .last_name`,
			want:     `SELECT "first_name", "last_name" FROM "actor"`,
			override: map[source.Type]string{mysql.Type: "SELECT `first_name`, `last_name` FROM `actor`"},
		},
		{
			name:     "select/handle-table/count-star",
			in:       `@sakila.actor | count(*)`,
			want:     `SELECT COUNT(*) FROM "actor"`,
			override: map[source.Type]string{mysql.Type: "SELECT COUNT(*) FROM `actor`"},
		},
		{
			name:     "select/handle-table/count-col",
			in:       `@sakila.actor | count(."first name")`,
			want:     `SELECT COUNT("first name") FROM "actor"`,
			override: map[source.Type]string{mysql.Type: "SELECT COUNT(`first name`) FROM `actor`"},
		},
		{
			name:     "select/count-alias",
			in:       `@sakila | .actor | count(*):quantity`,
			want:     `SELECT COUNT(*) AS "quantity" FROM "actor"`,
			override: map[source.Type]string{mysql.Type: "SELECT COUNT(*) AS `quantity` FROM `actor`"},
		},
		{
			name:     "join/fq-table-cols-equal",
			in:       `@sakila | .actor, .film_actor | join(.film_actor.actor_id == .actor.actor_id)`,
			want:     `SELECT * FROM "actor" INNER JOIN "film_actor" ON "film_actor"."actor_id" = "actor"."actor_id"`,
			override: map[source.Type]string{mysql.Type: "SELECT * FROM `actor` INNER JOIN `film_actor` ON `film_actor`.`actor_id` = `actor`.`actor_id`"},
		},
		{
			name:     "join/fq-table-cols-equal-whitespace",
			in:       `@sakila | .actor, ."film actor" | join(."film actor".actor_id == .actor.actor_id)`,
			want:     `SELECT * FROM "actor" INNER JOIN "film actor" ON "film actor"."actor_id" = "actor"."actor_id"`,
			override: map[source.Type]string{mysql.Type: "SELECT * FROM `actor` INNER JOIN `film actor` ON `film actor`.`actor_id` = `actor`.`actor_id`"},
		},
	}

	srcs := testh.New(t).NewSourceSet(sakila.SQLLatest()...)
	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			for _, src := range srcs.Items() {
				src := src

				t.Run(string(src.Type), func(t *testing.T) {
					in := strings.Replace(tc.in, "@sakila", src.Handle, 1)
					t.Logf(in)
					want := tc.want
					if overrideWant, ok := tc.override[src.Type]; ok {
						want = overrideWant
					}

					_, err := srcs.SetActive(src.Handle)
					require.NoError(t, err)

					th := testh.New(t)
					dbases := th.Databases()

					// drvr := th.DriverFor(src).(driver.SQLDriver)
					// drvr.AlterTableAddColumn()

					gotSQL, gotErr := libsq.SLQ2SQL(th.Context, th.Log, dbases, dbases, srcs, in)
					if tc.wantErr {
						require.Error(t, gotErr)
						return
					}

					require.NoError(t, gotErr)
					require.Equal(t, want, gotSQL)
				})
			}
		})
	}
}
