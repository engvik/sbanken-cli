package sbanken

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"testing"

	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

var testTransaction = sbanken.Transaction{
	TransactionDetails: sbanken.TransactionDetails{
		ID:                     "test-id",
		FormattedAccountNumber: "test-formatted-account",
		CID:                    "test-cid",
		AmountDescription:      "test-desc",
		ReceiverName:           "rest-receiver-name",
		PayerName:              "test-payer-name",
		RegistrationDate:       "timestamp",
		NumericReference:       15,
	},
	CardDetails: sbanken.CardDetails{
		CardNumber:                  "test-card-number",
		MerchantCategoryCode:        "test-merachant-cat-code",
		MerchantCategoryDescription: "test-merchant-desc",
		MerchantCity:                "test-merchant-city",
		MerchantName:                "test-merchant-name",
		OriginalCurrencyCode:        "test-org-currency-code",
		PurchaseDate:                "test-purchase-date",
		TransactionID:               "test-transaction-id",
		CurrencyAmount:              1337,
		CurrencyRate:                13.37,
	},
	AccountingDate:              "timestamp",
	InterestDate:                "timestampt",
	OtherAccountNumber:          "test-other-accounts",
	Text:                        "test-text",
	TransactionType:             "test-type",
	TransactionTypeText:         "test-type-text",
	ReservationType:             "test-reservation-type",
	Source:                      "source",
	Amount:                      999.99,
	IsReservation:               true,
	OtherAccountNumberSpecified: true,
	TransactionDetailSpecified:  true,
	CardDetailsSpecified:        true,
}

var testListTransactions = `+-------+-----------+--------+-----------------+---------------+------------------+----------------+
| INDEX | TEXT      | AMOUNT | ACCOUNTING DATE | INTEREST DATE | TRANSACTION TYPE | IS RESERVATION |
+-------+-----------+--------+-----------------+---------------+------------------+----------------+
|     0 | test-text | 999.99 | timestamp       | timestampt    | test-type        | true           |
+-------+-----------+--------+-----------------+---------------+------------------+----------------+
|       |           | 999.99 |                 |               |                  |                |
+-------+-----------+--------+-----------------+---------------+------------------+----------------+
`

var testListTransactionsDetailsText = `
To see detailed output, use: sbanken transactions list --id=<ID> --details
Detailed fields includes: Other Account Number, Transaction Type Text, Transaction Type Code, Reservation Type, Source

Some transactions contains card details, to list them use: sbanken transactions list --id=<ID> --card-details
Some transactions has more transaction details, to list them use: sbanken transactions list --id=<ID> --transaction-details
`

var testListTransactionsDetails = `+-------+----------------------+-----------------------+-----------------------+-----------------------+--------+
| INDEX | OTHER ACCOUNT NUMBER | TRANSACTION TYPE TEXT | TRANSACTION TYPE CODE | RESERVATION TYPE      | SOURCE |
+-------+----------------------+-----------------------+-----------------------+-----------------------+--------+
|     0 | test-other-accounts  | test-type-text        |                     0 | test-reservation-type | source |
+-------+----------------------+-----------------------+-----------------------+-----------------------+--------+
`

var testListTransactionsCardDetails = `+-------+-------------------------+-------------------------------+--------------------+--------------------+------------------------+--------------------+---------------------+-----------------+---------------+
| INDEX | MERCHANT CATEGORY CODE  | MERCHANT CATEGORY DESCRIPTION | MERCHANT CITY      | MERCHANT NAME      | ORIGINAL CURRENCY CODE | PURCHASE DATE      | TRANSACTION ID      | CURRENCY AMOUNT | CURRENCY RATE |
+-------+-------------------------+-------------------------------+--------------------+--------------------+------------------------+--------------------+---------------------+-----------------+---------------+
|     0 | test-merachant-cat-code | test-merchant-desc            | test-merchant-city | test-merchant-name | test-org-currency-code | test-purchase-date | test-transaction-id |            1337 |         13.37 |
+-------+-------------------------+-------------------------------+--------------------+--------------------+------------------------+--------------------+---------------------+-----------------+---------------+
`

var testListTransactionsTransactionDetails = `+-------+---------+--------------------------+----------+--------------------+--------------------+-----------------+-------------------+-------------------+
| INDEX | ID      | FORMATTED ACCOUNT NUMBER | CID      | AMOUNT DESCRIPTION | RECEIVER NAME      | PAYER NAME      | REGISTRATION DATE | NUMERIC REFERENCE |
+-------+---------+--------------------------+----------+--------------------+--------------------+-----------------+-------------------+-------------------+
|     0 | test-id | test-formatted-account   | test-cid | test-desc          | rest-receiver-name | test-payer-name | timestamp         |                15 |
+-------+---------+--------------------------+----------+--------------------+--------------------+-----------------+-------------------+-------------------+
`

func (c testClient) ListTransactions(context.Context, string, *sbanken.TransactionListQuery) ([]sbanken.Transaction, error) {
	return []sbanken.Transaction{testTransaction}, nil
}

func TestListTransactions(t *testing.T) {
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
			exp: fmt.Sprintf(
				"%s%s",
				testListTransactions,
				testListTransactionsDetailsText,
			),
		},
		{
			name: "should print expected output with detailed output",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("details", flag.ExitOnError)
				fs.Bool("details", true, "")
				return fs
			}(),
			exp: fmt.Sprintf(
				"%s%s",
				testListTransactions,
				testListTransactionsDetails,
			),
		},
		{
			name: "should print expected output with card details",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("card-details", flag.ExitOnError)
				fs.Bool("card-details", true, "")
				return fs
			}(),
			exp: fmt.Sprintf(
				"%s%s",
				testListTransactions,
				testListTransactionsCardDetails,
			),
		},
		{
			name: "should print expected output with transaction details",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("transaction-details", flag.ExitOnError)
				fs.Bool("transaction-details", true, "")
				return fs
			}(),
			exp: fmt.Sprintf(
				"%s%s",
				testListTransactions,
				testListTransactionsTransactionDetails,
			),
		},
		{
			name: "should print expected output with all details",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("all-details", flag.ExitOnError)
				fs.Bool("details", true, "")
				fs.Bool("card-details", true, "")
				fs.Bool("transaction-details", true, "")
				return fs
			}(),
			exp: fmt.Sprintf(
				"%s%s%s%s",
				testListTransactions,
				testListTransactionsDetails,
				testListTransactionsCardDetails,
				testListTransactionsTransactionDetails,
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			conn.output = &buf

			ctx := cli.NewContext(nil, tc.fs, nil)

			if err := conn.ListTransactions(ctx); err != nil {
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
