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

type Connection struct {
	Client sbankenClient
	output io.Writer
}

func NewEmptyConnection() *Connection {
	return &Connection{
		output: os.Stdout,
	}
}

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

	c.Client = sClient

	return nil
}
