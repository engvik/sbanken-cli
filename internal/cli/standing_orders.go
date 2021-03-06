package cli

import (
	"github.com/urfave/cli/v2"
)

type standingOrders interface {
	ListStandingOrders(*cli.Context) error
}

func getStandingOrdersCommand(conn standingOrders) *cli.Command {
	return &cli.Command{
		Name:    "standingorders",
		Usage:   "list standing orders",
		Aliases: []string{"s"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Usage:    "account id (or name) to list payments from",
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "details",
				Usage:   "list standing orders details",
				Aliases: []string{"d"},
			},
		},
		Action: conn.ListStandingOrders,
	}
}
