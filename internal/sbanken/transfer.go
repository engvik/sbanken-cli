package sbanken

import (
	"fmt"

	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

func (c *Connection) Transfer(ctx *cli.Context) error {
	q := parseTransferQuery(ctx)

	if err := c.Client.Transfer(ctx.Context, q); err != nil {
		return err
	}

	if q.Message != "" {
		fmt.Printf("%f successfully transfered from %s to %s: %s\n", q.Amount, q.FromAccountID, q.ToAccountID, q.Message)
	} else {
		fmt.Printf("%f successfully transfered from %s to %s\n", q.Amount, q.FromAccountID, q.ToAccountID)
	}

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
