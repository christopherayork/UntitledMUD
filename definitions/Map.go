package definitions

import "fmt"

type Map struct {
	Tangible
	grid *Grid
	name string
}

func (m Map) Enter(target interface{}, x, y int) bool {
	if tan, ok := target.(Region); ok {
		gridSuccess := m.grid.Enter(tan, x, y)
		return gridSuccess
	} else { return false }
}

func (m Map) Entered(target *Tangible) {

}

func (m Map) Exit(target *Tangible, x, y int) bool {
	return true
}

func (m Map) Exited(target *Tangible) {

}

var maps = 0

func NewMap() (*Map, error) {
	maps++
	m := Map{name: fmt.Sprintf("Map%v", maps)}
	grid := NewGrid(&m)
	m.grid = grid
	return &m, nil
}
