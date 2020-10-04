package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func getTransactionsCommand(conn sbankenConn) *cli.Command {
	return &cli.Command{
		Name:  "transactions",
		Usage: "interact with transactions",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Usage:    "account id to list payments from",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "start-date",
				Usage: "start date to filter on",
			},
			&cli.StringFlag{
				Name:  "end-date",
				Usage: "end date to filter on",
			},
			&cli.StringFlag{
				Name:  "index",
				Usage: "index to filter on",
			},
			&cli.StringFlag{
				Name:  "length",
				Usage: "length to filter on",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("transactions")
			return nil
		},
	}
}
