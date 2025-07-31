package builder

import "testing"

func TestBuilder(t *testing.T) {
	b := &HouseBuilder{}
	house := b.BuildWalls("Brick").BuildRoof("Tile").GetResult()
	if house.Walls != "Brick" || house.Roof != "Tile" {
		t.Error("Builder failed to construct correct house")
	}
}
