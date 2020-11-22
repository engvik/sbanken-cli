package sbanken

import (
	"time"

	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

// ListEfakturas handles the efakturas list command.
func (c *Connection) ListEfakturas(ctx *cli.Context) error {
	q, err := parseEfakturaListQuery(ctx)
	if err != nil {
		return err
	}

	efakturas, err := c.client.ListEfakturas(ctx.Context, q)
	if err != nil {
		return err
	}

	c.writer.ListEfakturas(efakturas)

	return nil
}

// PayEfaktura handles the efakturas pay command.
func (c *Connection) PayEfaktura(ctx *cli.Context) error {
	q := parseEfakturaPayQuery(ctx)

	if err := c.client.PayEfaktura(ctx.Context, q); err != nil {
		return err
	}

	c.writer.PayEfaktura(q)

	return nil
}

// ListNewEfakturas handles the efakturas list command with the --new option set.
func (c *Connection) ListNewEfakturas(ctx *cli.Context) error {
	q, err := parseEfakturaListQuery(ctx)
	if err != nil {
		return err
	}

	efakturas, err := c.client.ListNewEfakturas(ctx.Context, q)
	if err != nil {
		return err
	}

	c.writer.ListEfakturas(efakturas)

	return nil
}

// ReadEfaktura handles the read efakturas command.
func (c *Connection) ReadEfaktura(ctx *cli.Context) error {
	ID := ctx.String("id")

	efaktura, err := c.client.ReadEfaktura(ctx.Context, ID)
	if err != nil {
		return err
	}

	c.writer.ReadEfaktura(efaktura)

	return nil
}

func parseEfakturaListQuery(ctx *cli.Context) (*sbanken.EfakturaListQuery, error) {
	startDate := ctx.String("start-date")
	endDate := ctx.String("end-date")

	var startDateTime time.Time
	var endDateTime time.Time

	if startDate != "" {
		t, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			return nil, err
		}

		startDateTime = t
	}

	if endDate != "" {
		t, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			return nil, err
		}

		endDateTime = t
	}

	q := &sbanken.EfakturaListQuery{
		StartDate: startDateTime,
		EndDate:   endDateTime,
		Status:    ctx.String("status"),
		Index:     ctx.String("index"),
		Length:    ctx.String("length"),
	}

	return q, nil
}

func parseEfakturaPayQuery(ctx *cli.Context) *sbanken.EfakturaPayQuery {
	q := &sbanken.EfakturaPayQuery{
		ID:                   ctx.String("id"),
		AccountID:            ctx.String("account-id"),
		PayOnlyMinimumAmount: ctx.Bool("pay-minimum"),
	}

	return q
}
