package definitions

import "fmt"

/*
	Maps need reworked so that each tier of things has a centralized grid that it's laid out on. It can still be referenced within it's parent's
	individual grid, it's just that calculations become needlessly complex if we have to rebuild a visual of the surroundings by joining all sorts of parents together
	by having a synced joined grid, perhaps we can minimize expensive calculations during runtime, with minimal costs on grid updates

 */

type Map struct {
	Tangible
	grid *Grid
	name string
}

func (m Map) String() string {
	return fmt.Sprintf("Type: %T, Name: %v, Description: %v, Grid: %v", m, m.name, m.description, m.grid)
	//return fmt.Sprintf("%v", m.grid.grid)
}

func (m Map) Enter(target Display, x, y int) bool {
	if tan, ok := target.(Region); ok {
		gridSuccess := m.grid.Enter(tan, x, y)
		return gridSuccess
	} else { return false }
}

func (m Map) Entered(target Display) {

}

func (m Map) Exit(target Display, x, y int) bool {
	return true
}

func (m Map) Exited(target Display) {

}

var maps = 0

func NewMap() (*Map, error) {
	maps++
	m := Map{name: fmt.Sprintf("Map%v", maps)}
	m.grid = NewGrid(&m)
	return &m, nil
}
