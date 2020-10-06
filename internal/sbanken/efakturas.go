package sbanken

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/engvik/sbanken-go"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

func (c *Connection) ListEfakturas(cliCtx *cli.Context) error {
	ctx := context.Background()

	if err := c.ConnectClient(ctx, cliCtx); err != nil {
		return err
	}

	q, err := parseEfakturaListQuery(cliCtx)
	if err != nil {
		return err
	}

	efakturas, err := c.Client.ListEfakturas(ctx, q)
	if err != nil {
		return err
	}

	printEfakturas(efakturas)

	return nil
}

func (c *Connection) ListNewEfakturas(cliCtx *cli.Context) error {
	ctx := context.Background()

	if err := c.ConnectClient(ctx, cliCtx); err != nil {
		return err
	}

	q, err := parseEfakturaListQuery(cliCtx)
	if err != nil {
		return err
	}

	efakturas, err := c.Client.ListNewEfakturas(ctx, q)
	if err != nil {
		return err
	}

	printEfakturas(efakturas)

	return nil
}

func (c *Connection) ReadEfaktura(cliCtx *cli.Context) error {
	ctx := context.Background()
	ID := cliCtx.String("id")

	if err := c.ConnectClient(ctx, cliCtx); err != nil {
		return err
	}

	efaktura, err := c.Client.ReadEfaktura(ctx, ID)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
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

func printEfakturas(efakturas []sbanken.Efaktura) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
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

	fmt.Println()
	fmt.Println("To see all fields, use: sbanken efakturas read --id=<ID>")
	fmt.Println("Detailed fields includes: Issuer ID, Reference, Update Due Date, Updated Amount, Credit Account Number")
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

	log.Println(q)

	return q, nil
}
