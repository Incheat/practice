package prototype

type Cloner interface {
	Clone() Cloner
}

type Person struct {
	Name string
}

func (p *Person) Clone() Cloner {
	copy := *p
	return &copy
}
