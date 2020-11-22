package table

import (
	"github.com/engvik/sbanken-go"
	"github.com/jedib0t/go-pretty/v6/table"
)

func (w *Writer) ListAccounts(accounts []sbanken.Account) {
	w.table.AppendHeader(table.Row{"ID", "Type", "Name", "Number", "Balance", "Available", "Credit Limit"})

	var rows []table.Row
	var balance float32
	var available float32
	var creditLimit float32

	for _, a := range accounts {
		rows = append(rows, table.Row{a.ID, a.Type, a.Name, a.Number, a.Balance, a.Available, a.CreditLimit})
		balance += a.Balance
		available += a.Available
		creditLimit += a.CreditLimit
	}

	w.table.AppendRows(rows)
	w.table.AppendFooter(table.Row{"", "", "", "", balance, available, creditLimit})
	w.table.Render()
}

func (w *Writer) ReadAccount(account sbanken.Account) {
	w.table.AppendRow(table.Row{"ID", account.ID})
	w.table.AppendRow(table.Row{"Type", account.Type})
	w.table.AppendRow(table.Row{"Name", account.Name})
	w.table.AppendRow(table.Row{"Number", account.Number})
	w.table.AppendRow(table.Row{"Balance", account.Balance})
	w.table.AppendRow(table.Row{"Available", account.Available})
	w.table.AppendRow(table.Row{"Credit Limit", account.CreditLimit})

	w.table.Render()
}
