package sbanken

import (
	"regexp"
	"testing"

	"github.com/engvik/sbanken-cli/internal/table"
)

type testClient struct{}

func testNewConnection(t *testing.T) Connection {
	t.Helper()

	idRegexp, err := regexp.Compile(".?")
	if err != nil {
		t.Fatalf("error setting up test: %v", err)
	}

	return Connection{
		client:   testClient{},
		writer:   table.NewWriter(),
		idRegexp: idRegexp,
	}
}
