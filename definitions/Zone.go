package definitions

import "errors"

type Zone struct {
	Tangible
	grid *Grid
}

func NewZone(p Gridded) (*Zone, error) {
	_, ok := p.(Region)
	if !ok {
		return &Zone{}, errors.New("error: NewZone(), type Zone requires parameter p of type Region")
	}
	zone := Zone{}
	zone.grid = NewGrid(&zone)
	zone.loc = p
	return &zone, nil
}

func (z Zone) Enter(target Display, x, y int) bool {
	if tan, ok := target.(Area); ok {
		gridSuccess := z.grid.Enter(tan, x, y)
		if gridSuccess { defer z.Entered(tan) }
		return gridSuccess
	} else { return false }
}

func (z Zone) Entered(target Display) {

}

func (z Zone) Exit(target Display, x, y int) bool {
	return true
}

func (z Zone) Exited(target Display) {

}
