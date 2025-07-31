package abstractfactory

import "testing"

func TestAbstractFactory_Windows(t *testing.T) {
	factory := GetFactory("windows")
	if factory == nil {
		t.Fatal("Expected Windows factory, got nil")
	}

	btn := factory.CreateButton()
	if btn.Paint() != "Render Windows Button" {
		t.Errorf("Expected Windows Button, got %s", btn.Paint())
	}

	chk := factory.CreateCheckbox()
	if chk.Paint() != "Render Windows Checkbox" {
		t.Errorf("Expected Windows Checkbox, got %s", chk.Paint())
	}
}

func TestAbstractFactory_Mac(t *testing.T) {
	factory := GetFactory("mac")
	if factory == nil {
		t.Fatal("Expected Mac factory, got nil")
	}

	btn := factory.CreateButton()
	if btn.Paint() != "Render Mac Button" {
		t.Errorf("Expected Mac Button, got %s", btn.Paint())
	}

	chk := factory.CreateCheckbox()
	if chk.Paint() != "Render Mac Checkbox" {
		t.Errorf("Expected Mac Checkbox, got %s", chk.Paint())
	}
}
