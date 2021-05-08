package sbanken

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"testing"

	"github.com/engvik/sbanken-cli/internal/output/json"
	"github.com/engvik/sbanken-cli/internal/output/table"
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
	TransactionID:               "",
	ReservationType:             "test-reservation-type",
	Source:                      "source",
	Amount:                      999.99,
	IsReservation:               true,
	OtherAccountNumberSpecified: true,
	TransactionDetailSpecified:  true,
	CardDetailsSpecified:        true,
}

var testListTransactionsTable = `+-------+----+-----------+--------+-----------------+---------------+------------------+----------------+
| INDEX | ID | TEXT      | AMOUNT | ACCOUNTING DATE | INTEREST DATE | TRANSACTION TYPE | IS RESERVATION |
+-------+----+-----------+--------+-----------------+---------------+------------------+----------------+
|     0 |    | test-text | 999.99 | timestamp       | timestampt    | test-type        | true           |
+-------+----+-----------+--------+-----------------+---------------+------------------+----------------+
|       |    | 999.99    |        |                 |               |                  |                |
+-------+----+-----------+--------+-----------------+---------------+------------------+----------------+
`

var testListTransactionsTableDetailsText = `
To see detailed output, use: sbanken transactions list --id=<ID> --details
Detailed fields includes: Other Account Number, Transaction Type Text, Transaction Type Code, Reservation Type, Source

Some transactions contains card details, to list them use: sbanken transactions list --id=<ID> --card-details
Some transactions has more transaction details, to list them use: sbanken transactions list --id=<ID> --transaction-details
`

var testListTransactionsTableDetails = `+-------+----------------------+-----------------------+-----------------------+-----------------------+--------+
| INDEX | OTHER ACCOUNT NUMBER | TRANSACTION TYPE TEXT | TRANSACTION TYPE CODE | RESERVATION TYPE      | SOURCE |
+-------+----------------------+-----------------------+-----------------------+-----------------------+--------+
|     0 | test-other-accounts  | test-type-text        |                     0 | test-reservation-type | source |
+-------+----------------------+-----------------------+-----------------------+-----------------------+--------+
`

var testListTransactionsTableCardDetails = `+-------+-------------------------+-------------------------------+--------------------+--------------------+------------------------+--------------------+---------------------+-----------------+---------------+
| INDEX | MERCHANT CATEGORY CODE  | MERCHANT CATEGORY DESCRIPTION | MERCHANT CITY      | MERCHANT NAME      | ORIGINAL CURRENCY CODE | PURCHASE DATE      | TRANSACTION ID      | CURRENCY AMOUNT | CURRENCY RATE |
+-------+-------------------------+-------------------------------+--------------------+--------------------+------------------------+--------------------+---------------------+-----------------+---------------+
|     0 | test-merachant-cat-code | test-merchant-desc            | test-merchant-city | test-merchant-name | test-org-currency-code | test-purchase-date | test-transaction-id |            1337 |         13.37 |
+-------+-------------------------+-------------------------------+--------------------+--------------------+------------------------+--------------------+---------------------+-----------------+---------------+
`

var testListTransactionsTableTransactionDetails = `+-------+---------+--------------------------+----------+--------------------+--------------------+-----------------+-------------------+-------------------+
| INDEX | ID      | FORMATTED ACCOUNT NUMBER | CID      | AMOUNT DESCRIPTION | RECEIVER NAME      | PAYER NAME      | REGISTRATION DATE | NUMERIC REFERENCE |
+-------+---------+--------------------------+----------+--------------------+--------------------+-----------------+-------------------+-------------------+
|     0 | test-id | test-formatted-account   | test-cid | test-desc          | rest-receiver-name | test-payer-name | timestamp         |                15 |
+-------+---------+--------------------------+----------+--------------------+--------------------+-----------------+-------------------+-------------------+
`

