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
}

// func (t Tangible) Apply(target Display) {}

func NewPlot(g Gridded, grid Grid, x, y int, coords ...map[string]map[string]bool) (*Plot, error) {
	if _, ok := g.(Area); !ok {
		return &Plot{}, errors.New("error: NewPlot(), argument for parameter g must be of type Area")
	}
	plot := Plot{}
	if !grid.Enter(plot, x, y) { return nil, errors.New("error: NewPlot(), could not enter the plot into the grid at that location") }
	plot.grid = &grid
	plot.loc = &g
	okp := g.Enter(plot, x, y)
	if !okp { fmt.Println("NewPlot() failed on Area.Enter()") }
	return &plot, nil
}

func (p Plot) Enter(target Display, x, y int) bool {
	if tan, ok := target.(Tile); ok {
		defer p.Entered(tan)
		return true
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