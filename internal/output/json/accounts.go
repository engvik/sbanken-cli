package json

import (
	"encoding/json"
	"log"

	"github.com/engvik/sbanken-go"
)

func (w *Writer) ListAccounts(accounts []sbanken.Account) {
	json, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		log.Printf("Error creating json output: %s\n", err)
		return
	}

	log.Println(string(json))
}

func (w *Writer) ReadAccount(account sbanken.Account) {
	json, err := json.MarshalIndent(account, "", "  ")
	if err != nil {
		log.Printf("Error creating json output: %s\n", err)
		return
	}

	log.Println(string(json))
}
