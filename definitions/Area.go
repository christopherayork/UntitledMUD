package definitions

import (
	"errors"
	"fmt"
)

type Area struct {
	Tangible
	grid *Grid
	north *Area
	south *Area
	east *Area
	west *Area
}

func NewArea(g Gridded, grid Grid, x, y int, coords ...map[string]map[string]bool) (*Area, error) {
	if _, ok := g.(Zone); !ok {
		return &Area{}, errors.New("error: NewArea(), argument for parameter g must be of type Zone")
	}
	area := Area{}
	if !grid.Enter(area, x, y) { return nil, errors.New("error: NewArea(), could not enter the area into the grid at that location") }
	area.grid = &grid
	area.loc = &g
	oka := g.Enter(area, x, y)
	if !oka { fmt.Println("NewArea() failed on Zone.Enter()") }
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
