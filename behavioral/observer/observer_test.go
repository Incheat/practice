package observer

import "testing"

func TestObserver(t *testing.T) {
	subject := &Subject{}
	obs := &ConcreteObserver{}
	subject.Add(obs)

	subject.Notify("Hello")

	if obs.LastMsg != "Hello" {
		t.Error("Observer did not receive message")
	}
}
