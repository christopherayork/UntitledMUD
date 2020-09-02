package definitions

type Tangible struct {
	name string
	health Stat
	durability Stat
	description string
	contents []*Tangible
	loc *Tangible
	x int
	y int
}

type Grid struct {
	grid map[string]map[string]*Tangible
	parent *Tangible
}

func (g Grid) Enter(target *Tangible, x int, y int) bool {
	success := false
	if _, ok := g.grid[string(x)]; !ok {
		g.grid[string(x)] = make(map[string]*Tangible)
	}
	g.grid[string(x)][string(y)] = target
	target.loc = g.parent
	target.x = x
	target.y = y
	success = true
	defer g.Entered(target)
	return success
}

func (g Grid) Entered(target *Tangible) {

}

func (g Grid) Exit(target *Tangible) bool {
	success := false
	if _, ok := g.grid[string(target.x)]; ok {
		if _, ok2 := g.grid[string(target.x)][string(target.y)]; ok2 {
			g.grid[string(target.x)][string(target.y)] = nil
		}
	}
	success = true
	target.loc = nil
	target.x = 0
	target.y = 0
	defer g.Exited(target)
	return success
}

func (g Grid) Exited(target *Tangible) {

}

// probably need to setup some interfaces for these types and methods

func (t Tangible) Enter(target Tangible) bool {
	// define some rules for movement procedures
	// this will check if movement is permitted into this tangible
	return false // lets just say false for now
}


func (t Tangible) Entered(target Tangible) {
	// if Enter() is validated as truthy, this will be called to complete the process and add addition room for behavior
	// this will return nothing as of now
}

func (t Tangible) Exit(target Tangible) bool {
	// called to permit exiting
	return false // no allowance for now
}

func (t Tangible) Exited(target Tangible) {
	// called on confirmation of Exit() and to complete the process plus add room for additional behavior

}

// perhaps the Move method should be migrated to a movable subsection of Tangible
func (t Tangible) Move(newloc Tangible) bool {
	// go through the process of checking whether the new target tangible and it's containers will allow this tangible to move into it
	return false // no for now
}

// maybe later add some collision checking



// base type that /region, /zone, /area, /plot, /tile,
// obj, and /mob will derive from

