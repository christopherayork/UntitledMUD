package definitions

import (
	"errors"
	"fmt"
)

type Zone struct {
	Tangible
	grid *Grid
}

func NewZone(g Gridded, grid Grid, x, y int, coords ...map[string]map[string]bool) (*Zone, error) {
	_, ok := g.(Region)
	if !ok {
		return &Zone{}, errors.New("error: NewZone(), type Zone requires parameter p of type Region")
	}
	zone := Zone{}
	if !grid.Enter(zone, x, y) { return nil, errors.New("error: NewZone(), could not enter the zone into the grid at that location") }
	zone.grid = &grid
	zone.loc = &g
	oke := g.Enter(zone, x, y)
	if !oke { fmt.Println("NewZone() failed on Region.Enter()") }
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
