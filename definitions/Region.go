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
	if (x > 0 && y > 0)  && !grid.Enter(region, x, y) {
		return nil, errors.New("error: NewRegion(), could not enter the region into the grid at that location")
	}
	region.grid = &grid // we could set this inside g.Enter(), but we would have to test which type, and map to an interface for all types that match
	if len(coords) > 0 {
		//return nil, nil
		//fmt.Println(coords[0])
		for _, set := range coords[0] {
			//return nil, nil
			if len(set) < 2 { continue } // this is an improper set, and shouldn't be used
			grid.Enter(region, set[0], set[1])
			// use every coordinate set in an Enter call for the grid, so it covers all the positions it needs to
		}
	}
	return &region, nil
}

func (r Region) GetLocs() [][]int {

	return make([][]int, 0, 1)
}
func (r Region) SetCoords(x, y int) {
	r.x = x
	r.y = y
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