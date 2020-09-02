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
	region.loc = &m
	return &region, nil
}

func (r Region) Enter(target *Tangible) bool {

	return true
}

func (r Region) Entered(target *Tangible) {

}

func (r Region) Exit(target *Tangible) bool {

	return true
}

func (r Region) Exited(target *Tangible) {

}