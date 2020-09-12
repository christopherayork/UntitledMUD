package definitions

import (
	"errors"
	"fmt"
)

type Region struct {
	Tangible
	north *Region
	south *Region
	east *Region
	west *Region
	parentMap *Map
	grid *Grid
}

func NewRegion(m Gridded, g Grid, x, y int, dirs ...map[string]*Region) (*Region, error) {
	if _, ok := m.(Map); !ok {
		return &Region{}, errors.New("error: NewRegion(), argument m required to be of type *Map")
	}
	region := Region{}
	region.loc = &m
	region.grid = &g
	ok := m.Enter(region, x, y)
	if !ok { fmt.Println("NewRegion() failed on Map.Enter()") }
	for i := range dirs {
		if i > 0 { break }
		if v, ok := dirs[i]["n"]; ok { region.SetDirection("n", v) }
		if v, ok := dirs[i]["s"]; ok { region.SetDirection("s", v) }
		if v, ok := dirs[i]["e"]; ok { region.SetDirection("e", v) }
		if v, ok := dirs[i]["w"]; ok { region.SetDirection("w", v) }
	}
	return &region, nil
}

func (r Region) SetDirection(d string, v *Region) bool {
	switch d {
		case "n":
			r.north = v
			v.south = &r // doubly linked
		case "s":
			r.south = v
			v.north = &r
		case "e":
			r.east = v
			v.west = &r
		case "w":
			r.west = v
			v.east = &r
		default: return false
	}
	return true
}

func (r Region) Enter(target Display, x, y int) bool {
	if tan, ok := target.(Zone); ok {
		gridSuccess := r.grid.Enter(tan, x, y)
		if gridSuccess { defer r.Entered(tan) }
		return gridSuccess
	}
	return false
}

func (r Region) Entered(target Display) {
	if _, ok := target.(Individual); ok{
		r.Apply(target)
	}
}

func (r Region) Exit(target Display, x, y int) bool {

	return true
}

func (r Region) Exited(target Display) {

}