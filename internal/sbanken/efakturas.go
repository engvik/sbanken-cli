package sbanken

import (
	"fmt"
	"io"
	"time"

	"github.com/engvik/sbanken-go"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

// ListEfakturas handles the efakturas list command.
func (c *Connection) ListEfakturas(ctx *cli.Context) error {
	q, err := parseEfakturaListQuery(ctx)
	if err != nil {
		return err
	}

	efakturas, err := c.Client.ListEfakturas(ctx.Context, q)
	if err != nil {
		return err
	}

	printEfakturas(efakturas, c.output)

	return nil
}

// PayEfaktura handles the efakturas pay command.
func (c *Connection) PayEfaktura(ctx *cli.Context) error {
	q := parseEfakturaPayQuery(ctx)

	if err := c.Client.PayEfaktura(ctx.Context, q); err != nil {
		return err
	}

	fmt.Fprintf(c.output, "Efaktura %s paid successfully with account %s\n", q.ID, q.AccountID)

	return nil
}

// ListNewEfakturas handles the efakturas list command with the --new option set.
func (c *Connection) ListNewEfakturas(ctx *cli.Context) error {
	q, err := parseEfakturaListQuery(ctx)
	if err != nil {
		return err
	}

	efakturas, err := c.Client.ListNewEfakturas(ctx.Context, q)
	if err != nil {
		return err
	}

	printEfakturas(efakturas, c.output)

	return nil
}

// ReadEfaktura handles the read efakturas command.
func (c *Connection) ReadEfaktura(ctx *cli.Context) error {
	ID := ctx.String("id")

	efaktura, err := c.Client.ReadEfaktura(ctx.Context, ID)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(c.output)
	t.AppendRow(table.Row{"ID", efaktura.ID})
	t.AppendRow(table.Row{"Issuer ID", efaktura.IssuerID})
	t.AppendRow(table.Row{"Issuer Name", efaktura.IssuerName})
	t.AppendRow(table.Row{"Reference", efaktura.Reference})
	t.AppendRow(table.Row{"Document Type", efaktura.DocumentType})
	t.AppendRow(table.Row{"Status", efaktura.Status})
	t.AppendRow(table.Row{"Notification Date", efaktura.NotificationDate})
	t.AppendRow(table.Row{"Original Due Date", efaktura.OriginalDueDate})
	t.AppendRow(table.Row{"Update Due Date", efaktura.UpdatedDueDate})
	t.AppendRow(table.Row{"Original Amount", efaktura.OriginalAmount})
	t.AppendRow(table.Row{"Update Amount", efaktura.UpdatedAmount})
	t.AppendRow(table.Row{"Minimum Amount", efaktura.MinimumAmount})
	t.AppendRow(table.Row{"KID", efaktura.KID})
	t.AppendRow(table.Row{"Credit Account Number", efaktura.CreditAccountNumber})

	t.Render()

	return nil
}

func printEfakturas(efakturas []sbanken.Efaktura, output io.Writer) {
	t := table.NewWriter()
	t.SetOutputMirror(output)
	t.AppendHeader(table.Row{
		"ID",
		"Issuer Name",
		"Document Type",
		"Status",
		"Notification Date",
		"Original Due Date",
		"Original Amount",
		"Minimum Amoount",
		"KID",
	})

	var rows []table.Row
	var originalAmount float32
	var minimumAmount float32

	for _, e := range efakturas {
		rows = append(rows, table.Row{
			e.ID,
			e.IssuerName,
			e.DocumentType,
			e.Status,
			e.NotificationDate,
			e.OriginalDueDate,
			e.OriginalAmount,
			e.MinimumAmount,
			e.KID,
		})

		originalAmount += e.OriginalAmount
		minimumAmount += e.MinimumAmount
	}

	t.AppendRows(rows)
	t.AppendFooter(table.Row{
		"",
		"",
		"",
		"",
		"",
		"",
		originalAmount,
		minimumAmount,
	})
	t.Render()

	fmt.Fprintln(output)
	fmt.Fprintln(output, "To see all fields, use: sbanken efakturas read --id=<ID>")
	fmt.Fprintln(output, "Detailed fields includes: Issuer ID, Reference, Update Due Date, Updated Amount, Credit Account Number")
}

func parseEfakturaListQuery(ctx *cli.Context) (*sbanken.EfakturaListQuery, error) {
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

	q := &sbanken.EfakturaListQuery{
		StartDate: startDateTime,
		EndDate:   endDateTime,
		Status:    ctx.String("status"),
		Index:     ctx.String("index"),
		Length:    ctx.String("length"),
	}

	return q, nil
}

func parseEfakturaPayQuery(ctx *cli.Context) *sbanken.EfakturaPayQuery {
	q := &sbanken.EfakturaPayQuery{
		ID:                   ctx.String("id"),
		AccountID:            ctx.String("account-id"),
		PayOnlyMinimumAmount: ctx.Bool("pay-minimum"),
	}

	return q
}
