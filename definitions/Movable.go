package definitions

type Movable struct {
	Tangible
	health Stat
	durability Stat
	loc *Tile
}

func (m Movable) Move(dir string) bool {
	worked := false
	switch dir {
	case "n": worked = true
	case "s": worked = true
	case "e": worked = true
	case "w": worked = true
	default: return false
	}
	if !worked { return false }
	next := m.loc.GetDir(dir)
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
