package definitions

import "errors"

type Region struct {
	Tangible
	grid *Grid
}

func NewRegion(m Gridded) (*Region, error) {
	if _, ok := m.(Map); !ok {
		return &Region{}, errors.New("error: NewRegion(), argument m required to be of type *Map")
	}
	region := Region{}
	region.grid = NewGrid(&region)
	region.loc = m
	return &region, nil
}

func (r Region) Enter(target Display, x, y int) bool {
	if tan, ok := target.(Zone); ok {
		gridSuccess := r.grid.Enter(tan, x, y)
		if gridSuccess { defer r.Entered(tan) }
		return gridSuccess
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