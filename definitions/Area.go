package definitions

import "errors"

type Area struct {
	Tangible
	grid *Grid
}

func NewArea(g Gridded) (*Area, error) {
	if _, ok := g.(Zone); !ok {
		return &Area{}, errors.New("error: NewArea(), argument for parameter g must be of type Zone")
	}
	area := Area{}
	area.grid = NewGrid(area)
	area.loc = g
	return &area, nil
}

func (a Area) Enter(target Display, x, y int) bool {
	if tan, ok := target.(Plot); ok {
		gridSuccess := a.grid.Enter(tan, x, y)
		if gridSuccess { defer a.Entered(tan) }
		return gridSuccess
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
