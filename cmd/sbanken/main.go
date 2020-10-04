package main

import (
	"log"
	"os"

	"github.com/engvik/sbanken-cli/internal/sbankencli"
)

func main() {
	app := sbankencli.New()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
