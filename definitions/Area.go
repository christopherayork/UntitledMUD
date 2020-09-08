package definitions

import "errors"

type Area struct {
	Tangible
	grid *Grid
	north *Area
	south *Area
	east *Area
	west *Area
}

func NewArea(g Gridded, x, y int, dirs ...map[string]*Area) (*Area, error) {
	if _, ok := g.(Zone); !ok {
		return &Area{}, errors.New("error: NewArea(), argument for parameter g must be of type Zone")
	}
	area := Area{}
	area.grid = NewGrid(area)
	area.loc = g
	g.Enter(area, x, y)
	for i := range dirs {
		if i > 0 { break }
		if v, ok := dirs[i]["n"]; ok { area.SetDirection("n", v) }
		if v, ok := dirs[i]["s"]; ok { area.SetDirection("s", v) }
		if v, ok := dirs[i]["e"]; ok { area.SetDirection("e", v) }
		if v, ok := dirs[i]["w"]; ok { area.SetDirection("w", v) }
	}
	return &area, nil
}

func (t Area) SetDirection(d string, v *Area) bool {
	switch d {
		case "n":
			t.north = v
			v.south = &t // doubly linked
		case "s":
			t.south = v
			v.north = &t
		case "e":
			t.east = v
			v.west = &t
		case "w":
			t.west = v
			v.east = &t
		default: return false
	}
	return true
}

func (a Area) Enter(target Display, x, y int) bool {
	if tan, ok := target.(Plot); ok {
		gridSuccess := a.grid.Enter(tan, x, y)
		if gridSuccess { defer a.Entered(tan) }
		return gridSuccess
	}
	return false
}

func (a Area) Entered(target Display) {
	if _, ok := target.(Individual); ok{
		a.Apply(target)
	}
}

func (a Area) Exit(target Display, x, y int) bool {
	return true
}

func (a Area) Exited(target Display) {

}
