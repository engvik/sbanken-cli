package cli

import (
	"context"
	"errors"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

type sbankenConn interface {
	ConnectClient(context.Context, *cli.Context) error
	ListAccounts(*cli.Context) error
	ReadAccount(*cli.Context) error
	ListCards(*cli.Context) error
	ListEfakturas(*cli.Context) error
	PayEfaktura(*cli.Context) error
	ListNewEfakturas(*cli.Context) error
	ReadEfaktura(*cli.Context) error
	ListPayments(*cli.Context) error
	ReadPayment(*cli.Context) error
	ListStandingOrders(*cli.Context) error
	ListTransactions(*cli.Context) error
	Transfer(*cli.Context) error
}

func New(ctx context.Context, conn sbankenConn, version string) *cli.App {
	flags := getGlobalFlags()

	app := &cli.App{
		Name:    "sbanken",
		Usage:   "provides an easy way to interact with your bank from the terminal",
		Version: version,
		Before: func(c *cli.Context) error {
			loadConfigFunc := altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config"))
			err := loadConfigFunc(c)

			if err != nil {
				if c.String("client-id") == "" {
					return errors.New("client-id is a required parameter")
				}

				if c.String("client-secret") == "" {
					return errors.New("client-secret is a required parameter")
				}

				if c.String("customer-id") == "" {
					return errors.New("customer-id is a required parameter")
				}
			}

			if err := conn.ConnectClient(ctx, c); err != nil {
				return err
			}

			return nil
		},
		Flags: flags,
		Commands: []*cli.Command{
			getAccountsCommand(conn),
			getCardsCommand(conn),
			getEfakturasCommand(conn),
			getPaymentsCommand(conn),
			getStandingOrdersCommand(conn),
			getTransactionsCommand(conn),
			getTransfersCommand(conn),
		},
	}

	app.EnableBashCompletion = true

	return app
}
