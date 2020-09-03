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

func (r Region) Enter(target interface{}, x, y int) bool {
	if tan, ok := target.(Zone); ok {
		gridSuccess := r.grid.Enter(tan, x, y)
		return gridSuccess
	} else { return false }
}

func (r Region) Entered(target *Tangible) {

}

func (r Region) Exit(target *Tangible, x, y int) bool {

	return true
}

func (r Region) Exited(target *Tangible) {

}