package sbanken

import (
	"github.com/urfave/cli/v2"
)

// ListStandingOrders handles the standing orders command.
func (c *Connection) ListStandingOrders(ctx *cli.Context) error {
	accountID := ctx.String("id")
	detailedOutput := ctx.Bool("details")

	standingOrders, err := c.client.ListStandingOrders(ctx.Context, accountID)
	if err != nil {
		return err
	}

	c.writer.ListStandingOrders(standingOrders, detailedOutput)

	return nil
}
