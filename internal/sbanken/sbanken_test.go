package sbanken

import (
	"regexp"
	"testing"
)

type testClient struct{}

func testNewConnection(t *testing.T, w outputWriter) Connection {
	t.Helper()

	idRegexp, err := regexp.Compile(".?")
	if err != nil {
		t.Fatalf("error setting up test: %v", err)
	}

	return Connection{
		client:   testClient{},
		writer:   w,
		idRegexp: idRegexp,
	}
}
