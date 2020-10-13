package sbanken

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

func (c *Connection) ListCards(ctx *cli.Context) error {
	cards, err := c.Client.ListCards(ctx.Context)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(c.output)
	t.AppendHeader(table.Row{"ID", "Type", "Product Code", "Number", "Account Number", "ExpiryDate", "Status", "Version Number"})

	var rows []table.Row

	for _, card := range cards {
		rows = append(rows, table.Row{card.ID, card.Type, card.ProductCode, card.Number, card.AccountNumber, card.ExpiryDate, card.Status, card.VersionNumber})
	}

	t.AppendRows(rows)
	t.Render()

	return nil
}
