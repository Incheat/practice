package factory

type Shape interface {
	Draw() string
}

type Circle struct{}
func (Circle) Draw() string { return "Circle" }

type Square struct{}
func (Square) Draw() string { return "Square" }

func NewShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return Circle{}
	case "square":
		return Square{}
	default:
		return nil
	}
}
