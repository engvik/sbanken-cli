package sbanken

import (
	"github.com/urfave/cli/v2"
)

// ListAccounts handles the accounts list command.
func (c *Connection) ListAccounts(ctx *cli.Context) error {
	accounts, err := c.client.ListAccounts(ctx.Context)
	if err != nil {
		return err
	}

	c.writer.ListAccounts(accounts)

	return nil
}

// ReadAccount handles the accounts read command.
func (c *Connection) ReadAccount(ctx *cli.Context) error {
	ID := ctx.String("id")

	account, err := c.client.ReadAccount(ctx.Context, ID)
	if err != nil {
		return err
	}

	c.writer.ReadAccount(account)

	return nil
}
