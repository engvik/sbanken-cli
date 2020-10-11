package cli

import (
	"context"

	"github.com/urfave/cli/v2"
)

type accounts interface {
	ConnectClient(context.Context, *cli.Context) error
	ListAccounts(*cli.Context) error
	ReadAccount(*cli.Context) error
}

func getAccountsCommand(conn accounts) *cli.Command {
	return &cli.Command{
		Name:    "accounts",
		Usage:   "interact with accounts",
		Aliases: []string{"a"},
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Usage:   "list all accounts",
				Aliases: []string{"l"},
				Action:  conn.ListAccounts,
			},
			{
				Name:    "read",
				Usage:   "read a single account",
				Aliases: []string{"r"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "account id to read",
						Required: true,
					},
				},
				Action: conn.ReadAccount,
			},
		},
	}
}
