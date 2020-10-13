package sbanken

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

func (c *Connection) ListStandingOrders(ctx *cli.Context) error {
	accountID := ctx.String("id")
	detailedOutput := ctx.Bool("details")

	standingOrders, err := c.Client.ListStandingOrders(ctx.Context, accountID)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
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

	t.AppendRows(rows)
	t.AppendFooter(table.Row{"", "", "", "", "", amount})
	t.Render()

	if detailedOutput {
		td := table.NewWriter()
		td.SetOutputMirror(os.Stdout)
		td.AppendHeader(table.Row{
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
		td.AppendRows(rows)
		td.Render()
	} else {
		fmt.Println()
		fmt.Println("To see detailed output, use: sbanken standingorders list --id=<ID> --details")
		fmt.Println("Detailed fields includes: CID, Beneficiary Name, Standing Order Start Date, Standing Order End Date, Standing Order Type, Free Terms")
	}

	return nil
}
