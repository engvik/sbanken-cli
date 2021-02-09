package cli

import "github.com/urfave/cli/v2"

type customers interface {
	GetCustomer(*cli.Context) error
}

func getCustomerCommand(conn customers) *cli.Command {
	return &cli.Command{
		Name:    "customer",
		Usage:   "get customer data",
		Aliases: []string{"cu"},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "customer-id",
				Usage: "include customer id in output",
			},
		},
		Action: conn.GetCustomer,
	}
}
