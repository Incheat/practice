package adapter

type OldPrinter interface {
	PrintOld(msg string)
}

type LegacyPrinter struct{}

func (LegacyPrinter) PrintOld(msg string) {
	println("Old Printer:", msg)
}

type NewPrinter interface {
	PrintNew(msg string)
}

type Adapter struct {
	OldPrinter OldPrinter
}

func (a Adapter) PrintNew(msg string) {
	a.OldPrinter.PrintOld(msg)
}
