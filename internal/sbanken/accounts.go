package sbanken

import (
	"context"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

func (c *Connection) ListAccounts(cliCtx *cli.Context) error {
	ctx := context.Background()

	if err := c.ConnectClient(ctx, cliCtx); err != nil {
		return err
	}

	accounts, err := c.Client.ListAccounts(ctx)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Type", "Name", "Number", "Balance", "Available", "Credit Limit"})

	var rows []table.Row
	var balance float32
	var available float32
	var creditLimit float32

	for _, a := range accounts {
		rows = append(rows, table.Row{a.ID, a.Type, a.Name, a.Number, a.Balance, a.Available, a.CreditLimit})
		balance += a.Balance
		available += a.Available
		creditLimit += a.CreditLimit
	}

	t.AppendRows(rows)
	t.AppendFooter(table.Row{"", "", "", "", balance, available, creditLimit})
	t.Render()

	return nil
}

func (c *Connection) ReadAccount(cliCtx *cli.Context) error {
	ctx := context.Background()
	ID := cliCtx.String("id")

	if err := c.ConnectClient(ctx, cliCtx); err != nil {
		return err
	}

	account, err := c.Client.ReadAccount(ctx, ID)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendRow(table.Row{"ID", account.ID})
	t.AppendRow(table.Row{"Type", account.Type})
	t.AppendRow(table.Row{"Name", account.Name})
	t.AppendRow(table.Row{"Number", account.Number})
	t.AppendRow(table.Row{"Balance", account.Balance})
	t.AppendRow(table.Row{"Available", account.Available})
	t.AppendRow(table.Row{"Credit Limit", account.CreditLimit})

	t.Render()

	return nil
}
