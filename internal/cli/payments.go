package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func getPaymentsCommand(conn sbankenConn) *cli.Command {
	return &cli.Command{
		Name:  "payments",
		Usage: "interact with payments",
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "list all payments",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "account id to list payments from",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("list payments")
					return nil
				},
			},
			{
				Name:  "read",
				Usage: "read a single payment",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "payment id",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "account-id",
						Usage:    "account id",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("read payment")
					return nil
				},
			},
		},
	}
}
