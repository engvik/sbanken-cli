package table

import (
	"github.com/engvik/sbanken-go"
	"github.com/jedib0t/go-pretty/v6/table"
)

func (w *Writer) ListCards(cards []sbanken.Card) {
	w.table.AppendHeader(table.Row{"ID", "Type", "Product Code", "Customer ID", "Account Owner", "Number", "Account Number", "ExpiryDate", "Status", "Version Number"})

	var rows []table.Row

	for _, card := range cards {
		rows = append(rows, table.Row{card.ID, card.Type, card.ProductCode, card.CustomerID, card.AccountOwner, card.Number, card.AccountNumber, card.ExpiryDate, card.Status, card.VersionNumber})
	}

	w.table.AppendRows(rows)
	w.table.Render()
}
