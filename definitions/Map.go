package definitions

import "fmt"

type Map struct {
	Tangible
	grid *Grid
	name string
}

func (m Map) Enter(target *Tangible) bool {
	return true
}

func (m Map) Entered(target *Tangible) {

}

func (m Map) Exit(target *Tangible) bool {
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
