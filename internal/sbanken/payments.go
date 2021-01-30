package sbanken

import (
	"fmt"
	"strings"

	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

// ListPayments handles the payments list command.
func (c *Connection) ListPayments(ctx *cli.Context) error {
	accountID, err := c.getAccountID(ctx)
	if err != nil {
		return err
	}

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
	paymentID := ctx.String("id")

	accountID, err := c.getAccountID(ctx)
	if err != nil {
		return err
	}

	if !c.idRegexp.MatchString(paymentID) {
		payments, err := c.client.ListPayments(ctx.Context, accountID, nil)
		if err != nil {
			return err
		}

		paymentID = strings.ToLower(paymentID)

		var found bool
		for _, p := range payments {
			if strings.ToLower(p.Text) == paymentID {
				found = true
				paymentID = p.ID
				break
			}
		}

		if !found {
			return fmt.Errorf("Unknown ID: %s", paymentID)
		}
	}

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
