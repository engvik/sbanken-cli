package json

import (
	"encoding/json"
	"log"

	"github.com/engvik/sbanken-go"
)

func (w *Writer) ListTransactions(transactions []sbanken.Transaction, _ bool, _ bool, _ bool) {
	json, err := json.MarshalIndent(transactions, "", "  ")
	if err != nil {
		log.Printf("Error creating json output: %s\n", err)
		return
	}

	log.Println(string(json))
}
