package decorator

import "testing"

func TestDecorator(t *testing.T) {
	email := EmailNotifier{}
	slack := SlackDecorator{Notifier: email}
	slack.Send("Test Message")
}
