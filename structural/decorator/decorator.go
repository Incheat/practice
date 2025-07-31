package decorator

type Notifier interface {
	Send(msg string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(msg string) {
	println("Email:", msg)
}

type SlackDecorator struct {
	Notifier Notifier
}

func (s SlackDecorator) Send(msg string) {
	s.Notifier.Send(msg)
	println("Slack:", msg)
}
