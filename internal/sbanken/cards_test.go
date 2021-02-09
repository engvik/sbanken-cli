package sbanken

import (
	"bytes"
	"context"
	"testing"

	"github.com/engvik/sbanken-cli/internal/output/json"
	"github.com/engvik/sbanken-cli/internal/output/table"
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

var testListCardsTable = `+-----------+-----------+--------------+-----------+----------------+------------+-------------+----------------+
| ID        | TYPE      | PRODUCT CODE | NUMBER    | ACCOUNT NUMBER | EXPIRYDATE | STATUS      | VERSION NUMBER |
+-----------+-----------+--------------+-----------+----------------+------------+-------------+----------------+
| test-card | test-type | test-code    | 123456789 | 987654321      | timestamp  | test-status |              5 |
+-----------+-----------+--------------+-----------+----------------+------------+-------------+----------------+
`

var testListCardsJSON = `[
  {
    "cardId": "test-card",
    "cardNumber": "123456789",
    "expiryDate": "timestamp",
    "status": "test-status",
    "cardType": "test-type",
    "productCode": "test-code",
    "accountNumber": "987654321",
    "cardVersionNumber": 5
  }
]
`

func (c testClient) ListCards(context.Context) ([]sbanken.Card, error) {
	return []sbanken.Card{testCard}, nil
}

func TestListCards(t *testing.T) {
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
			exp:  []byte(testListCardsTable),
		},
		{
			name: "should write json output correctly",
			conn: JSONConn,
			exp:  []byte(testListCardsJSON),
		},
	}

	var buf bytes.Buffer

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.conn.writer.SetOutputMirror(&buf)

			if err := tc.conn.ListCards(&cli.Context{}); err != nil {
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
