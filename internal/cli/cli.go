package cli

import (
	"context"

	"github.com/engvik/sbanken-cli/internal/sbanken"
	"github.com/urfave/cli/v2"
)

type sbankenConn interface {
	ConnectClient(context.Context, *sbanken.Config) error
}

func New(ctx context.Context, conn sbankenConn) *cli.App {
	app := &cli.App{
		Name:    "sbanken",
		Usage:   "interact with sbanken through the command line",
		Version: "1.0.0",
		Flags:   getGlobalFlags(),
		Action: func(c *cli.Context) error {
			err := conn.ConnectClient(ctx, &sbanken.Config{
				ClientID:     c.String("client-id"),
				ClientSecret: c.String("client-secret"),
				CustomerID:   c.String("customer-id"),
			})

			return err
		},
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
