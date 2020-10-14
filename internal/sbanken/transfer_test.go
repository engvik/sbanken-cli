package sbanken

import (
	"bytes"
	"context"
	"flag"
	"testing"

	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

func (c testClient) Transfer(context.Context, *sbanken.TransferQuery) error {
	return nil
}

func TestTransfer(t *testing.T) {
	conn := Connection{
		Client: testClient{},
	}

	tests := []struct {
		name string
		fs   *flag.FlagSet
		exp  string
	}{
		{
			name: "should print expected output without message",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("id", flag.ExitOnError)
				fs.String("from", "test-from-account-id", "")
				fs.String("to", "test-to-account-id", "")
				fs.Int("amount", 1337, "")
				return fs
			}(),
			exp: `1337.000000 successfully transferred from test-from-account-id to test-to-account-id
`,
		},
		{
			name: "should print expected output with message",
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("id", flag.ExitOnError)
				fs.String("from", "test-from-account-id", "")
				fs.String("to", "test-to-account-id", "")
				fs.String("message", "$$$", "")
				fs.Int("amount", 1337, "")
				return fs
			}(),
			exp: `1337.000000 successfully transferred from test-from-account-id to test-to-account-id: $$$
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			conn.output = &buf

			ctx := cli.NewContext(nil, tc.fs, nil)

			if err := conn.Transfer(ctx); err != nil {
				t.Errorf("error running test: %v", err)
			}

			exp := []byte(tc.exp)
			got := buf.Bytes()

			if bytes.Compare(got, exp) != 0 {
				t.Errorf("unexpected bytes: got %s, exp %s", got, exp)
			}
		})
	}
}
