package json

import (
	"io"
	"log"
)

type Writer struct{}

func NewWriter() *Writer {
	return &Writer{}
}

func (w *Writer) SetOutputMirror(m io.Writer) {
	log.SetFlags(0)
	log.SetOutput(m)
}

func (w *Writer) SetStyle(_ string) {}

func (w *Writer) SetColors(_ bool) {}
