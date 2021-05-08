package sbanken

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/engvik/sbanken-cli/internal/output/json"
	"github.com/engvik/sbanken-cli/internal/output/table"
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
	GetCustomer(context.Context) (sbanken.Customer, error)
}

type outputWriter interface {
	SetOutputMirror(io.Writer)
	SetStyle(string)
	SetColors(bool)
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
	GetCustomer(sbanken.Customer)
}

// Connection holds the sbanken client and the output writer.
type Connection struct {
	client   sbankenClient
	writer   outputWriter
	idRegexp *regexp.Regexp
	config   *Config
}

// NewEmptyConnection returns a new connection without a connected client.
func NewEmptyConnection() (*Connection, error) {
	idRegexp, err := regexp.Compile("([0-9A-F]){32}")
	if err != nil {
		return nil, err
	}

	return &Connection{
		idRegexp: idRegexp,
		config:   &Config{},
	}, nil
}

// ConnectClient sets up a connection to the sbanken client.
func (c *Connection) ConnectClient(ctx context.Context, cliCtx *cli.Context, version string) error {
	cfg := &sbanken.Config{
		ClientID:     cliCtx.String("client-id"),
		ClientSecret: cliCtx.String("client-secret"),
		UserAgent:    fmt.Sprintf("sbanken-cli/%s (github.com/engvik/sbanken-cli)", version),
	}

	httpClient := &http.Client{
		Timeout: time.Second * time.Duration(cliCtx.Int("http-timeout")),
	}

	sClient, err := sbanken.NewClient(ctx, cfg, httpClient)
	if err != nil {
		return err
	}

	c.client = sClient

	return nil
}

func (c *Connection) SetConfig(cfg *Config) {
	c.config = cfg
}

func (c *Connection) SetWriter(ctx *cli.Context) {
	switch ctx.String("output") {
	case "json":
		c.writer = json.NewWriter()
		c.writer.SetOutputMirror(os.Stdout)
	case "table":
		fallthrough
	default:
		c.writer = table.NewWriter()
		c.writer.SetOutputMirror(os.Stdout)
		c.writer.SetStyle(ctx.String("style"))
		c.writer.SetColors(ctx.Bool("colors"))
	}
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
