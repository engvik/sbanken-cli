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
		Action:  conn.GetCustomer,
	}
}
