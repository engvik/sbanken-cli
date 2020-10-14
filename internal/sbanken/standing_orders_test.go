package sbanken

import (
	"bytes"
	"context"
	"flag"
	"testing"

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

var testListStandingOrders = `+-------------------+----------------------------+---------------------------+---------------+-------------------+--------+----------------+
| STANDING ORDER ID | CREDIT ACCOUNT NUMBER      | DEBIT ACCOUNT NUMBER      | NEXT DUE DATE | LAST PAYMENT DATE | AMOUNT | FREQUENCY      |
+-------------------+----------------------------+---------------------------+---------------+-------------------+--------+----------------+
|                19 | test-credit-account-number | test-debit-account-number | timestamp     | timestamp         |   1337 | test-frequency |
+-------------------+----------------------------+---------------------------+---------------+-------------------+--------+----------------+
|                   |                            |                           |               |                   |   1337 |                |
+-------------------+----------------------------+---------------------------+---------------+-------------------+--------+----------------+

To see detailed output, use: sbanken standingorders list --id=<ID> --details
Detailed fields includes: CID, Beneficiary Name, Standing Order Start Date, Standing Order End Date, Standing Order Type, Free Terms
`

var testListStandingOrdersDetails = `+-------------------+----------------------------+---------------------------+---------------+-------------------+--------+----------------+
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

func (c testClient) ListStandingOrders(context.Context, string) ([]sbanken.StandingOrder, error) {
	return []sbanken.StandingOrder{testStandingOrder}, nil
}

func TestListStandingOrders(t *testing.T) {
	conn := Connection{
		Client: testClient{},
	}

	tests := []struct {
		name string
		fs   *flag.FlagSet
		exp  string
	}{
		{
			name: "should print expected output without any details",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("id", flag.ExitOnError)
				return fs
			}(),
			exp: testListStandingOrders,
		},
		{
			name: "should print expected output with details",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("id", flag.ExitOnError)
				fs.Bool("details", true, "")
				return fs
			}(),
			exp: testListStandingOrdersDetails,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			conn.output = &buf

			ctx := cli.NewContext(nil, tc.fs, nil)

			if err := conn.ListStandingOrders(ctx); err != nil {
				t.Errorf("error running test: %v", err)
			}

			exp := []byte(tc.exp)
			got := buf.Bytes()

			if bytes.Compare(got, exp) != 0 {
				t.Errorf("unexpected bytes: got %s, exp %s", got, exp)
			}
		})
	}
}
