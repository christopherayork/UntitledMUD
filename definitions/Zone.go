package definitions

import (
	"errors"
	"fmt"
)

type Zone struct {
	Tangible
	grid *Grid
	north *Zone
	south *Zone
	east *Zone
	west *Zone
}

func NewZone(g Gridded, grid Grid, x, y int, dirs ...map[string]*Zone) (*Zone, error) {
	_, ok := g.(Region)
	if !ok {
		return &Zone{}, errors.New("error: NewZone(), type Zone requires parameter p of type Region")
	}
	zone := Zone{}
	zone.grid = &grid
	zone.loc = &g
	oke := g.Enter(zone, x, y)
	if !oke { fmt.Println("NewZone() failed on Region.Enter()") }
	for i := range dirs {
		if i > 0 { break }
		if v, ok := dirs[i]["n"]; ok { zone.SetDirection("n", v) }
		if v, ok := dirs[i]["s"]; ok { zone.SetDirection("s", v) }
		if v, ok := dirs[i]["e"]; ok { zone.SetDirection("e", v) }
		if v, ok := dirs[i]["w"]; ok { zone.SetDirection("w", v) }
	}
	return &zone, nil
}

func (t Zone) SetDirection(d string, v *Zone) bool {
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

func (z Zone) Enter(target Display, x, y int) bool {
	if tan, ok := target.(Area); ok {
		gridSuccess := z.grid.Enter(tan, x, y)
		if gridSuccess { defer z.Entered(tan) }
		return gridSuccess
	}
	return false
}

func (z Zone) Entered(target Display) {
	if _, ok := target.(Individual); ok{
		z.Apply(target)
	}
}

func (z Zone) Exit(target Display, x, y int) bool {
	return true
}

func (z Zone) Exited(target Display) {

}
