package cli

import (
	"context"

	"github.com/urfave/cli/v2"
)

type sbankenConn interface {
	ConnectClient(context.Context, *cli.Context) error
	ListAccounts(*cli.Context) error
	ReadAccount(*cli.Context) error
	ListCards(*cli.Context) error
	ListNewEfakturas(*cli.Context) error
	ReadEfaktura(*cli.Context) error
}

func New(ctx context.Context, conn sbankenConn) *cli.App {
	app := &cli.App{
		Name:    "sbanken",
		Usage:   "interact with sbanken through the command line",
		Version: "1.0.0",
		Flags:   getGlobalFlags(),
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
