package sbanken

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/engvik/sbanken-go"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

func (c *Connection) ListPayments(cliCtx *cli.Context) error {
	ctx := context.Background()

	if err := c.ConnectClient(ctx, cliCtx); err != nil {
		return err
	}

	accountID := cliCtx.String("id")

	payments, err := c.Client.ListPayments(ctx, accountID, nil)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
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

	fmt.Println()
	fmt.Println("To see all fields, use: sbanken payments read --id=<ID>")
	fmt.Println("Detailed fields includes: Allowed New Status Types, Status Details, Product Type, Payment Number, Is Active")

	return nil
}

func (c *Connection) ReadPayment(cliCtx *cli.Context) error {
	ctx := context.Background()

	if err := c.ConnectClient(ctx, cliCtx); err != nil {
		return err
	}

	accountID := cliCtx.String("account-id")
	paymentID := cliCtx.String("id")

	payment, err := c.Client.ReadPayment(ctx, accountID, paymentID)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

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

func parsePaymentsListQuery(ctx *cli.Context) *sbanken.PaymentListQuery {
	q := &sbanken.PaymentListQuery{
		Index:  ctx.String("index"),
		Length: ctx.String("length"),
	}

	return q
}
