package definitions

type Movable struct {
	Tangible
	health Stat
	durability Stat
	loc *Tile
}

func (m Movable) Move(dir string) bool {
	var next *Tile
	switch dir {
	case "n": next = m.loc.north
	case "s": next = m.loc.south
	case "e": next = m.loc.east
	case "w": next = m.loc.west
	default: return false
	}
	if next != nil {
		canExit := m.loc.Exit(m)
		if !canExit { return false }
		success := next.Enter(m)
		if !success { return false }
		return true
	}
	return false // we'll figure this out later!
}

type Individual interface {
	Move(dir string) bool
	String() string
}
