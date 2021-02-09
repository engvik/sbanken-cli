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

var testGetCustomerTable = `+----------------+--------------------------+
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

var testGetCustomerJSON = `{
  "customerID": "",
  "firstName": "test-first",
  "lastName": "test-last",
  "emailAddress": "test@test.com",
  "dateOfBirth": "2021-01-31T10:05:54.590Z",
  "postalAddress": {
    "addressLine1": "Tester street 1",
    "addressLine2": "",
    "addressLine3": "",
    "country": "",
    "zipCode": "",
    "city": ""
  },
  "streetAddress": {
    "addressLine1": "Tester street 1",
    "addressLine2": "",
    "addressLine3": "",
    "country": "",
    "zipCode": "",
    "city": ""
  },
  "phoneNumbers": [
    {
      "countryCode": "1",
      "number": "1337133713371337"
    }
  ]
}
`

func (c testClient) GetCustomer(context.Context) (sbanken.Customer, error) {
	return testCustomer, nil
}

func TestGetCustomer(t *testing.T) {
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
			exp:  []byte(testGetCustomerTable),
		},
		{
			name: "should write json output correctly",
			conn: JSONConn,
			exp:  []byte(testGetCustomerJSON),
		},
	}

	var buf bytes.Buffer

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.conn.writer.SetOutputMirror(&buf)

			if err := tc.conn.GetCustomer(&cli.Context{}); err != nil {
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
