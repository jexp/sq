package drivers_test

import (
	"testing"

	"github.com/neilotoole/sq/drivers/postgres"
	"github.com/neilotoole/sq/drivers/sqlite3"
	"github.com/neilotoole/sq/drivers/sqlserver"

	"github.com/neilotoole/sq/testh/sakila"

	"github.com/neilotoole/sq/drivers/mysql"

	"github.com/neilotoole/sq/libsq/source"

	_ "github.com/mattn/go-sqlite3"
)

func TestQuery_datetime(t *testing.T) {
	testCases := []queryTestCase{
		{
			name:     "datetime/strftime/sqlite",
			in:       `@sakila | .payment | strftime("%m", .payment_date)`,
			wantSQL:  `SELECT strftime('%m', "payment_date") AS "strftime(""%m"",.payment_date)" FROM "payment"`,
			onlyFor:  []source.Type{sqlite3.Type},
			wantRecs: sakila.TblPaymentCount,
		},
		{
			name:     "datetime/date_trunc/postgres",
			in:       `@sakila | .payment | date_trunc("month", .payment_date)`,
			wantSQL:  `SELECT date_trunc('month', "payment_date") AS "date_trunc(""month"",.payment_date)" FROM "payment"`,
			onlyFor:  []source.Type{postgres.Type},
			wantRecs: sakila.TblPaymentCount,
		},
		{
			name:     "datetime/month/sqlserver",
			in:       `@sakila | .payment | month(.payment_date)`,
			wantSQL:  `SELECT month("payment_date") AS "month(.payment_date)" FROM "payment"`,
			onlyFor:  []source.Type{sqlserver.Type},
			wantRecs: sakila.TblPaymentCount,
		},
		{
			name:     "datetime/date_format/mysql",
			in:       `@sakila | .payment | date_format(.payment_date, "%m")`,
			wantSQL:  "SELECT date_format(`payment_date`, '%m') AS `date_format(.payment_date,\"%m\")` FROM `payment`",
			onlyFor:  []source.Type{mysql.Type},
			wantRecs: sakila.TblPaymentCount,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			execQueryTestCase(t, tc)
		})
	}
}
