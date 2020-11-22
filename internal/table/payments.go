package table

import (
	"fmt"
	"strings"

	"github.com/engvik/sbanken-go"
	"github.com/jedib0t/go-pretty/v6/table"
)

func (w *Writer) ListPayments(payments []sbanken.Payment) {
	if w.colors {
		w.setPaymentsColors()
	}

	w.table.AppendHeader(table.Row{
		"ID",
		"Beneficiary Name",
		"Recipient Account Number",
		"Due Date",
		"KID",
		"Text",
		"Status",
		"Amount",
	})

	var rows []table.Row
	var amount float32

	for _, p := range payments {
		rows = append(rows, table.Row{
			p.ID,
			p.BeneficiaryName,
			p.RecipientAccountNumber,
			p.DueDate,
			p.KID,
			p.Text,
			p.Status,
			p.Amount,
		})

		amount += p.Amount
	}

	w.table.AppendRows(rows)
	w.table.AppendFooter(table.Row{"", "", "", "", "", "", "", amount})
	w.table.Render()

	fmt.Fprintln(w.output)
	fmt.Fprintln(w.output, "To see all fields, use: sbanken payments read --id=<ID>")
	fmt.Fprintln(w.output, "Detailed fields includes: Allowed New Status Types, Status Details, Product Type, Payment Number, Is Active")
}

func (w *Writer) ReadPayment(payment sbanken.Payment) {
	w.table.AppendRow(table.Row{"ID", payment.ID})
	w.table.AppendRow(table.Row{"Beneficiary Name", payment.BeneficiaryName})
	w.table.AppendRow(table.Row{"Recipient Account Number", payment.RecipientAccountNumber})
	w.table.AppendRow(table.Row{"Due Date", payment.DueDate})
	w.table.AppendRow(table.Row{"KID", payment.KID})
	w.table.AppendRow(table.Row{"Text", payment.Text})
	w.table.AppendRow(table.Row{"Status", payment.Status})
	w.table.AppendRow(table.Row{"Amount", payment.Amount})
	w.table.AppendRow(table.Row{"Allowed New Status Types", strings.Join(payment.AllowedNewStatusTypes, ",")})
	w.table.AppendRow(table.Row{"Status Details", payment.StatusDetails})
	w.table.AppendRow(table.Row{"Product Type", payment.ProductType})
	w.table.AppendRow(table.Row{"Payment Number", payment.PaymentNumber})
	w.table.AppendRow(table.Row{"Is Active", payment.IsActive})

	w.table.Render()
}

func (w *Writer) setPaymentsColors() {
	w.table.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:              "Amount",
			Transformer:       w.colorValuesTransformer,
			TransformerFooter: w.colorValuesTransformer,
		},
	})
}
