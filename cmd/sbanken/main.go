package main

import (
	"context"
	"log"
	"os"

	"github.com/engvik/sbanken-cli/internal/cli"
	"github.com/engvik/sbanken-cli/internal/sbanken"
)

func main() {
	ctx := context.Background()
	conn := sbanken.NewEmptyConnection()
	app := cli.New(ctx, conn)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
