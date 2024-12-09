package main

import "fmt"

type Printer interface {
	Print()
}

type LegacyPrinter struct{}

func (lp *LegacyPrinter) PrintDocument() {
	fmt.Println("legacy printer is printing a document...")
}

type PrinterAdapter struct {
	legacyPrinter *LegacyPrinter
}

func (pa *PrinterAdapter) Print() {
	pa.legacyPrinter.PrintDocument()
}

func clientCode(printer Printer) {
	printer.Print()
}

func main() {
	adapter := &PrinterAdapter{legacyPrinter: &LegacyPrinter{}}
	clientCode(adapter)
}
