package json

import (
	"encoding/json"
	"log"

	"github.com/engvik/sbanken-go"
)

func (w *Writer) ListPayments(payments []sbanken.Payment) {
	json, err := json.MarshalIndent(payments, "", "  ")
	if err != nil {
		log.Printf("Error creating json output: %s\n", err)
		return
	}

	log.Println(string(json))
}

func (w *Writer) ReadPayment(payment sbanken.Payment) {
	json, err := json.MarshalIndent(payment, "", "  ")
	if err != nil {
		log.Printf("Error creating json output: %s\n", err)
		return
	}

	log.Println(string(json))
}
