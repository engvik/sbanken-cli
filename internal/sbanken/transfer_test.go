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

func (c testClient) Transfer(context.Context, *sbanken.TransferQuery) error {
	return nil
}

func TestTransfer(t *testing.T) {
	tableConn := testNewConnection(t, table.NewWriter())
	JSONConn := testNewConnection(t, json.NewWriter())

	var buf bytes.Buffer

	tests := []struct {
		name string
		fs   *flag.FlagSet
		conn Connection
		exp  []byte
	}{
		{
			name: "should write expected table output without message",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("id", flag.ExitOnError)
				fs.String("from", "test-from-account-id", "")
				fs.String("to", "test-to-account-id", "")
				fs.Int("amount", 1337, "")
				return fs
			}(),
			conn: tableConn,
			exp: []byte(`1337.000000 successfully transferred from test-from-account-id to test-to-account-id
`),
		},
		{
			name: "should write expected table output with message",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("id", flag.ExitOnError)
				fs.String("from", "test-from-account-id", "")
				fs.String("to", "test-to-account-id", "")
				fs.String("message", "$$$", "")
				fs.Int("amount", 1337, "")
				return fs
			}(),
			conn: tableConn,
			exp: []byte(`1337.000000 successfully transferred from test-from-account-id to test-to-account-id: $$$
`),
		},
		{
			name: "should write expected json output",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("id", flag.ExitOnError)
				fs.String("from", "test-from-account-id", "")
				fs.String("to", "test-to-account-id", "")
				fs.String("message", "$$$", "")
				fs.Int("amount", 1337, "")
				return fs
			}(),
			conn: JSONConn,
			exp: []byte(`{
  "fromAccountId": "test-from-account-id",
  "toAccountId": "test-to-account-id",
  "message": "$$$",
  "amount": 1337
}
`),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.conn.writer.SetOutputMirror(&buf)
			ctx := cli.NewContext(nil, tc.fs, nil)

			if err := tc.conn.Transfer(ctx); err != nil {
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
