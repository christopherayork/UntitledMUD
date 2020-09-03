package definitions

import "errors"

type Area struct {
	Tangible
	grid *Grid
}

func (a Area) String() string {
	return a.description
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

func (a Area) Enter(target interface{}, x, y int) bool {
	if tan, ok := target.(Plot); ok {
		gridSuccess := a.grid.Enter(tan, x, y)
		return gridSuccess
	} else { return false }
}

func (a Area) Entered(target *Tangible) {

}

func (a Area) Exit(target *Tangible, x, y int) bool {
	return true
}

func (a Area) Exited(target *Tangible) {

}
