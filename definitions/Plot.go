package definitions

import (
	"errors"
)

// smallest grid available
// add extra handlers to support adding tiles to it's grid
type Plot struct {
	Tangible
	grid *Grid
}

// func (t Tangible) Apply(target Display) {}

func NewPlot(grid Grid, x, y int, coords ...[][]int) (*Plot, error) {
	plot := Plot{}
	if !grid.Enter(plot, x, y) { return nil, errors.New("error: NewPlot(), could not enter the plot into the grid at that location") }
	plot.grid = &grid
	if len(coords) > 0 {
		for _, set := range coords[0] {
			if len(set) < 2 { continue } // this is an improper set, and shouldn't be used
			grid.Enter(plot, set[0], set[1])
			// use every coordinate set in an Enter call for the grid, so it covers all the positions it needs to
		}
	}
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

func (p Plot) GetLocs() [][]int {

	return make([][]int, 0, 1)
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