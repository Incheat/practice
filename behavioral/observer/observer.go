package observer

type Observer interface {
	Update(msg string)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Add(o Observer) {
	s.observers = append(s.observers, o)
}

func (s Subject) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

type ConcreteObserver struct {
	LastMsg string
}

func (co *ConcreteObserver) Update(msg string) {
	co.LastMsg = msg
}
