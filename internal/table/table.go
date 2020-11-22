package table

import (
	"io"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type Writer struct {
	table                  table.Writer
	output                 io.Writer
	colorValuesTransformer text.Transformer
	colors                 bool
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

func (w *Writer) SetStyle(style string) {
	style = strings.ToLower(style)
	switch style {
	case "bold":
		w.table.SetStyle(table.StyleBold)
	case "colored-bright":
		w.table.SetStyle(table.StyleColoredBright)
	case "colored-dark":
		w.table.SetStyle(table.StyleColoredDark)
	case "colored-black-on-blue-white":
		w.table.SetStyle(table.StyleColoredBlackOnBlueWhite)
	case "colored-black-on-cyan-white":
		w.table.SetStyle(table.StyleColoredBlackOnCyanWhite)
	case "colored-black-on-green-white":
		w.table.SetStyle(table.StyleColoredBlackOnGreenWhite)
	case "colored-black-on-magenta-white":
		w.table.SetStyle(table.StyleColoredBlackOnMagentaWhite)
	case "colored-black-on-yellow-white":
		w.table.SetStyle(table.StyleColoredBlackOnYellowWhite)
	case "colored-black-on-red-white":
		w.table.SetStyle(table.StyleColoredBlackOnRedWhite)
	case "colored-blue-white-on-black":
		w.table.SetStyle(table.StyleColoredBlueWhiteOnBlack)
	case "colored-cyan-white-on-black":
		w.table.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	case "colored-green-white-on-black":
		w.table.SetStyle(table.StyleColoredGreenWhiteOnBlack)
	case "colored-magenta-white-on-black":
		w.table.SetStyle(table.StyleColoredMagentaWhiteOnBlack)
	case "colored-red-white-on-black":
		w.table.SetStyle(table.StyleColoredRedWhiteOnBlack)
	case "colored-Yellow-white-on-black":
		w.table.SetStyle(table.StyleColoredYellowWhiteOnBlack)
	case "double":
		w.table.SetStyle(table.StyleDouble)
	case "light":
		w.table.SetStyle(table.StyleLight)
	case "rounded":
		w.table.SetStyle(table.StyleRounded)
	}
}

func (w *Writer) SetColors(colors bool) {
	w.colors = colors
	w.setColorValuesTranformer()
}

func (w *Writer) setColorValuesTranformer() {
	w.colorValuesTransformer = text.Transformer(func(val interface{}) string {
		value := val.(float32)
		if value < 0 {
			return text.Color.Sprint(text.FgRed, val)
		}

		if value == 0 {
			return "0"
		}

		return text.Color.Sprint(text.FgGreen, val)
	})
}
