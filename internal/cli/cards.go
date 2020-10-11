package cli

import (
	"github.com/urfave/cli/v2"
)

type cards interface {
	ListCards(*cli.Context) error
}

func getCardsCommand(conn cards) *cli.Command {
	return &cli.Command{
		Name:    "cards",
		Usage:   "list cards",
		Aliases: []string{"c"},
		Action:  conn.ListCards,
	}
}
