package json

import (
	"encoding/json"
	"log"

	"github.com/engvik/sbanken-go"
)

func (w *Writer) Transfer(q *sbanken.TransferQuery) {
	json, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		log.Printf("Error creating json output: %s\n", err)
		return
	}

	log.Println(string(json))
}
