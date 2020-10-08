package cli

import (
	"github.com/urfave/cli/v2"
)

type standingOrders interface {
	ListStandingOrders(*cli.Context) error
}

func getStandingOrdersCommand(conn standingOrders) *cli.Command {
	return &cli.Command{
		Name:  "standingorders",
		Usage: "interact with standing orders",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Usage:    "account id to list payments from",
				Required: true,
			},
			&cli.BoolFlag{
				Name:  "details",
				Usage: "list standing orders details",
			},
		},
		Action: conn.ListStandingOrders,
	}
}
