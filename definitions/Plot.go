package definitions

import (
	"errors"
	"fmt"
)

// smallest grid available
// add extra handlers to support adding tiles to it's grid
type Plot struct {
	Tangible
	grid *Grid
	north *Plot
	south *Plot
	east *Plot
	west *Plot
}

// func (t Tangible) Apply(target Display) {}

func NewPlot(g Gridded, x, y int, dirs ...map[string]*Plot) (*Plot, error) {
	if _, ok := g.(Area); !ok {
		return &Plot{}, errors.New("error: NewPlot(), argument for parameter g must be of type Area")
	}
	plot := Plot{}
	plot.grid = NewGrid(plot)
	plot.loc = g
	okp := g.Enter(plot, x, y)
	if !okp { fmt.Println("NewPlot() failed on Area.Enter()") }
	for i := range dirs {
		if i > 0 { break }
		if v, ok := dirs[i]["n"]; ok { plot.SetDirection("n", v) }
		if v, ok := dirs[i]["s"]; ok { plot.SetDirection("s", v) }
		if v, ok := dirs[i]["e"]; ok { plot.SetDirection("e", v) }
		if v, ok := dirs[i]["w"]; ok { plot.SetDirection("w", v) }
	}
	return &plot, nil
}

func (t Plot) SetDirection(d string, v *Plot) bool {
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

func (p Plot) Enter(target Display, x, y int) bool {
	if tan, ok := target.(Tile); ok {
		gridSuccess := p.grid.Enter(&tan, x, y)
		if gridSuccess { defer p.Entered(tan) }
		return gridSuccess
	}
	if ind, ok := target.(Individual); ok {
		// output description to target
		p.Apply(ind)
		return true
		// currently this only outputs messages from the Plot, and doesn't go up the chain. We'll add that later.
	}
	return false
}

func (p Plot) Entered(target Display) {
	if _, ok := target.(Individual); ok{
		p.Apply(target)
	}
}

func (p Plot) Exit(target Display, x, y int) bool {
	return true
}

func (p Plot) Exited(target Display) {

}