package cli

import (
	"github.com/urfave/cli/v2"
)

type efakturas interface {
	ListEfakturas(*cli.Context) error
	PayEfaktura(*cli.Context) error
	ListNewEfakturas(*cli.Context) error
	ReadEfaktura(*cli.Context) error
}

func getEfakturasCommand(conn sbankenConn) *cli.Command {
	return &cli.Command{
		Name:  "efakturas",
		Usage: "interact with efakturas",
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "list all efakturas",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "new",
						Usage: "list only new efakturas",
					},
					&cli.StringFlag{
						Name:  "start-date",
						Usage: "start date to filter on (YYYY-MM-DD)",
					},
					&cli.StringFlag{
						Name:  "end-date",
						Usage: "end date to filter on (YYYY-MM-DD)",
					},
					&cli.StringFlag{
						Name:  "status",
						Usage: "status to filter on",
					},
					&cli.StringFlag{
						Name:  "index",
						Usage: "index to filter on",
					},
					&cli.StringFlag{
						Name:  "length",
						Usage: "length to filter on",
					},
				},
				Action: func(c *cli.Context) error {
					n := c.Args().Get(0)
					if n == "new" {
						return conn.ListNewEfakturas(c)

					}

					return conn.ListEfakturas(c)
				},
			},
			{
				Name:  "pay",
				Usage: "pay efaktura",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "efaktura id to pay",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "account-id",
						Usage:    "account id to pay from",
						Required: true,
					},
					&cli.BoolFlag{
						Name:  "pay-minimum",
						Usage: "pay only minimum",
					},
				},
				Action: conn.PayEfaktura,
			},
			{
				Name:  "read",
				Usage: "read a single efaktura",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "efaktura id to read",
						Required: true,
					},
				},
				Action: conn.ReadEfaktura,
			},
		},
	}
}
