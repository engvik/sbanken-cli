package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "sbanken",
		Usage:   "interact with sbanken through the command line",
		Version: "1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "client-id",
				Usage:    "the client id",
				Required: true,
				EnvVars:  []string{"SBANKEN_CLIENT_ID"},
			},
			&cli.StringFlag{
				Name:     "client-secret",
				Usage:    "the client secret",
				Required: true,
				EnvVars:  []string{"SBANKEN_CLIENT_SECRET"},
			},
			&cli.StringFlag{
				Name:     "customer-id",
				Usage:    "customer id",
				Required: true,
				EnvVars:  []string{"SBANKEN_CUSTOMER_ID"},
			},
		},
		Commands: []*cli.Command{
			{
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
			},
			{
				Name:  "cards",
				Usage: "interact with cards",
				Action: func(c *cli.Context) error {
					fmt.Println("cards")
					return nil
				},
			},
			{
				Name:  "efakturas",
				Usage: "interact with efakturas",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list all efakturas",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:  "new",
								Usage: "list only new efakturas",
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
								Name:  "status",
								Usage: "status to filter on",
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
							n := c.Args().Get(0)
							if n == "new" {
								fmt.Println("list new efakturas")
								return nil
							}

							fmt.Println("list all efakturas")
							return nil
						},
					},
					{
						Name:  "pay",
						Usage: "pay efaktura",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "id",
								Usage:    "efaktura id to pay",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "account-id",
								Usage:    "account id to pay from",
								Required: true,
							},
							&cli.BoolFlag{
								Name:     "pay-minimum",
								Usage:    "pay only minimum",
								Required: true,
							},
						},
						Action: func(c *cli.Context) error {
							fmt.Println("pay efakturas")
							return nil
						},
					},
					{
						Name:  "read",
						Usage: "read a single efaktura",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "id",
								Usage:    "efaktura id to read",
								Required: true,
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
							fmt.Println("read efaktura")
							return nil
						},
					},
				},
			},
			{
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
			},
			{
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
			},
			{
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
			},
			{
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
			},
		},
	}

	app.EnableBashCompletion = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
