package table

import (
	"fmt"

	"github.com/engvik/sbanken-go"
	"github.com/jedib0t/go-pretty/v6/table"
)

func (w *Writer) GetCustomer(customer sbanken.Customer) {
	w.table.AppendRow(table.Row{"CustomerID", customer.CustomerID})
	w.table.AppendRow(table.Row{"First Name", customer.FirstName})
	w.table.AppendRow(table.Row{"Last Name", customer.LastName})
	w.table.AppendRow(table.Row{"Email Address", customer.EmailAddress})
	w.table.AppendRow(table.Row{"Date of Birth", customer.DateOfBirth})
	w.table.AppendRow(table.Row{"Postal Address", customer.PostalAddress})
	w.table.AppendRow(table.Row{"Street Address", customer.PostalAddress})

	w.table.AppendRow(table.Row{"Phone Numbers", ""})
	for _, n := range customer.PhoneNumbers {
		w.table.AppendRow(table.Row{"", fmt.Sprintf("%s %s", n.CountryCode, n.Number)})
	}

	w.table.Render()
	w.table.ResetHeaders()
	w.table.ResetRows()
	w.table.ResetFooters()
}
