package definitions

import (
	"errors"
)

type Region struct {
	Tangible
	parentMap *Map
	grid *Grid
}

func NewRegion(grid Grid, x, y int, coords ...[][]int) (*Region, error) {
	region := Region{}
	if !grid.Enter(region, x, y) { return nil, errors.New("error: NewRegion(), could not enter the region into the grid at that location") }
	region.grid = &grid // we could set this inside g.Enter(), but we would have to test which type, and map to an interface for all types that match
	return &region, nil
}

func (r Region) Enter(target Display, x, y int) bool {
	if tan, ok := target.(Zone); ok {
		defer r.Entered(tan)
		return true
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