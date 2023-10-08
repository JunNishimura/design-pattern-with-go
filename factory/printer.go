package factory

import (
	"fmt"
	"io"
	"os"
)

type Printer interface {
	Output(string)
}

type PrinterType int

const (
	Stdout PrinterType = iota
	File
)

func NewPrinter(pt PrinterType) Printer {
	switch pt {
	case Stdout:
		return &StdoutPrinter{}
	case File:
		f, _ := os.Create("log.txt")
		return &FilePrinter{
			w: f,
		}
	default:
		return &StdoutPrinter{}
	}
}

type StdoutPrinter struct{}

func (p *StdoutPrinter) Output(message string) {
	fmt.Println(message)
}

type FilePrinter struct {
	w io.Writer
}

func (p *FilePrinter) Output(message string) {
	fmt.Fprintln(p.w, message)
}
