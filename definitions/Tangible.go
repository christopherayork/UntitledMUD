package definitions

import "fmt"

type Tangible struct {
	name string
	description string
	contents []*Individual
	loc *Gridded
	x int
	y int
}

type Gridded interface {
	Enter(target Display, x, y int) bool
	Entered(target Display)
	Exit(target Display, x, y int) bool
	Exited(target Display)
}
type Mapped interface {
	GetLocs() [][]int
	SetCoords(x, y int)
}

type Display interface {
	String() string
}

func (t Tangible) Apply(target Display) {
	fmt.Println(t.String())
	// probably want to format this better later and make it more immersive
}

func (t Tangible) String() string {
	return fmt.Sprintf("Type: %T, Name: %v, Description: %v", t, t.name, t.description)
}

func (t Tangible) Enter(target Display, x, y int) bool {
	// define some rules for movement procedures
	// this will check if movement is permitted into this tangible
	return true
}


func (t Tangible) Entered(target Display) {
	// if Enter() is validated as truthy, this will be called to complete the process and add addition room for behavior
	// this will return nothing as of now
}

func (t Tangible) Exit(target Display, x, y int) bool {
	// called to permit exiting
	return false // no allowance for now
}

func (t Tangible) Exited(target Display) {
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

