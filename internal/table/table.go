package table

import (
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Writer struct {
	table  table.Writer
	output io.Writer
}

func NewWriter() *Writer {
	t := table.NewWriter()

	return &Writer{
		table: t,
	}
}

func (w *Writer) SetOutputMirror(m io.Writer) {
	w.output = m
	w.table.SetOutputMirror(m)
}
