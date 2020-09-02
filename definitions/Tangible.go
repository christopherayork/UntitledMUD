package definitions


type Tangible struct {
	name string
	health Stat
	durability Stat
	description string
	contents []*Tangible
	loc *Gridded
	x int
	y int
}

type Gridded interface {
	Enter(target *Tangible) bool
	Entered(target *Tangible)
	Exit(target *Tangible) bool
	Exited(target *Tangible)
}

func (t Tangible) Enter(target *Tangible) bool {
	// define some rules for movement procedures
	// this will check if movement is permitted into this tangible
	return false // lets just say false for now
}


func (t Tangible) Entered(target *Tangible) {
	// if Enter() is validated as truthy, this will be called to complete the process and add addition room for behavior
	// this will return nothing as of now
}

func (t Tangible) Exit(target *Tangible) bool {
	// called to permit exiting
	return false // no allowance for now
}

func (t Tangible) Exited(target *Tangible) {
	// called on confirmation of Exit() and to complete the process plus add room for additional behavior

}

// perhaps the Move method should be migrated to a movable subsection of Tangible
func (t Tangible) Move(newloc *Tangible) bool {
	// go through the process of checking whether the new target tangible and it's containers will allow this tangible to move into it
	return false // no for now
}

// maybe later add some collision checking


// for interfacing purposes, /map will also be considered of interface type Gridded
// base type that /region, /zone, /area, /plot, /tile,
// obj, and /mob will derive from

