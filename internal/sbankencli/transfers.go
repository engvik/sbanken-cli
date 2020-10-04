package sbankencli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func getTransfersCommand() *cli.Command {
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
				Name:     "message",
				Usage:    "transfer message",
				Required: true,
			},
			&cli.Float64Flag{
				Name:     "amount",
				Usage:    "the amount to transfer",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("transfers")
			return nil
		},
	}
}
