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
	return &area, nil
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
