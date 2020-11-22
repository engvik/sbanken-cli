package table

import (
	"fmt"

	"github.com/engvik/sbanken-go"
)

func (w *Writer) Transfer(q *sbanken.TransferQuery) {
	if q.Message != "" {
		fmt.Fprintf(w.output, "%f successfully transferred from %s to %s: %s\n", q.Amount, q.FromAccountID, q.ToAccountID, q.Message)
	} else {
		fmt.Fprintf(w.output, "%f successfully transferred from %s to %s\n", q.Amount, q.FromAccountID, q.ToAccountID)
	}
}
