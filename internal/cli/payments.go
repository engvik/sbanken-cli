package cli

import (
	"github.com/urfave/cli/v2"
)

type payments interface {
	ListPayments(*cli.Context) error
	ReadPayment(*cli.Context) error
}

func getPaymentsCommand(conn payments) *cli.Command {
	return &cli.Command{
		Name:    "payments",
		Usage:   "interact with payments",
		Aliases: []string{"p"},
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Usage:   "list all payments",
				Aliases: []string{"l"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "account id to list payments from",
						Required: true,
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
				},
				Action: conn.ListPayments,
			},
			{
				Name:    "read",
				Usage:   "read a single payment",
				Aliases: []string{"r"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "payment id",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "account-id",
						Usage:    "account id",
						Aliases:  []string{"aid"},
						Required: true,
					},
				},
				Action: conn.ReadPayment,
			},
		},
	}
}
