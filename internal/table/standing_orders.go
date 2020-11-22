package table

import (
	"fmt"
	"strings"

	"github.com/engvik/sbanken-go"
	"github.com/jedib0t/go-pretty/v6/table"
)

func (w *Writer) ListStandingOrders(standingOrders []sbanken.StandingOrder, detailedOutput bool) {
	w.table.AppendHeader(table.Row{
		"Standing Order ID",
		"Credit Account Number",
		"Debit Account Number",
		"Next Due Date",
		"Last Payment Date",
		"Amount",
		"Frequency",
	})

	var rows []table.Row
	var amount float32

	for _, so := range standingOrders {
		rows = append(rows, table.Row{
			so.StandingOrderID,
			so.CreditAccountNumber,
			so.DebitAccountNumber,
			so.NextDueDate,
			so.LastPaymentDate,
			so.Amount,
			so.Frequency,
		})

		amount += so.Amount
	}

	w.table.AppendRows(rows)
	w.table.AppendFooter(table.Row{"", "", "", "", "", amount})
	w.table.Render()
	w.table.ResetHeaders()
	w.table.ResetRows()
	w.table.ResetFooters()

	if detailedOutput {
		w.table.AppendHeader(table.Row{
			"Standing Order ID",
			"CID",
			"Beneficiary Name",
			"Standing Order Start Date",
			"Standing Order End Date",
			"Standing Order Type",
			"Free Terms",
		})

		var rows []table.Row

		for _, so := range standingOrders {
			rows = append(rows, table.Row{
				so.StandingOrderID,
				so.CID,
				so.BeneficiaryName,
				so.StandingOrderStartDate,
				so.StandingOrderEndDate,
				so.StandingOrderType,
				strings.Join(so.FreeTerms, ","),
			})
		}
		w.table.AppendRows(rows)
		w.table.Render()
	} else {
		fmt.Fprintln(w.output)
		fmt.Fprintln(w.output, "To see detailed output, use: sbanken standingorders list --id=<ID> --details")
		fmt.Fprintln(w.output, "Detailed fields includes: CID, Beneficiary Name, Standing Order Start Date, Standing Order End Date, Standing Order Type, Free Terms")
	}
}
