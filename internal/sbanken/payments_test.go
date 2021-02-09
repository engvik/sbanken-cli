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

var testPayment = sbanken.Payment{
	ID:                     "test-id",
	RecipientAccountNumber: "test-recipient-account-number",
	DueDate:                "test-timestamp",
	KID:                    "test-kid",
	Text:                   "test-text",
	Status:                 "test-status",
	AllowedNewStatusTypes:  []string{"new-status"},
	StatusDetails:          "test-details",
	ProductType:            "test-type",
	PaymentType:            "test-payment-type",
	BeneficiaryName:        "test-name",
	Amount:                 1337.00,
	PaymentNumber:          4,
	IsActive:               true,
}

var testListPaymentsTable = `+---------+------------------+-------------------------------+----------------+----------+-----------+-------------+--------+
| ID      | BENEFICIARY NAME | RECIPIENT ACCOUNT NUMBER      | DUE DATE       | KID      | TEXT      | STATUS      | AMOUNT |
+---------+------------------+-------------------------------+----------------+----------+-----------+-------------+--------+
| test-id | test-name        | test-recipient-account-number | test-timestamp | test-kid | test-text | test-status |   1337 |
+---------+------------------+-------------------------------+----------------+----------+-----------+-------------+--------+
|         |                  |                               |                |          |           |             |   1337 |
+---------+------------------+-------------------------------+----------------+----------+-----------+-------------+--------+

To see all fields, use: sbanken payments read --id=<ID>
Detailed fields includes: Allowed New Status Types, Status Details, Product Type, Payment Number, Is Active
`

var testListPaymentsJSON = `[
  {
    "allowedNewStatusTypes": [
      "new-status"
    ],
    "paymentId": "test-id",
    "recipientAccountNumber": "test-recipient-account-number",
    "dueDate": "test-timestamp",
    "kid": "test-kid",
    "text": "test-text",
    "status": "test-status",
    "statusDetails": "test-details",
    "productType": "test-type",
    "paymentType": "test-payment-type",
    "beneficiaryName": "test-name",
    "amount": 1337,
    "paymentNumber": 4,
    "isActive": true
  }
]
`

var testReadPaymentTable = `+--------------------------+-------------------------------+
| ID                       | test-id                       |
| Beneficiary Name         | test-name                     |
| Recipient Account Number | test-recipient-account-number |
| Due Date                 | test-timestamp                |
| KID                      | test-kid                      |
| Text                     | test-text                     |
| Status                   | test-status                   |
| Amount                   | 1337                          |
| Allowed New Status Types | new-status                    |
| Status Details           | test-details                  |
| Product Type             | test-type                     |
| Payment Number           | 4                             |
| Is Active                | true                          |
+--------------------------+-------------------------------+
`

var testReadPaymentJSON = `{
  "allowedNewStatusTypes": [
    "new-status"
  ],
  "paymentId": "test-id",
  "recipientAccountNumber": "test-recipient-account-number",
  "dueDate": "test-timestamp",
  "kid": "test-kid",
  "text": "test-text",
  "status": "test-status",
  "statusDetails": "test-details",
  "productType": "test-type",
  "paymentType": "test-payment-type",
  "beneficiaryName": "test-name",
  "amount": 1337,
  "paymentNumber": 4,
  "isActive": true
}
`

func (c testClient) ListPayments(context.Context, string, *sbanken.PaymentListQuery) ([]sbanken.Payment, error) {
	return []sbanken.Payment{testPayment}, nil
}

func (c testClient) ReadPayment(context.Context, string, string) (sbanken.Payment, error) {
	return testPayment, nil
}

func TestListPayments(t *testing.T) {
	tableConn := testNewConnection(t, table.NewWriter())
	JSONConn := testNewConnection(t, json.NewWriter())

	tests := []struct {
		name string
		conn Connection
		exp  []byte
	}{
		{
			name: "should write table output correctly",
			conn: tableConn,
			exp:  []byte(testListPaymentsTable),
		},
		{
			name: "should write json output correctly",
			conn: JSONConn,
			exp:  []byte(testListPaymentsJSON),
		},
	}

	var buf bytes.Buffer

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.conn.writer.SetOutputMirror(&buf)

			fs := flag.NewFlagSet("id", flag.ExitOnError)
			ctx := cli.NewContext(nil, fs, nil)

			if err := tc.conn.ListPayments(ctx); err != nil {
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

func TestReadPayment(t *testing.T) {
	tableConn := testNewConnection(t, table.NewWriter())
	JSONConn := testNewConnection(t, json.NewWriter())

	tests := []struct {
		name string
		conn Connection
		exp  []byte
	}{
		{
			name: "should write table output correctly",
			conn: tableConn,
			exp:  []byte(testReadPaymentTable),
		},
		{
			name: "should write json output correctly",
			conn: JSONConn,
			exp:  []byte(testReadPaymentJSON),
		},
	}

	var buf bytes.Buffer

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.conn.writer.SetOutputMirror(&buf)

			fs := flag.NewFlagSet("id", flag.ExitOnError)
			ctx := cli.NewContext(nil, fs, nil)

			if err := tc.conn.ReadPayment(ctx); err != nil {
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
