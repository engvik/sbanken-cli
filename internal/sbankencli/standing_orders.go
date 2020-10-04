package sbankencli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func getStandingOrdersCommand() *cli.Command {
	return &cli.Command{
		Name:  "standingorders",
		Usage: "interact with standing orders",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Usage:    "account id to list payments from",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("standingorders")
			return nil
		},
	}
}
