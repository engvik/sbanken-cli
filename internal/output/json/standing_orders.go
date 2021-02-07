package json

import (
	"encoding/json"
	"log"

	"github.com/engvik/sbanken-go"
)

func (w *Writer) ListStandingOrders(standingOrders []sbanken.StandingOrder, _ bool) {
	json, err := json.MarshalIndent(standingOrders, "", "  ")
	if err != nil {
		log.Printf("Error creating json output: %s\n", err)
		return
	}

	log.Println(string(json))
}
