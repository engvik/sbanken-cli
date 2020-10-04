package sbankencli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func getAccountsCommand() *cli.Command {
	return &cli.Command{
		Name:  "accounts",
		Usage: "interact with accounts",
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "list all accounts",
				Action: func(c *cli.Context) error {
					fmt.Println("list account")
					return nil
				},
			},
			{
				Name:  "read",
				Usage: "read a single account",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "account id to read",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Printf("read account: %s\n", c.String("id"))
					return nil
				},
			},
		},
	}
}
