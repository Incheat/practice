package strategy

import "testing"

func TestStrategy(t *testing.T) {
	ctx := Context{Strategy: Add{}}
	if ctx.ExecuteStrategy(3, 5) != 8 {
		t.Error("Add strategy failed")
	}

	ctx = Context{Strategy: Multiply{}}
	if ctx.ExecuteStrategy(3, 5) != 15 {
		t.Error("Multiply strategy failed")
	}
}
