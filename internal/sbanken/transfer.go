package sbanken

import (
	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

// Transfer handles the transfer command.
func (c *Connection) Transfer(ctx *cli.Context) error {
	q := parseTransferQuery(ctx)

	if err := c.client.Transfer(ctx.Context, q); err != nil {
		return err
	}

	c.writer.Transfer(q)

	return nil
}

func parseTransferQuery(ctx *cli.Context) *sbanken.TransferQuery {
	q := &sbanken.TransferQuery{
		FromAccountID: ctx.String("from"),
		ToAccountID:   ctx.String("to"),
		Message:       ctx.String("message"),
		Amount:        float32(ctx.Int("amount")),
	}

	return q
}
