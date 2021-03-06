package definitions

import (
	"errors"
)

type Zone struct {
	Tangible
	grid *Grid
}

func NewZone(grid Grid, x, y int, coords ...[][]int) (*Zone, error) {
	zone := Zone{}
	if !grid.Enter(zone, x, y) { return nil, errors.New("error: NewZone(), could not enter the zone into the grid at that location") }
	zone.grid = &grid
	if len(coords) > 0 {
		for _, set := range coords[0] {
			if len(set) < 2 { continue } // this is an improper set, and shouldn't be used
			grid.Enter(zone, set[0], set[1])
			// use every coordinate set in an Enter call for the grid, so it covers all the positions it needs to
		}
	}
	return &zone, nil
}

func (z Zone) GetLocs() [][]int {

	return make([][]int, 0, 1)
}
func (z Zone) SetCoords(x, y int) {
	z.x = x
	z.y = y
}

func (z Zone) Enter(target Display, x, y int) bool {
	if tan, ok := target.(Area); ok {
		defer z.Entered(tan)
		return true
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
