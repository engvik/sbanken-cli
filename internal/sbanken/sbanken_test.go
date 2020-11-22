package sbanken

import (
	"testing"

	"github.com/engvik/sbanken-cli/internal/table"
)

type testClient struct{}

func testNewConnection(t *testing.T) Connection {
	t.Helper()

	return Connection{
		client: testClient{},
		writer: table.NewWriter(),
	}

}
