package sbanken

import (
	"github.com/urfave/cli/v2"
)

// ListStandingOrders handles the standing orders command.
func (c *Connection) ListStandingOrders(ctx *cli.Context) error {
	accountID, err := c.getAccountID(ctx)
	if err != nil {
		return err
	}

	standingOrders, err := c.client.ListStandingOrders(ctx.Context, accountID)
	if err != nil {
		return err
	}

	detailedOutput := ctx.Bool("details")

	c.writer.ListStandingOrders(standingOrders, detailedOutput)

	return nil
}
