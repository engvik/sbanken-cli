package sbanken

import (
	"fmt"
	"io"
	"time"

	"github.com/engvik/sbanken-go"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

func (c *Connection) ListTransactions(ctx *cli.Context) error {
	accountID := ctx.String("id")
	detailedOutput := ctx.Bool("details")
	cardDetails := ctx.Bool("card-details")
	transactionDetails := ctx.Bool("transaction-details")
	q, err := parseTransactionListQuery(ctx)
	if err != nil {
		return err
	}

	transactions, err := c.Client.ListTransactions(ctx.Context, accountID, q)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(c.output)
	t.AppendHeader(table.Row{
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

	t.AppendRows(rows)
	t.AppendFooter(table.Row{"", "", amount})
	t.Render()

	if detailedOutput {
		printDetails(transactions, c.output)
	}

	if cardDetails {
		printCardDetails(transactions, c.output)
	}

	if transactionDetails {
		printTransactionDetails(transactions, c.output)
	}

	if !detailedOutput && !cardDetails && !transactionDetails {
		fmt.Fprintln(c.output)
		fmt.Fprintln(c.output, "To see detailed output, use: sbanken transactions list --id=<ID> --details")
		fmt.Fprintln(c.output, "Detailed fields includes: Other Account Number, Transaction Type Text, Transaction Type Code, Reservation Type, Source")
		fmt.Fprintln(c.output)
		fmt.Fprintln(c.output, "Some transactions contains card details, to list them use: sbanken transactions list --id=<ID> --card-details")
		fmt.Fprintln(c.output, "Some transactions has more transaction details, to list them use: sbanken transactions list --id=<ID> --transaction-details")
	}

	return nil
}

func parseTransactionListQuery(ctx *cli.Context) (*sbanken.TransactionListQuery, error) {
	startDate := ctx.String("start-date")
	endDate := ctx.String("end-date")

	var startDateTime time.Time
	var endDateTime time.Time

	if startDate != "" {
		t, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			return nil, err
		}

		startDateTime = t
	}

	if endDate != "" {
		t, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			return nil, err
		}

		endDateTime = t
	}

	q := &sbanken.TransactionListQuery{
		StartDate: startDateTime,
		EndDate:   endDateTime,
		Index:     ctx.String("index"),
		Length:    ctx.String("length"),
	}

	return q, nil
}

func printDetails(transactions []sbanken.Transaction, output io.Writer) {
	td := table.NewWriter()
	td.SetOutputMirror(output)
	td.AppendHeader(table.Row{
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
	td.AppendRows(rows)
	td.Render()
}

func printCardDetails(transactions []sbanken.Transaction, output io.Writer) {
	td := table.NewWriter()
	td.SetOutputMirror(output)
	td.AppendHeader(table.Row{
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
	td.AppendRows(rows)
	td.Render()
}

func printTransactionDetails(transactions []sbanken.Transaction, output io.Writer) {
	td := table.NewWriter()
	td.SetOutputMirror(output)
	td.AppendHeader(table.Row{
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

	td.AppendRows(rows)
	td.Render()
}
