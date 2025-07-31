package decorator

import "testing"

func TestDecorator(t *testing.T) {
	email := EmailNotifier{}
	slack := SlackDecorator{notifier: email}
	slack.Send("Test Message")
}
