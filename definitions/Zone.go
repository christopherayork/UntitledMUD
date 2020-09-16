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
	return &zone, nil
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
