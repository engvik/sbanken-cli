package cli

import (
	"github.com/urfave/cli/v2"
)

type transactions interface {
	ListTransactions(*cli.Context) error
}

func getTransactionsCommand(conn transactions) *cli.Command {
	return &cli.Command{
		Name:    "transactions",
		Usage:   "list transactions",
		Aliases: []string{"ta"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Usage:    "account id (or name) to list transactions from",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "start-date",
				Usage:   "start date to filter on",
				Aliases: []string{"sd"},
			},
			&cli.StringFlag{
				Name:    "end-date",
				Usage:   "end date to filter on",
				Aliases: []string{"ed"},
			},
			&cli.StringFlag{
				Name:    "index",
				Usage:   "index to filter on",
				Aliases: []string{"i"},
			},
			&cli.StringFlag{
				Name:    "length",
				Usage:   "length to filter on",
				Aliases: []string{"l"},
			},
			&cli.BoolFlag{
				Name:    "details",
				Usage:   "list transaction details",
				Aliases: []string{"d"},
			},
			&cli.BoolFlag{
				Name:    "card-details",
				Usage:   "list card details details if applicable",
				Aliases: []string{"cd"},
			},
			&cli.BoolFlag{
				Name:    "transaction-details",
				Usage:   "list more transaction details if applicable",
				Aliases: []string{"td"},
			},
			&cli.BoolFlag{
				Name:    "archived",
				Usage:   "lists only archived transactions",
				Aliases: []string{"a"},
			},
		},
		Action: conn.ListTransactions,
	}
}
