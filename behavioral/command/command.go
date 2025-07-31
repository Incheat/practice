package command

type Command interface {
	Execute()
}

type Light struct {
	OnState bool
}

func (l *Light) On()  { l.OnState = true }
func (l *Light) Off() { l.OnState = false }

type LightOnCommand struct {
	light *Light
}

func (c LightOnCommand) Execute() { c.light.On() }

type LightOffCommand struct {
	light *Light
}

func (c LightOffCommand) Execute() { c.light.Off() }
