package sbanken

import "github.com/urfave/cli/v2"

func (c *Connection) GetCustomer(ctx *cli.Context) error {
	customer, err := c.client.GetCustomer(ctx.Context)
	if err != nil {
		return err
	}

	includeCustomerID := ctx.Bool("customer-id")

	if !includeCustomerID {
		customer.CustomerID = ""
	}

	c.writer.GetCustomer(customer)

	return nil
}
