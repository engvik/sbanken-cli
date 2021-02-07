package main

import (
	"context"
	"log"
	"os"

	"github.com/engvik/sbanken-cli/internal/cli"
	"github.com/engvik/sbanken-cli/internal/sbanken"
	"github.com/engvik/sbanken-cli/internal/table"
)

// VERSION is the current sbanken-cli version
const VERSION string = "1.4.0"

func main() {
	ctx := context.Background()
	writer := table.NewWriter()
	writer.SetOutputMirror(os.Stdout)

	conn, err := sbanken.NewEmptyConnection(writer)
	if err != nil {
		log.Fatal(err)
	}

	app := cli.New(ctx, conn, writer, VERSION)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
