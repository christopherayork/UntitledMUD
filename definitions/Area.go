package definitions

import (
	"errors"
)

type Area struct {
	Tangible
	grid *Grid
}

func NewArea(grid Grid, x, y int, coords ...[][]int) (*Area, error) {

	area := Area{}
	if !grid.Enter(area, x, y) { return nil, errors.New("error: NewArea(), could not enter the area into the grid at that location") }
	area.grid = &grid
	if len(coords) > 0 {
		for _, set := range coords[0] {
			if len(set) < 2 { continue } // this is an improper set, and shouldn't be used
			grid.Enter(area, set[0], set[1])
			// use every coordinate set in an Enter call for the grid, so it covers all the positions it needs to
		}
	}
	return &area, nil
}

func (a Area) GetLocs() [][]int {

	return make([][]int, 0, 1)
}

func (a Area) Enter(target Display, x, y int) bool {
	if tan, ok := target.(Plot); ok {
		defer a.Entered(tan)
		return true
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
