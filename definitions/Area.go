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
	area.loc = &g
	return &area, nil
}

func (a Area) Enter(target *Tangible) bool {
	return true
}

func (a Area) Entered(target *Tangible) {

}

func (a Area) Exit(target *Tangible) bool {
	return true
}

func (a Area) Exited(target *Tangible) {

}
