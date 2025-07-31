package strategy

type Strategy interface {
	Execute(a, b int) int
}

type Add struct{}
func (Add) Execute(a, b int) int { return a + b }

type Multiply struct{}
func (Multiply) Execute(a, b int) int { return a * b }

type Context struct {
	Strategy Strategy
}

func (c Context) ExecuteStrategy(a, b int) int {
	return c.Strategy.Execute(a, b)
}
