package sbanken

import (
	"time"

	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

// ListTransactions handles the list transactions command.
func (c *Connection) ListTransactions(ctx *cli.Context) error {
	accountID := ctx.String("id")
	detailedOutput := ctx.Bool("details")
	cardDetails := ctx.Bool("card-details")
	transactionDetails := ctx.Bool("transaction-details")
	q, err := parseTransactionListQuery(ctx)
	if err != nil {
		return err
	}

	if !c.idRegexp.MatchString(accountID) {
		var err error
		accountID, err = c.getAccountID(ctx.Context, accountID)
		if err != nil {
			return err
		}
	}

	transactions, err := c.client.ListTransactions(ctx.Context, accountID, q)
	if err != nil {
		return err
	}

	c.writer.ListTransactions(transactions, detailedOutput, cardDetails, transactionDetails)

	return nil
}

func parseTransactionListQuery(ctx *cli.Context) (*sbanken.TransactionListQuery, error) {
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

	q := &sbanken.TransactionListQuery{
		StartDate: startDateTime,
		EndDate:   endDateTime,
		Index:     ctx.String("index"),
		Length:    ctx.String("length"),
	}

	return q, nil
}
