package cli

import (
	"github.com/urfave/cli/v2"
)

type transfer interface {
	Transfer(*cli.Context) error
}

func getTransfersCommand(conn transfer) *cli.Command {
	return &cli.Command{
		Name:  "transfers",
		Usage: "interact with transfers",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "to",
				Usage:    "account id to transfer to",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "from",
				Usage:    "account id to transfer from",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "message",
				Usage: "transfer message",
			},
			&cli.IntFlag{
				Name:     "amount",
				Usage:    "the amount to transfer",
				Required: true,
			},
		},
		Action: conn.Transfer,
	}
}
