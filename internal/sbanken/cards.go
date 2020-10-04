package sbanken

import (
	"context"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/urfave/cli/v2"
)

func (c *Connection) ListCards(cliCtx *cli.Context) error {
	ctx := context.Background()

	if err := c.ConnectClient(ctx, cliCtx); err != nil {
		return err
	}

	cards, err := c.Client.ListCards(ctx)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Type", "Product Code", "Number", "Account Number", "ExpiryDate", "Status", "Version Number"})

	var rows []table.Row

	for _, card := range cards {
		rows = append(rows, table.Row{card.ID, card.Type, card.ProductCode, card.Number, card.AccountNumber, card.ExpiryDate, card.Status, card.VersionNumber})
	}

	t.AppendRows(rows)
	t.Render()

	return nil
}
