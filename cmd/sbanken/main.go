package main

import (
	"context"
	"log"
	"os"

	"github.com/engvik/sbanken-cli/internal/cli"
	"github.com/engvik/sbanken-cli/internal/sbanken"
)

// VERSION is the current sbanken-cli version
const VERSION string = "1.6.0"

func main() {
	ctx := context.Background()

	conn, err := sbanken.NewEmptyConnection()
	if err != nil {
		log.Fatal(err)
	}

	app := cli.New(ctx, conn, VERSION)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
