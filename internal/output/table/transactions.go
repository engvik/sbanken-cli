package table

import (
	"fmt"

	"github.com/engvik/sbanken-go"
	"github.com/jedib0t/go-pretty/v6/table"
)

func (w *Writer) ListTransactions(transactions []sbanken.Transaction, detailedOutput bool, cardDetails bool, transactionDetails bool) {
	if w.colors {
		w.setTransactionsColors()
	}

	w.table.AppendHeader(table.Row{
		"Index",
		"Text",
		"Amount",
		"Accounting Date",
		"Interest Date",
		"Transaction Type",
		"Is Reservation",
	})

	var rows []table.Row
	var amount float32

	for i, tr := range transactions {
		rows = append(rows, table.Row{
			i,
			tr.Text,
			tr.Amount,
			tr.AccountingDate,
			tr.InterestDate,
			tr.TransactionType,
			tr.IsReservation,
		})

		amount += tr.Amount
	}

	w.table.AppendRows(rows)
	w.table.AppendFooter(table.Row{"", "", amount})
	w.table.Render()
	w.table.ResetHeaders()
	w.table.ResetRows()
	w.table.ResetFooters()

	if detailedOutput {
		w.transactionPrintDetails(transactions)
		w.table.ResetHeaders()
		w.table.ResetRows()
		w.table.ResetFooters()
	}

	if cardDetails {
		w.transactionPrintCardDetails(transactions)
		w.table.ResetHeaders()
		w.table.ResetRows()
		w.table.ResetFooters()
	}

	if transactionDetails {
		w.transactionPrintTransactionDetails(transactions)
		w.table.ResetHeaders()
		w.table.ResetRows()
		w.table.ResetFooters()
	}

	if !detailedOutput && !cardDetails && !transactionDetails {
		fmt.Fprintln(w.output)
		fmt.Fprintln(w.output, "To see detailed output, use: sbanken transactions list --id=<ID> --details")
		fmt.Fprintln(w.output, "Detailed fields includes: Other Account Number, Transaction Type Text, Transaction Type Code, Reservation Type, Source")
		fmt.Fprintln(w.output)
		fmt.Fprintln(w.output, "Some transactions contains card details, to list them use: sbanken transactions list --id=<ID> --card-details")
		fmt.Fprintln(w.output, "Some transactions has more transaction details, to list them use: sbanken transactions list --id=<ID> --transaction-details")
	}
}

func (w *Writer) transactionPrintDetails(transactions []sbanken.Transaction) {
	w.table.AppendHeader(table.Row{
		"Index",
		"Other Account Number",
		"Transaction Type Text",
		"Transaction Type Code",
		"Reservation Type",
		"Source",
	})

	var rows []table.Row

	for i, tr := range transactions {
		rows = append(rows, table.Row{
			i,
			tr.OtherAccountNumber,
			tr.TransactionTypeText,
			tr.TransactionTypeCode,
			tr.ReservationType,
			tr.Source,
		})
	}

	w.table.AppendRows(rows)
	w.table.Render()
}

func (w *Writer) transactionPrintCardDetails(transactions []sbanken.Transaction) {
	w.table.AppendHeader(table.Row{
		"Index",
		"Merchant Category Code",
		"Merchant Category Description",
		"Merchant City",
		"Merchant Name",
		"Original Currency Code",
		"Purchase Date",
		"Transaction ID",
		"Currency Amount",
		"Currency Rate",
	})

	var rows []table.Row

	for i, tr := range transactions {
		if tr.CardDetailsSpecified {
			rows = append(rows, table.Row{
				i,
				tr.CardDetails.MerchantCategoryCode,
				tr.CardDetails.MerchantCategoryDescription,
				tr.CardDetails.MerchantCity,
				tr.CardDetails.MerchantName,
				tr.CardDetails.OriginalCurrencyCode,
				tr.CardDetails.PurchaseDate,
				tr.CardDetails.TransactionID,
				tr.CardDetails.CurrencyAmount,
				tr.CardDetails.CurrencyRate,
			})
		}
	}

	w.table.AppendRows(rows)
	w.table.Render()
}

func (w *Writer) transactionPrintTransactionDetails(transactions []sbanken.Transaction) {
	w.table.AppendHeader(table.Row{
		"Index",
		"ID",
		"Formatted Account Number",
		"CID",
		"Amount Description",
		"Receiver Name",
		"Payer Name",
		"Registration Date",
		"Numeric Reference",
	})

	var rows []table.Row

	for i, tr := range transactions {
		if tr.TransactionDetailSpecified {
			rows = append(rows, table.Row{
				i,
				tr.TransactionDetails.ID,
				tr.TransactionDetails.FormattedAccountNumber,
				tr.TransactionDetails.CID,
				tr.TransactionDetails.AmountDescription,
				tr.TransactionDetails.ReceiverName,
				tr.TransactionDetails.PayerName,
				tr.TransactionDetails.RegistrationDate,
				tr.TransactionDetails.NumericReference,
			})
		}
	}

	w.table.AppendRows(rows)
	w.table.Render()
}

func (w *Writer) setTransactionsColors() {
	w.table.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:              "Amount",
			Transformer:       w.colorValuesTransformer,
			TransformerFooter: w.colorValuesTransformer,
		},
		{
			Name:              "Currency Amount",
			Transformer:       w.colorValuesTransformer,
			TransformerFooter: w.colorValuesTransformer,
		},
		{
			Name:              "Currency Rate",
			Transformer:       w.colorValuesTransformer,
			TransformerFooter: w.colorValuesTransformer,
		},
	})
}
