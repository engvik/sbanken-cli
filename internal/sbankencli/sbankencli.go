package sbankencli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func New() *cli.App {
	app := &cli.App{
		Name:    "sbanken",
		Usage:   "interact with sbanken through the command line",
		Version: "1.0.0",
		Flags:   getGlobalFlags(),
		Action: func(c *cli.Context) error {
			clID := c.String("client-id")
			secret := c.String("client-secret")
			cuID := c.String("customer-id")

			fmt.Println(clID, secret, cuID)
			return nil
		},
		Commands: []*cli.Command{
			getAccountsCommand(),
			getCardsCommand(),
			getEfakturasCommand(),
			getPaymentsCommand(),
			getStandingOrdersCommand(),
			getTransactionsCommand(),
			getTransfersCommand(),
		},
	}

	app.EnableBashCompletion = true

	return app
}
