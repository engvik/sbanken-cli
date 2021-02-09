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

var testAccount = sbanken.Account{
	ID:          "test-id",
	Name:        "test-name",
	Type:        "test-type",
	Number:      "test-number",
	Available:   13.37,
	Balance:     13.37,
	CreditLimit: 0,
}

var testListAccountsTable = `+---------+-----------+-----------+-------------+---------+-----------+--------------+
| ID      | TYPE      | NAME      | NUMBER      | BALANCE | AVAILABLE | CREDIT LIMIT |
+---------+-----------+-----------+-------------+---------+-----------+--------------+
| test-id | test-type | test-name | test-number |   13.37 |     13.37 |            0 |
+---------+-----------+-----------+-------------+---------+-----------+--------------+
|         |           |           |             |   13.37 |     13.37 |            0 |
+---------+-----------+-----------+-------------+---------+-----------+--------------+
`

var testListAccountsJSON = `[
  {
    "accountId": "test-id",
    "name": "test-name",
    "accountType": "test-type",
    "accountNumber": "test-number",
    "available": 13.37,
    "balance": 13.37,
    "creditLimit": 0
  }
]
`

var testReadAccountTable = `+--------------+-------------+
| ID           | test-id     |
| Type         | test-type   |
| Name         | test-name   |
| Number       | test-number |
| Balance      | 13.37       |
| Available    | 13.37       |
| Credit Limit | 0           |
+--------------+-------------+
`

var testReadAccountJSON = `{
  "accountId": "test-id",
  "name": "test-name",
  "accountType": "test-type",
  "accountNumber": "test-number",
  "available": 13.37,
  "balance": 13.37,
  "creditLimit": 0
}
`

func (c testClient) ListAccounts(context.Context) ([]sbanken.Account, error) {
	return []sbanken.Account{testAccount}, nil
}

func (c testClient) ReadAccount(context.Context, string) (sbanken.Account, error) {
	return testAccount, nil

}

func TestListAccounts(t *testing.T) {
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
			exp:  []byte(testListAccountsTable),
		},
		{
			name: "should write json output correctly",
			conn: JSONConn,
			exp:  []byte(testListAccountsJSON),
		},
	}

	var buf bytes.Buffer

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.conn.writer.SetOutputMirror(&buf)

			if err := tc.conn.ListAccounts(&cli.Context{}); err != nil {
				t.Errorf("error running test: %v", err)
			}

			got := buf.Bytes()

			if bytes.Compare(got, tc.exp) != 0 {
				t.Errorf("unexpected bytes: got %s, exp %s", got, tc.exp)
			}
		})

		buf.Reset()
	}
}

func TestReadAccount(t *testing.T) {
	tests := []struct {
		name string
		conn Connection
		exp  []byte
	}{
		{
			name: "should list table output correctly",
			conn: testNewConnection(t, table.NewWriter()),
			exp:  []byte(testReadAccountTable),
		},
		{
			name: "should list json output correctly",
			conn: testNewConnection(t, json.NewWriter()),
			exp:  []byte(testReadAccountJSON),
		},
	}

	var buf bytes.Buffer

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.conn.writer.SetOutputMirror(&buf)

			fs := flag.NewFlagSet("id", flag.ExitOnError)
			ctx := cli.NewContext(nil, fs, nil)

			if err := tc.conn.ReadAccount(ctx); err != nil {
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
