package json

import (
	"encoding/json"
	"log"

	"github.com/engvik/sbanken-go"
)

func (w *Writer) ListEfakturas(efakturas []sbanken.Efaktura) {
	json, err := json.MarshalIndent(efakturas, "", "  ")
	if err != nil {
		log.Printf("Error creating json output: %s\n", err)
		return
	}

	log.Println(string(json))
}

func (w *Writer) PayEfaktura(q *sbanken.EfakturaPayQuery) {
	json, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		log.Printf("Error creating json output: %s\n", err)
		return
	}

	log.Println(string(json))
}

func (w *Writer) ReadEfaktura(efaktura sbanken.Efaktura) {
	json, err := json.MarshalIndent(efaktura, "", "  ")
	if err != nil {
		log.Printf("Error creating json output: %s\n", err)
		return
	}

	log.Println(string(json))
}
