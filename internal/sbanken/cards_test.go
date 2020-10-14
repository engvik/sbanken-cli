package sbanken

import (
	"bytes"
	"context"
	"testing"

	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

var testCard = sbanken.Card{
	ID:            "test-card",
	Number:        "123456789",
	ExpiryDate:    "timestamp",
	Status:        "test-status",
	Type:          "test-type",
	ProductCode:   "test-code",
	AccountNumber: "987654321",
	VersionNumber: 5,
}

var testListCards = `+-----------+-----------+--------------+-----------+----------------+------------+-------------+----------------+
| ID        | TYPE      | PRODUCT CODE | NUMBER    | ACCOUNT NUMBER | EXPIRYDATE | STATUS      | VERSION NUMBER |
+-----------+-----------+--------------+-----------+----------------+------------+-------------+----------------+
| test-card | test-type | test-code    | 123456789 | 987654321      | timestamp  | test-status |              5 |
+-----------+-----------+--------------+-----------+----------------+------------+-------------+----------------+
`

func (c testClient) ListCards(context.Context) ([]sbanken.Card, error) {
	return []sbanken.Card{testCard}, nil
}

func TestListCards(t *testing.T) {
	conn := Connection{
		Client: testClient{},
	}

	var buf bytes.Buffer
	conn.output = &buf

	if err := conn.ListCards(&cli.Context{}); err != nil {
		t.Errorf("error running test: %v", err)
	}

	exp := []byte(testListCards)
	got := buf.Bytes()

	if bytes.Compare(got, exp) != 0 {
		t.Errorf("unexpected bytes: got %s, exp %s", got, exp)
	}
}
