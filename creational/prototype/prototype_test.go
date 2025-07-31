package prototype

import "testing"

func TestPrototype(t *testing.T) {
	p1 := &Person{Name: "John"}
	p2 := p1.Clone().(*Person)
	if p1 == p2 {
		t.Error("Expected different object instances")
	}
	if p1.Name != p2.Name {
		t.Error("Expected same data after clone")
	}
}
