package sbanken

import (
	"bytes"
	"context"
	"flag"
	"testing"

	"github.com/engvik/sbanken-cli/internal/output/json"
	"github.com/engvik/sbanken-cli/internal/output/table"
	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

var testStandingOrder = sbanken.StandingOrder{
	BeneficiaryName:        "test-name",
	CID:                    "test-cid",
	CreditAccountNumber:    "test-credit-account-number",
	DebitAccountNumber:     "test-debit-account-number",
	Frequency:              "test-frequency",
	LastPaymentDate:        "timestamp",
	NextDueDate:            "timestamp",
	StandingOrderEndDate:   "timestamp",
	StandingOrderStartDate: "timestamp",
	StandingOrderType:      "test-type",
	Amount:                 1337.00,
	StandingOrderID:        19,
}

var testListStandingOrdersTable = `+-------------------+----------------------------+---------------------------+---------------+-------------------+--------+----------------+
| STANDING ORDER ID | CREDIT ACCOUNT NUMBER      | DEBIT ACCOUNT NUMBER      | NEXT DUE DATE | LAST PAYMENT DATE | AMOUNT | FREQUENCY      |
+-------------------+----------------------------+---------------------------+---------------+-------------------+--------+----------------+
|                19 | test-credit-account-number | test-debit-account-number | timestamp     | timestamp         |   1337 | test-frequency |
+-------------------+----------------------------+---------------------------+---------------+-------------------+--------+----------------+
|                   |                            |                           |               |                   |   1337 |                |
+-------------------+----------------------------+---------------------------+---------------+-------------------+--------+----------------+

To see detailed output, use: sbanken standingorders list --id=<ID> --details
Detailed fields includes: CID, Beneficiary Name, Standing Order Start Date, Standing Order End Date, Standing Order Type, Free Terms
`

var testListStandingOrdersTableDetails = `+-------------------+----------------------------+---------------------------+---------------+-------------------+--------+----------------+
| STANDING ORDER ID | CREDIT ACCOUNT NUMBER      | DEBIT ACCOUNT NUMBER      | NEXT DUE DATE | LAST PAYMENT DATE | AMOUNT | FREQUENCY      |
+-------------------+----------------------------+---------------------------+---------------+-------------------+--------+----------------+
|                19 | test-credit-account-number | test-debit-account-number | timestamp     | timestamp         |   1337 | test-frequency |
+-------------------+----------------------------+---------------------------+---------------+-------------------+--------+----------------+
|                   |                            |                           |               |                   |   1337 |                |
+-------------------+----------------------------+---------------------------+---------------+-------------------+--------+----------------+
+-------------------+----------+------------------+---------------------------+-------------------------+---------------------+------------+
| STANDING ORDER ID | CID      | BENEFICIARY NAME | STANDING ORDER START DATE | STANDING ORDER END DATE | STANDING ORDER TYPE | FREE TERMS |
+-------------------+----------+------------------+---------------------------+-------------------------+---------------------+------------+
|                19 | test-cid | test-name        | timestamp                 | timestamp               | test-type           |            |
+-------------------+----------+------------------+---------------------------+-------------------------+---------------------+------------+
`

var testListStandingOrdersJSON = `[
  {
    "freeTerms": null,
    "beneficiaryName": "test-name",
    "cId": "test-cid",
    "creditAccountNumber": "test-credit-account-number",
    "debitAccountNumber": "test-debit-account-number",
    "frequency": "test-frequency",
    "lastPaymentDate": "timestamp",
    "nextDueDate": "timestamp",
    "standingOrderEndDate": "timestamp",
    "standingOrderStartDate": "timestamp",
    "standingOrderType": "test-type",
    "amount": 1337,
    "standingOrderId": 19
  }
]
`

func (c testClient) ListStandingOrders(context.Context, string) ([]sbanken.StandingOrder, error) {
	return []sbanken.StandingOrder{testStandingOrder}, nil
}

func TestListStandingOrders(t *testing.T) {
	tableConn := testNewConnection(t, table.NewWriter())
	JSONConn := testNewConnection(t, json.NewWriter())

	var buf bytes.Buffer

	tests := []struct {
		name string
		fs   *flag.FlagSet
		conn Connection
		exp  []byte
	}{
		{
			name: "should write table output without any details correctly",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("id", flag.ExitOnError)
				return fs
			}(),
			conn: tableConn,
			exp:  []byte(testListStandingOrdersTable),
		},
		{
			name: "should write table output with details correctly",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("id", flag.ExitOnError)
				fs.Bool("details", true, "")
				return fs
			}(),
			conn: tableConn,
			exp:  []byte(testListStandingOrdersTableDetails),
		},
		{
			name: "should write json output correctly",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("id", flag.ExitOnError)
				fs.Bool("details", true, "")
				return fs
			}(),
			conn: JSONConn,
			exp:  []byte(testListStandingOrdersJSON),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.conn.writer.SetOutputMirror(&buf)

			ctx := cli.NewContext(nil, tc.fs, nil)

			if err := tc.conn.ListStandingOrders(ctx); err != nil {
				t.Errorf("error running test: %v", err)
			}

			got := buf.Bytes()

			if bytes.Compare(got, tc.exp) != 0 {
				t.Errorf("unexpected bytes: got %s, exp %s", got, tc.exp)
			}

			buf.Reset()
		})
	}
}
