package sbanken

import (
	"bytes"
	"context"
	"flag"
	"testing"

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

var testListAccounts = `+---------+-----------+-----------+-------------+---------+-----------+--------------+
| ID      | TYPE      | NAME      | NUMBER      | BALANCE | AVAILABLE | CREDIT LIMIT |
+---------+-----------+-----------+-------------+---------+-----------+--------------+
| test-id | test-type | test-name | test-number |   13.37 |     13.37 |            0 |
+---------+-----------+-----------+-------------+---------+-----------+--------------+
|         |           |           |             |   13.37 |     13.37 |            0 |
+---------+-----------+-----------+-------------+---------+-----------+--------------+
`

var testReadAccount = `+--------------+-------------+
| ID           | test-id     |
| Type         | test-type   |
| Name         | test-name   |
| Number       | test-number |
| Balance      | 13.37       |
| Available    | 13.37       |
| Credit Limit | 0           |
+--------------+-------------+
`

func (c testClient) ListAccounts(context.Context) ([]sbanken.Account, error) {
	return []sbanken.Account{testAccount}, nil
}

func (c testClient) ReadAccount(context.Context, string) (sbanken.Account, error) {
	return testAccount, nil

}

func TestListAccounts(t *testing.T) {
	conn := testNewConnection(t)

	var buf bytes.Buffer
	conn.writer.SetOutputMirror(&buf)

	if err := conn.ListAccounts(&cli.Context{}); err != nil {
		t.Errorf("error running test: %v", err)
	}

	exp := []byte(testListAccounts)
	got := buf.Bytes()

	if bytes.Compare(got, exp) != 0 {
		t.Errorf("unexpected bytes: got %s, exp %s", got, exp)
	}
}

func TestReadAccount(t *testing.T) {
	conn := testNewConnection(t)

	var buf bytes.Buffer
	conn.writer.SetOutputMirror(&buf)

	fs := flag.NewFlagSet("id", flag.ExitOnError)
	ctx := cli.NewContext(nil, fs, nil)

	if err := conn.ReadAccount(ctx); err != nil {
		t.Errorf("error running test: %v", err)
	}

	exp := []byte(testReadAccount)
	got := buf.Bytes()

	if bytes.Compare(got, exp) != 0 {
		t.Errorf("unexpected bytes: got %s, exp %s", got, exp)
	}
}
