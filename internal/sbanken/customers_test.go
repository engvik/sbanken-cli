package sbanken

import (
	"bytes"
	"context"
	"testing"

	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

var testCustomer = sbanken.Customer{
	FirstName:    "test-first",
	LastName:     "test-last",
	EmailAddress: "test@test.com",
	DateOfBirth:  "2021-01-31T10:05:54.590Z",
	PostalAddress: sbanken.Address{
		AddressLine1: "Tester street 1",
	},
	StreetAddress: sbanken.Address{
		AddressLine1: "Tester street 1",
	},
	PhoneNumbers: []sbanken.PhoneNumber{
		{
			"1",
			"1337133713371337",
		},
	},
}

var testGetCustomer = `+----------------+--------------------------+
| First Name     | test-first               |
| Last Name      | test-last                |
| Email Address  | test@test.com            |
| Date of Birth  | 2021-01-31T10:05:54.590Z |
| Postal Address | {Tester street 1     }   |
| Street Address | {Tester street 1     }   |
| Phone Numbers  |                          |
|                | 1 1337133713371337       |
+----------------+--------------------------+
`

func (c testClient) GetCustomer(context.Context) (sbanken.Customer, error) {
	return testCustomer, nil
}

func TestGetCustomer(t *testing.T) {
	conn := testNewConnection(t)

	var buf bytes.Buffer
	conn.writer.SetOutputMirror(&buf)

	if err := conn.GetCustomer(&cli.Context{}); err != nil {
		t.Errorf("error running test: %v", err)
	}

	exp := []byte(testGetCustomer)
	got := buf.Bytes()

	if bytes.Compare(got, exp) != 0 {
		t.Errorf("unexpected bytes: got %s, exp %s", got, exp)
	}

}
