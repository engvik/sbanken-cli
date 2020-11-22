package sbanken

import (
	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

// ListPayments handles the payments list command.
func (c *Connection) ListPayments(ctx *cli.Context) error {
	accountID := ctx.String("id")
	q := parsePaymentListQuery(ctx)

	payments, err := c.client.ListPayments(ctx.Context, accountID, q)
	if err != nil {
		return err
	}

	c.writer.ListPayments(payments)

	return nil
}

// ReadPayment handles the payments read command.
func (c *Connection) ReadPayment(ctx *cli.Context) error {
	accountID := ctx.String("account-id")
	paymentID := ctx.String("id")

	payment, err := c.client.ReadPayment(ctx.Context, accountID, paymentID)
	if err != nil {
		return err
	}

	c.writer.ReadPayment(payment)

	return nil
}

func parsePaymentListQuery(ctx *cli.Context) *sbanken.PaymentListQuery {
	q := &sbanken.PaymentListQuery{
		Index:  ctx.String("index"),
		Length: ctx.String("length"),
	}

	return q
}
