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
	Light *Light
}

func (c LightOnCommand) Execute() { c.Light.On() }

type LightOffCommand struct {
	Light *Light
}

func (c LightOffCommand) Execute() { c.Light.Off() }
