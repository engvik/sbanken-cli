package sbanken

import (
	"context"
	"io"
	"os"

	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

type sbankenClient interface {
	ListAccounts(context.Context) ([]sbanken.Account, error)
	ReadAccount(context.Context, string) (sbanken.Account, error)
	ListCards(context.Context) ([]sbanken.Card, error)
	ListEfakturas(context.Context, *sbanken.EfakturaListQuery) ([]sbanken.Efaktura, error)
	PayEfaktura(context.Context, *sbanken.EfakturaPayQuery) error
	ListNewEfakturas(context.Context, *sbanken.EfakturaListQuery) ([]sbanken.Efaktura, error)
	ReadEfaktura(context.Context, string) (sbanken.Efaktura, error)
	ListPayments(context.Context, string, *sbanken.PaymentListQuery) ([]sbanken.Payment, error)
	ReadPayment(context.Context, string, string) (sbanken.Payment, error)
	ListStandingOrders(context.Context, string) ([]sbanken.StandingOrder, error)
	ListTransactions(context.Context, string, *sbanken.TransactionListQuery) ([]sbanken.Transaction, error)
	Transfer(context.Context, *sbanken.TransferQuery) error
}

type tableWriter interface {
	SetOutputMirror(io.Writer)
	ListAccounts([]sbanken.Account)
	ReadAccount(sbanken.Account)
	ListCards([]sbanken.Card)
	ListEfakturas([]sbanken.Efaktura)
	PayEfaktura(*sbanken.EfakturaPayQuery)
	ReadEfaktura(sbanken.Efaktura)
	ListPayments([]sbanken.Payment)
	ReadPayment(sbanken.Payment)
	ListStandingOrders([]sbanken.StandingOrder, bool)
	ListTransactions([]sbanken.Transaction, bool, bool, bool)
	Transfer(*sbanken.TransferQuery)
}

// Connection holds the sbanken client and the output writer.
type Connection struct {
	client sbankenClient
	writer tableWriter
	output io.Writer
}

// NewEmptyConnection returns a new connection without a connected client.
func NewEmptyConnection(tw tableWriter) *Connection {
	return &Connection{
		writer: tw,
		output: os.Stdout,
	}
}

// ConnectClient sets up a connection to the sbanken client.
func (c *Connection) ConnectClient(ctx context.Context, cliCtx *cli.Context) error {
	cfg := &sbanken.Config{
		ClientID:     cliCtx.String("client-id"),
		ClientSecret: cliCtx.String("client-secret"),
		CustomerID:   cliCtx.String("customer-id"),
	}
	sClient, err := sbanken.NewClient(ctx, cfg, nil)
	if err != nil {
		return err
	}

	c.client = sClient

	return nil
}