var testListTransactionsJSON = `[
  {
    "cardDetails": {
      "cardNumber": "test-card-number",
      "merchantCategoryCode": "test-merachant-cat-code",
      "merchantCategoryDescription": "test-merchant-desc",
      "merchantCity": "test-merchant-city",
      "merchantName": "test-merchant-name",
      "originalCurrencyCode": "test-org-currency-code",
      "purchaseDate": "test-purchase-date",
      "transactionId": "test-transaction-id",
      "currencyAmount": 1337,
      "currencyRate": 13.37
    },
    "transactionDetails": {
      "transactionId": "test-id",
      "formattedAccountNumber": "test-formatted-account",
      "cid": "test-cid",
      "amountDescription": "test-desc",
      "receiverName": "rest-receiver-name",
      "payerName": "test-payer-name",
      "registrationDate": "timestamp",
      "numericReference": 15
    },
    "accountingDate": "timestamp",
    "interestDate": "timestampt",
    "otherAccountNumber": "test-other-accounts",
    "text": "test-text",
    "transactionType": "test-type",
    "transactionTypeText": "test-type-text",
    "reservationType": "test-reservation-type",
    "transactionId": "",
    "source": "source",
    "amount": 999.99,
    "transactionTypeCode": 0,
    "isReservation": true,
    "cardDetailsSpecified": true,
    "otherAccountNumberSpecified": true,
    "transactionDetailSpecified": true
  }
]
`

func (c testClient) ListTransactions(context.Context, string, *sbanken.TransactionListQuery) ([]sbanken.Transaction, error) {
	return []sbanken.Transaction{testTransaction}, nil
}

func (c testClient) ListArchivedTransactions(context.Context, string, *sbanken.TransactionListQuery) ([]sbanken.Transaction, error) {
	return []sbanken.Transaction{testTransaction}, nil
}

func TestListTransactions(t *testing.T) {
	tableConn := testNewConnection(t, table.NewWriter())
	JSONConn := testNewConnection(t, json.NewWriter())

	var buf bytes.Buffer

	tests := []struct {
		name string
		fs   *flag.FlagSet
		conn Connection
		exp  string
	}{
		{
			name: "should write expected table output without any details",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("id", flag.ExitOnError)
				return fs
			}(),
			conn: tableConn,
			exp: fmt.Sprintf(
				"%s%s",
				testListTransactionsTable,
				testListTransactionsTableDetailsText,
			),
		},
		{
			name: "should write expected table output with detailed output",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("details", flag.ExitOnError)
				fs.Bool("details", true, "")
				return fs
			}(),
			conn: tableConn,
			exp: fmt.Sprintf(
				"%s%s",
				testListTransactionsTable,
				testListTransactionsTableDetails,
			),
		},
		{
			name: "should write expected table output with card details",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("card-details", flag.ExitOnError)
				fs.Bool("card-details", true, "")
				return fs
			}(),
			conn: tableConn,
			exp: fmt.Sprintf(
				"%s%s",
				testListTransactionsTable,
				testListTransactionsTableCardDetails,
			),
		},
		{
			name: "should write expected table output with transaction details",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("transaction-details", flag.ExitOnError)
				fs.Bool("transaction-details", true, "")
				return fs
			}(),
			conn: tableConn,
			exp: fmt.Sprintf(
				"%s%s",
				testListTransactionsTable,
				testListTransactionsTableTransactionDetails,
			),
		},
		{
			name: "should write expected table output with all details",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("all-details", flag.ExitOnError)
				fs.Bool("details", true, "")
				fs.Bool("card-details", true, "")
				fs.Bool("transaction-details", true, "")
				return fs
			}(),
			conn: tableConn,
			exp: fmt.Sprintf(
				"%s%s%s%s",
				testListTransactionsTable,
				testListTransactionsTableDetails,
				testListTransactionsTableCardDetails,
				testListTransactionsTableTransactionDetails,
			),
		},
		{
			name: "should write expected json output",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("all-details", flag.ExitOnError)
				fs.Bool("details", true, "")
				fs.Bool("card-details", true, "")
				fs.Bool("transaction-details", true, "")
				return fs
			}(),
			conn: JSONConn,
			exp:  testListTransactionsJSON,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.conn.writer.SetOutputMirror(&buf)

			ctx := cli.NewContext(nil, tc.fs, nil)

			if err := tc.conn.ListTransactions(ctx); err != nil {
				t.Errorf("error running test: %v", err)
			}

			exp := []byte(tc.exp)
			got := buf.Bytes()

			if bytes.Compare(got, exp) != 0 {
				t.Errorf("unexpected bytes: got %s, exp %s", got, exp)
			}

			buf.Reset()
		})
	}
}
