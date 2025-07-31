package command

import "testing"

func TestCommand(t *testing.T) {
	light := &Light{}
	onCmd := LightOnCommand{Light: light}
	offCmd := LightOffCommand{Light: light}

	onCmd.Execute()
	if !light.OnState {
		t.Error("Light should be ON")
	}

	offCmd.Execute()
	if light.OnState {
		t.Error("Light should be OFF")
	}
}
