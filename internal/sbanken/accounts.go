package sbanken

import (
	"context"
	"fmt"
	"strings"

	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

func (c *Connection) ListAccounts(cliCtx *cli.Context) error {
	ctx := context.Background()

	if err := c.ConnectClient(ctx, cliCtx); err != nil {
		return err
	}

	accounts, err := c.Client.ListAccounts(ctx)
	if err != nil {
		return err
	}

	printAccountHeader()

	for _, a := range accounts {
		printAccount(a)
	}

	return nil
}

func (c *Connection) ReadAccount(cliCtx *cli.Context) error {
	ctx := context.Background()
	ID := cliCtx.String("id")

	if err := c.ConnectClient(ctx, cliCtx); err != nil {
		return err
	}

	account, err := c.Client.ReadAccount(ctx, ID)
	if err != nil {
		return err
	}

	printAccountHeader()
	printAccount(account)

	return nil

}

func printAccountHeader() {
	header := "ID\tName\tNumber\tType\tBalance\tAvailable\tCreditLimit"
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))
}

func printAccount(a sbanken.Account) {
	formattedAccount := fmt.Sprintf(
		"%s\t%s\t%s\t%s\t%.2f\t%.2f\t%.2f",
		a.ID,
		a.Name,
		a.Number,
		a.Type,
		a.Balance,
		a.Available,
		a.CreditLimit,
	)

	fmt.Println(formattedAccount)
}
