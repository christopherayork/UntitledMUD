package definitions

type Movable struct {
	Tangible
	health Stat
	durability Stat
}

func (m Movable) Move(dir string) bool {

	return true // we'll figure this out later!
}

type Individual interface {
	Move(dir string) bool
}
