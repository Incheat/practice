package builder

type House struct {
	Walls string
	Roof  string
}

type HouseBuilder struct {
	house House
}

func (b *HouseBuilder) BuildWalls(walls string) *HouseBuilder {
	b.house.Walls = walls
	return b
}
func (b *HouseBuilder) BuildRoof(roof string) *HouseBuilder {
	b.house.Roof = roof
	return b
}
func (b *HouseBuilder) GetResult() House {
	return b.house
}
