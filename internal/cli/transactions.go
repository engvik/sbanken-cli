package cli

import (
	"github.com/urfave/cli/v2"
)

type transactions interface {
	ListTransactions(*cli.Context) error
}

func getTransactionsCommand(conn transactions) *cli.Command {
	return &cli.Command{
		Name:  "transactions",
		Usage: "interact with transactions",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Usage:    "account id to list payments from",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "start-date",
				Usage: "start date to filter on",
			},
			&cli.StringFlag{
				Name:  "end-date",
				Usage: "end date to filter on",
			},
			&cli.StringFlag{
				Name:  "index",
				Usage: "index to filter on",
			},
			&cli.StringFlag{
				Name:  "length",
				Usage: "length to filter on",
			},
			&cli.BoolFlag{
				Name:  "details",
				Usage: "list transaction details",
			},
			&cli.BoolFlag{
				Name:  "card-details",
				Usage: "list card details details if applicable",
			},
			&cli.BoolFlag{
				Name:  "transaction-details",
				Usage: "list more transaction details if applicable",
			},
		},
		Action: conn.ListTransactions,
	}
}
