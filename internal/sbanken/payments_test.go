package sbanken

import (
	"bytes"
	"context"
	"flag"
	"testing"

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

var testListPayments = `+---------+------------------+-------------------------------+----------------+----------+-----------+-------------+--------+
| ID      | BENEFICIARY NAME | RECIPIENT ACCOUNT NUMBER      | DUE DATE       | KID      | TEXT      | STATUS      | AMOUNT |
+---------+------------------+-------------------------------+----------------+----------+-----------+-------------+--------+
| test-id | test-name        | test-recipient-account-number | test-timestamp | test-kid | test-text | test-status |   1337 |
+---------+------------------+-------------------------------+----------------+----------+-----------+-------------+--------+
|         |                  |                               |                |          |           |             |   1337 |
+---------+------------------+-------------------------------+----------------+----------+-----------+-------------+--------+

To see all fields, use: sbanken payments read --id=<ID>
Detailed fields includes: Allowed New Status Types, Status Details, Product Type, Payment Number, Is Active
`

var testReadPayment = `+--------------------------+-------------------------------+
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

func (c testClient) ListPayments(context.Context, string, *sbanken.PaymentListQuery) ([]sbanken.Payment, error) {
	return []sbanken.Payment{testPayment}, nil
}

func (c testClient) ReadPayment(context.Context, string, string) (sbanken.Payment, error) {
	return testPayment, nil
}

func TestListPayments(t *testing.T) {
	conn := Connection{
		Client: testClient{},
	}

	var buf bytes.Buffer
	conn.output = &buf

	fs := flag.NewFlagSet("id", flag.ExitOnError)
	ctx := cli.NewContext(nil, fs, nil)

	if err := conn.ListPayments(ctx); err != nil {
		t.Errorf("error running test: %v", err)
	}

	exp := []byte(testListPayments)
	got := buf.Bytes()

	if bytes.Compare(got, exp) != 0 {
		t.Errorf("unexpected bytes: got %s, exp %s", got, exp)
	}
}

func TestReadPayment(t *testing.T) {
	conn := Connection{
		Client: testClient{},
	}

	var buf bytes.Buffer
	conn.output = &buf

	fs := flag.NewFlagSet("id", flag.ExitOnError)
	ctx := cli.NewContext(nil, fs, nil)

	if err := conn.ReadPayment(ctx); err != nil {
		t.Errorf("error running test: %v", err)
	}

	exp := []byte(testReadPayment)
	got := buf.Bytes()

	if bytes.Compare(got, exp) != 0 {
		t.Errorf("unexpected bytes: got %s, exp %s", got, exp)
	}
}
