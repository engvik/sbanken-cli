package sbanken

import (
	"github.com/urfave/cli/v2"
)

// ListCards handles the cards command.
func (c *Connection) ListCards(ctx *cli.Context) error {
	cards, err := c.client.ListCards(ctx.Context)
	if err != nil {
		return err
	}

	c.writer.ListCards(cards)

	return nil
}
