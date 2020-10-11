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

func getEfakturasCommand(conn efakturas) *cli.Command {
	return &cli.Command{
		Name:    "efakturas",
		Usage:   "interact with efakturas",
		Aliases: []string{"e"},
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Usage:   "list all efakturas",
				Aliases: []string{"l"},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "new",
						Usage:   "list only new efakturas",
						Aliases: []string{"n"},
					},
					&cli.StringFlag{
						Name:    "start-date",
						Usage:   "start date to filter on (YYYY-MM-DD)",
						Aliases: []string{"sd"},
					},
					&cli.StringFlag{
						Name:    "end-date",
						Usage:   "end date to filter on (YYYY-MM-DD)",
						Aliases: []string{"ed"},
					},
					&cli.StringFlag{
						Name:    "status",
						Usage:   "status to filter on",
						Aliases: []string{"s"},
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
				Action: func(c *cli.Context) error {
					n := c.Args().Get(0)
					if n == "new" {
						return conn.ListNewEfakturas(c)

					}

					return conn.ListEfakturas(c)
				},
			},
			{
				Name:    "pay",
				Usage:   "pay efaktura",
				Aliases: []string{"p"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "efaktura id to pay",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "account-id",
						Usage:    "account id to pay from",
						Aliases:  []string{"aid"},
						Required: true,
					},
					&cli.BoolFlag{
						Name:    "pay-minimum",
						Usage:   "pay only minimum",
						Aliases: []string{"m"},
					},
				},
				Action: conn.PayEfaktura,
			},
			{
				Name:    "read",
				Usage:   "read a single efaktura",
				Aliases: []string{"r"},
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
