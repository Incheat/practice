package decorator

type Notifier interface {
	Send(msg string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(msg string) {
	println("Email:", msg)
}

type SlackDecorator struct {
	notifier Notifier
}

func (s SlackDecorator) Send(msg string) {
	s.notifier.Send(msg)
	println("Slack:", msg)
}
