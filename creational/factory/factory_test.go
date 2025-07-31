package factory

import "testing"

func TestFactory(t *testing.T) {
	shape := NewShape("circle")
	if shape.Draw() != "Circle" {
		t.Error("Expected Circle")
	}

	shape2 := NewShape("square")
	if shape2.Draw() != "Square" {
		t.Error("Expected Square")
	}
}
