package sbanken

import (
	"bytes"
	"context"
	"flag"
	"testing"

	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

var testEfaktura = sbanken.Efaktura{
	ID:                  "test-efaktura",
	IssuerID:            "test-issuer",
	Reference:           "test-ref",
	DocumentType:        "test-doctype",
	Status:              "test-status",
	KID:                 "test-kid",
	OriginalDueDate:     "timestamp",
	NotificationDate:    "timestamp",
	IssuerName:          "test-issuer-name",
	OriginalAmount:      133.33,
	MinimumAmount:       100.00,
	CreditAccountNumber: "test-credit-card-number",
}

var testListEfakturas = `+---------------+------------------+---------------+-------------+-------------------+-------------------+-----------------+-----------------+----------+
| ID            | ISSUER NAME      | DOCUMENT TYPE | STATUS      | NOTIFICATION DATE | ORIGINAL DUE DATE | ORIGINAL AMOUNT | MINIMUM AMOOUNT | KID      |
+---------------+------------------+---------------+-------------+-------------------+-------------------+-----------------+-----------------+----------+
| test-efaktura | test-issuer-name | test-doctype  | test-status | timestamp         | timestamp         |          133.33 |             100 | test-kid |
+---------------+------------------+---------------+-------------+-------------------+-------------------+-----------------+-----------------+----------+
|               |                  |               |             |                   |                   |          133.33 |             100 |          |
+---------------+------------------+---------------+-------------+-------------------+-------------------+-----------------+-----------------+----------+

To see all fields, use: sbanken efakturas read --id=<ID>
Detailed fields includes: Issuer ID, Reference, Update Due Date, Updated Amount, Credit Account Number
`

var testReadEfaktura = `+-----------------------+-------------------------+
| ID                    | test-efaktura           |
| Issuer ID             | test-issuer             |
| Issuer Name           | test-issuer-name        |
| Reference             | test-ref                |
| Document Type         | test-doctype            |
| Status                | test-status             |
| Notification Date     | timestamp               |
| Original Due Date     | timestamp               |
| Update Due Date       |                         |
| Original Amount       | 133.33                  |
| Update Amount         | 0                       |
| Minimum Amount        | 100                     |
| KID                   | test-kid                |
| Credit Account Number | test-credit-card-number |
+-----------------------+-------------------------+
`

func (c testClient) ListEfakturas(context.Context, *sbanken.EfakturaListQuery) ([]sbanken.Efaktura, error) {
	return []sbanken.Efaktura{testEfaktura}, nil
}

func (c testClient) PayEfaktura(context.Context, *sbanken.EfakturaPayQuery) error {
	return nil
}

func (c testClient) ListNewEfakturas(context.Context, *sbanken.EfakturaListQuery) ([]sbanken.Efaktura, error) {
	return []sbanken.Efaktura{testEfaktura}, nil
}

func (c testClient) ReadEfaktura(context.Context, string) (sbanken.Efaktura, error) {
	return testEfaktura, nil
}

func TestListEfakturas(t *testing.T) {
	conn := Connection{
		Client: testClient{},
	}

	var buf bytes.Buffer
	conn.output = &buf

	fs := flag.NewFlagSet("start-date", flag.ExitOnError)
	ctx := cli.NewContext(nil, fs, nil)

	if err := conn.ListEfakturas(ctx); err != nil {
		t.Errorf("error running test: %v", err)
	}

	exp := []byte(testListEfakturas)
	got := buf.Bytes()

	if bytes.Compare(got, exp) != 0 {
		t.Errorf("unexpected bytes: got %s, exp %s", got, exp)
	}
}

func TestPayEfaktura(t *testing.T) {
	conn := Connection{
		Client: testClient{},
	}

	var buf bytes.Buffer
	conn.output = &buf

	fs := flag.NewFlagSet("efaktura", flag.ExitOnError)
	fs.String("id", "test-id", "")
	fs.String("account-id", "test-account-id", "")
	ctx := cli.NewContext(nil, fs, nil)

	if err := conn.PayEfaktura(ctx); err != nil {
		t.Errorf("error running test: %v", err)
	}

	exp := []byte(`Efaktura test-id payed successfully with account test-account-id
`)
	got := buf.Bytes()

	if bytes.Compare(got, exp) != 0 {
		t.Errorf("unexpected bytes: got %s, exp %s", got, exp)
	}

}

func TestListNewEfakturas(t *testing.T) {
	conn := Connection{
		Client: testClient{},
	}

	var buf bytes.Buffer
	conn.output = &buf

	fs := flag.NewFlagSet("start-date", flag.ExitOnError)
	ctx := cli.NewContext(nil, fs, nil)

	if err := conn.ListNewEfakturas(ctx); err != nil {
		t.Errorf("error running test: %v", err)
	}

	exp := []byte(testListEfakturas)
	got := buf.Bytes()

	if bytes.Compare(got, exp) != 0 {
		t.Errorf("unexpected bytes: got %s, exp %s", got, exp)
	}
}

func TestReadEfakturas(t *testing.T) {
	conn := Connection{
		Client: testClient{},
	}

	var buf bytes.Buffer
	conn.output = &buf

	fs := flag.NewFlagSet("id", flag.ExitOnError)
	ctx := cli.NewContext(nil, fs, nil)

	if err := conn.ReadEfaktura(ctx); err != nil {
		t.Errorf("error running test: %v", err)
	}

	exp := []byte(testReadEfaktura)
	got := buf.Bytes()

	if bytes.Compare(got, exp) != 0 {
		t.Errorf("unexpected bytes: got %s, exp %s", got, exp)
	}
}
