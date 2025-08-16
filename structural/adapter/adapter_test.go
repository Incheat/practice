package adapter

import "testing"

func TestAdapter(t *testing.T) {
	old := LegacyPrinter{}
	adapter := Adapter{OldPrinter: old}
	adapter.PrintNew("Hello")
	// 觀察輸出，實際業務中可用 mock capture log
}
