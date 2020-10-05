package sbanken

import (
	"context"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/urfave/cli/v2"
)

func (c *Connection) ListNewEfakturas(cliCtx *cli.Context) error {
	ctx := context.Background()

	if err := c.ConnectClient(ctx, cliCtx); err != nil {
		return err
	}

	efakturas, err := c.Client.ListNewEfakturas(ctx, nil)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"ID",
		"Issuer ID",
		"Issuer Name",
		"Reference",
		"Document Type",
		"Status",
		"Notification Date",
		"Original Due Date",
		"Updated Due Date",
		"Original Amount",
		"Updated Amount",
		"Minimum Amoount",
		"KID",
		"Credit Account Number",
	})

	var rows []table.Row
	var originalAmount float32
	var updatedAmount float32
	var minimumAmount float32

	for _, e := range efakturas {
		rows = append(rows, table.Row{
			e.ID,
			e.IssuerID,
			e.IssuerName,
			e.Reference,
			e.DocumentType,
			e.Status,
			e.NotificationDate,
			e.OriginalDueDate,
			e.UpdatedDueDate,
			e.OriginalAmount,
			e.UpdatedAmount,
			e.MinimumAmount,
			e.KID,
			e.CreditAccountNumber,
		})

		originalAmount += e.OriginalAmount
		updatedAmount += e.UpdatedAmount
		minimumAmount += e.MinimumAmount
	}

	t.AppendRows(rows)
	t.AppendFooter(table.Row{
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		originalAmount,
		updatedAmount,
		minimumAmount,
	})
	t.Render()

	return nil
}
