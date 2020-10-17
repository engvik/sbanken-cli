package sbanken

import (
	"fmt"
	"strings"

	"github.com/engvik/sbanken-go"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

// ListPayments handles the payments list command.
func (c *Connection) ListPayments(ctx *cli.Context) error {
	accountID := ctx.String("id")
	q := parsePaymentListQuery(ctx)

	payments, err := c.Client.ListPayments(ctx.Context, accountID, q)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(c.output)
	t.AppendHeader(table.Row{
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

	t.AppendRows(rows)
	t.AppendFooter(table.Row{"", "", "", "", "", "", "", amount})
	t.Render()

	fmt.Fprintln(c.output)
	fmt.Fprintln(c.output, "To see all fields, use: sbanken payments read --id=<ID>")
	fmt.Fprintln(c.output, "Detailed fields includes: Allowed New Status Types, Status Details, Product Type, Payment Number, Is Active")

	return nil
}

// ReadPayment handles the payments read command.
func (c *Connection) ReadPayment(ctx *cli.Context) error {
	accountID := ctx.String("account-id")
	paymentID := ctx.String("id")

	payment, err := c.Client.ReadPayment(ctx.Context, accountID, paymentID)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(c.output)

	t.AppendRow(table.Row{"ID", payment.ID})
	t.AppendRow(table.Row{"Beneficiary Name", payment.BeneficiaryName})
	t.AppendRow(table.Row{"Recipient Account Number", payment.RecipientAccountNumber})
	t.AppendRow(table.Row{"Due Date", payment.DueDate})
	t.AppendRow(table.Row{"KID", payment.KID})
	t.AppendRow(table.Row{"Text", payment.Text})
	t.AppendRow(table.Row{"Status", payment.Status})
	t.AppendRow(table.Row{"Amount", payment.Amount})
	t.AppendRow(table.Row{"Allowed New Status Types", strings.Join(payment.AllowedNewStatusTypes, ",")})
	t.AppendRow(table.Row{"Status Details", payment.StatusDetails})
	t.AppendRow(table.Row{"Product Type", payment.ProductType})
	t.AppendRow(table.Row{"Payment Number", payment.PaymentNumber})
	t.AppendRow(table.Row{"Is Active", payment.IsActive})

	t.Render()

	return nil
}

func parsePaymentListQuery(ctx *cli.Context) *sbanken.PaymentListQuery {
	q := &sbanken.PaymentListQuery{
		Index:  ctx.String("index"),
		Length: ctx.String("length"),
	}

	return q
}
