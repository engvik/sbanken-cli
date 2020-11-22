package table

import (
	"fmt"

	"github.com/engvik/sbanken-go"
	"github.com/jedib0t/go-pretty/v6/table"
)

func (w *Writer) ListEfakturas(efakturas []sbanken.Efaktura) {
	w.table.AppendHeader(table.Row{
		"ID",
		"Issuer Name",
		"Document Type",
		"Status",
		"Notification Date",
		"Original Due Date",
		"Original Amount",
		"Minimum Amoount",
		"KID",
	})

	var rows []table.Row
	var originalAmount float32
	var minimumAmount float32

	for _, e := range efakturas {
		rows = append(rows, table.Row{
			e.ID,
			e.IssuerName,
			e.DocumentType,
			e.Status,
			e.NotificationDate,
			e.OriginalDueDate,
			e.OriginalAmount,
			e.MinimumAmount,
			e.KID,
		})

		originalAmount += e.OriginalAmount
		minimumAmount += e.MinimumAmount
	}

	w.table.AppendRows(rows)
	w.table.AppendFooter(table.Row{
		"",
		"",
		"",
		"",
		"",
		"",
		originalAmount,
		minimumAmount,
	})
	w.table.Render()

	fmt.Fprintln(w.output)
	fmt.Fprintln(w.output, "To see all fields, use: sbanken efakturas read --id=<ID>")
	fmt.Fprintln(w.output, "Detailed fields includes: Issuer ID, Reference, Update Due Date, Updated Amount, Credit Account Number")
}

func (w *Writer) PayEfaktura(q *sbanken.EfakturaPayQuery) {
	fmt.Fprintf(w.output, "Efaktura %s paid successfully with account %s\n", q.ID, q.AccountID)
}

func (w *Writer) ReadEfaktura(efaktura sbanken.Efaktura) {
	w.table.AppendRow(table.Row{"ID", efaktura.ID})
	w.table.AppendRow(table.Row{"Issuer ID", efaktura.IssuerID})
	w.table.AppendRow(table.Row{"Issuer Name", efaktura.IssuerName})
	w.table.AppendRow(table.Row{"Reference", efaktura.Reference})
	w.table.AppendRow(table.Row{"Document Type", efaktura.DocumentType})
	w.table.AppendRow(table.Row{"Status", efaktura.Status})
	w.table.AppendRow(table.Row{"Notification Date", efaktura.NotificationDate})
	w.table.AppendRow(table.Row{"Original Due Date", efaktura.OriginalDueDate})
	w.table.AppendRow(table.Row{"Update Due Date", efaktura.UpdatedDueDate})
	w.table.AppendRow(table.Row{"Original Amount", efaktura.OriginalAmount})
	w.table.AppendRow(table.Row{"Update Amount", efaktura.UpdatedAmount})
	w.table.AppendRow(table.Row{"Minimum Amount", efaktura.MinimumAmount})
	w.table.AppendRow(table.Row{"KID", efaktura.KID})
	w.table.AppendRow(table.Row{"Credit Account Number", efaktura.CreditAccountNumber})

	w.table.Render()
}
