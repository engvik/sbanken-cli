package cli

import (
	"github.com/urfave/cli/v2"
)

type transfer interface {
	Transfer(*cli.Context) error
}

func getTransfersCommand(conn transfer) *cli.Command {
	return &cli.Command{
		Name:    "transfers",
		Usage:   "interact with transfers",
		Aliases: []string{"tf"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "to",
				Usage:    "account id to transfer to",
				Aliases:  []string{"t"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "from",
				Usage:    "account id to transfer from",
				Aliases:  []string{"f"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "message",
				Usage:   "transfer message",
				Aliases: []string{"m"},
			},
			&cli.IntFlag{
				Name:     "amount",
				Usage:    "the amount to transfer",
				Aliases:  []string{"a"},
				Required: true,
			},
		},
		Action: conn.Transfer,
	}
}
