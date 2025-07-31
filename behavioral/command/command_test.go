package command

import "testing"

func TestCommand(t *testing.T) {
	light := &Light{}
	onCmd := LightOnCommand{light: light}
	offCmd := LightOffCommand{light: light}

	onCmd.Execute()
	if !light.OnState {
		t.Error("Light should be ON")
	}

	offCmd.Execute()
	if light.OnState {
		t.Error("Light should be OFF")
	}
}
