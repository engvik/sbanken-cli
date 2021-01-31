package sbanken

import (
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

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
	client   sbankenClient
	writer   tableWriter
	output   io.Writer
	idRegexp *regexp.Regexp
	config   *Config
}

// NewEmptyConnection returns a new connection without a connected client.
func NewEmptyConnection(tw tableWriter) (*Connection, error) {
	idRegexp, err := regexp.Compile("([0-9A-F]){32}")
	if err != nil {
		return nil, err
	}

	return &Connection{
		writer:   tw,
		output:   os.Stdout,
		idRegexp: idRegexp,
	}, nil
}

// ConnectClient sets up a connection to the sbanken client.
func (c *Connection) ConnectClient(ctx context.Context, cliCtx *cli.Context, version string) error {
	cfg := &sbanken.Config{
		ClientID:     cliCtx.String("client-id"),
		ClientSecret: cliCtx.String("client-secret"),
		CustomerID:   cliCtx.String("customer-id"),
		UserAgent:    fmt.Sprintf("sbanken-cli/%s (github.com/engvik/sbanken-cli)", version),
	}
	sClient, err := sbanken.NewClient(ctx, cfg, nil)
	if err != nil {
		return err
	}

	c.client = sClient

	return nil
}

func (c *Connection) SetConfig(cfg *Config) {
	c.config = cfg
}

func (c *Connection) getAccountID(ctx *cli.Context) (string, error) {
	ID := ctx.String("id")

	if !c.idRegexp.MatchString(ID) {
		aliasID := c.getAccountIDFromAlias(ID)
		if aliasID != "" {
			return aliasID, nil
		}

		return c.getAccountIDByName(ctx.Context, ID)
	}

	return ID, nil
}

func (c *Connection) getAccountIDWithID(ctx context.Context, ID string) (string, error) {
	if !c.idRegexp.MatchString(ID) {
		aliasID := c.getAccountIDFromAlias(ID)
		if aliasID != "" {
			return aliasID, nil
		}

		return c.getAccountIDByName(ctx, ID)
	}

	return ID, nil
}

func (c *Connection) getAccountIDByName(ctx context.Context, ID string) (string, error) {
	accounts, err := c.client.ListAccounts(ctx)
	if err != nil {
		return "", err
	}

	ID = strings.ToLower(ID)

	var found bool
	for _, a := range accounts {
		if strings.ToLower(a.Name) == ID {
			found = true
			ID = a.ID
			break
		}
	}

	if !found {
		return "", fmt.Errorf("Unknown ID: %s", ID)
	}

	return ID, nil
}

func (c *Connection) getAccountIDFromAlias(alias string) string {
	if len(c.config.AccountAliases) == 0 {
		return ""
	}

	for k, v := range c.config.AccountAliases {
		if v == alias {
			return k
		}
	}

	return ""
}
