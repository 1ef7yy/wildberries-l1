package main

import "fmt"

// паттерн адаптер
// для решения проблемы между интерфейсами которые не совместимы
// на примере этой программы
// клиент код должен работать с интерфейсом Printer, но LegacyPrinter не реализует этот интерфейс
// адаптер помогает клиенту коду работать с LegacyPrinter

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
