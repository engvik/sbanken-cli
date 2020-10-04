package cli

import (
	"github.com/urfave/cli/v2"
)

type cards interface {
	ListCards(*cli.Context) error
}

func getCardsCommand(conn sbankenConn) *cli.Command {
	return &cli.Command{
		Name:   "cards",
		Usage:  "interact with cards",
		Action: conn.ListCards,
	}
}
