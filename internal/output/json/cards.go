package json

import (
	"encoding/json"
	"log"

	"github.com/engvik/sbanken-go"
)

func (w *Writer) ListCards(cards []sbanken.Card) {
	json, err := json.MarshalIndent(cards, "", "  ")
	if err != nil {
		log.Printf("Error creating json output: %s\n", err)
		return
	}

	log.Println(string(json))
}
