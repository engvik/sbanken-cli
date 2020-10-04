package sbankencli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func getCardsCommand() *cli.Command {
	return &cli.Command{
		Name:  "cards",
		Usage: "interact with cards",
		Action: func(c *cli.Context) error {
			fmt.Println("cards")
			return nil
		},
	}
}
