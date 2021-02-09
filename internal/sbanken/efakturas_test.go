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

var testListEfakturasTable = `+---------------+------------------+---------------+-------------+-------------------+-------------------+-----------------+-----------------+----------+
| ID            | ISSUER NAME      | DOCUMENT TYPE | STATUS      | NOTIFICATION DATE | ORIGINAL DUE DATE | ORIGINAL AMOUNT | MINIMUM AMOOUNT | KID      |
+---------------+------------------+---------------+-------------+-------------------+-------------------+-----------------+-----------------+----------+
| test-efaktura | test-issuer-name | test-doctype  | test-status | timestamp         | timestamp         |          133.33 |             100 | test-kid |
+---------------+------------------+---------------+-------------+-------------------+-------------------+-----------------+-----------------+----------+
|               |                  |               |             |                   |                   |          133.33 |             100 |          |
+---------------+------------------+---------------+-------------+-------------------+-------------------+-----------------+-----------------+----------+

To see all fields, use: sbanken efakturas read --id=<ID>
Detailed fields includes: Issuer ID, Reference, Update Due Date, Updated Amount, Credit Account Number
`

var testListEfakturasJSON = `[
  {
    "eFakturaId": "test-efaktura",
    "issuerId": "test-issuer",
    "eFakturaReference": "test-ref",
    "documentType": "test-doctype",
    "status": "test-status",
    "kid": "test-kid",
    "originalDueDate": "timestamp",
    "updatedDueDate": "",
    "notificationDate": "timestamp",
    "issuerName": "test-issuer-name",
    "creditAccountNumber": "test-credit-card-number",
    "originalAmount": 133.33,
    "updatedAmount": 0,
    "minimumAmount": 100
  }
]
`

var testReadEfakturaTable = `+-----------------------+-------------------------+
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

var testReadEfakturaJSON = `{
  "eFakturaId": "test-efaktura",
  "issuerId": "test-issuer",
  "eFakturaReference": "test-ref",
  "documentType": "test-doctype",
  "status": "test-status",
  "kid": "test-kid",
  "originalDueDate": "timestamp",
  "updatedDueDate": "",
  "notificationDate": "timestamp",
  "issuerName": "test-issuer-name",
  "creditAccountNumber": "test-credit-card-number",
  "originalAmount": 133.33,
  "updatedAmount": 0,
  "minimumAmount": 100
}
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
			exp:  []byte(testListEfakturasTable),
		},
		{
			name: "should write json output correctly",
			conn: JSONConn,
			exp:  []byte(testListEfakturasJSON),
		},
	}

	var buf bytes.Buffer

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.conn.writer.SetOutputMirror(&buf)

			fs := flag.NewFlagSet("start-date", flag.ExitOnError)
			ctx := cli.NewContext(nil, fs, nil)

			if err := tc.conn.ListEfakturas(ctx); err != nil {
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

func TestPayEfaktura(t *testing.T) {
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
			exp: []byte(`Efaktura test-id paid successfully with account test-account-id
`),
		},
		{
			name: "should write json output correctly",
			conn: JSONConn,
			exp: []byte(`{
  "eFakturaId": "test-id",
  "accountId": "test-account-id",
  "PayOnlyMinimumAmount": false
}
`),
		},
	}

	var buf bytes.Buffer

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.conn.writer.SetOutputMirror(&buf)

			fs := flag.NewFlagSet("efaktura", flag.ExitOnError)
			fs.String("id", "test-id", "")
			fs.String("account-id", "test-account-id", "")
			ctx := cli.NewContext(nil, fs, nil)

			if err := tc.conn.PayEfaktura(ctx); err != nil {
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

func TestListNewEfakturas(t *testing.T) {
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
			exp:  []byte(testListEfakturasTable),
		},
		{
			name: "should write json output correctly",
			conn: JSONConn,
			exp:  []byte(testListEfakturasJSON),
		},
	}

	var buf bytes.Buffer

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.conn.writer.SetOutputMirror(&buf)

			fs := flag.NewFlagSet("start-date", flag.ExitOnError)
			ctx := cli.NewContext(nil, fs, nil)

			if err := tc.conn.ListNewEfakturas(ctx); err != nil {
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

func TestReadEfakturas(t *testing.T) {
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
			exp:  []byte(testReadEfakturaTable),
		},
		{
			name: "should write json output correctly",
			conn: JSONConn,
			exp:  []byte(testReadEfakturaJSON),
		},
	}

	var buf bytes.Buffer

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.conn.writer.SetOutputMirror(&buf)

			fs := flag.NewFlagSet("id", flag.ExitOnError)
			ctx := cli.NewContext(nil, fs, nil)

			if err := tc.conn.ReadEfaktura(ctx); err != nil {
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
