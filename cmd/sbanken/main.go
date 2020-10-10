package main

import (
	"context"
	"log"
	"os"

	"github.com/engvik/sbanken-cli/internal/cli"
	"github.com/engvik/sbanken-cli/internal/sbanken"
)

const VERSION string = "1.0.0"

func main() {
	ctx := context.Background()
	conn := sbanken.NewEmptyConnection()
	app := cli.New(ctx, conn, VERSION)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
