package definitions

import "errors"



type Tile struct {
	Tangible
}

func NewTile(g Gridded) (*Tile, error) {
	if tan, ok := g.(Plot); ok {
		tile := Tile{}
		// call g.Enter(tile and get permission back before setting tile.loc)
		tile.loc = tan
		return &tile, nil
	} else {
		return &Tile{}, errors.New("error: NewTile(), argument for parameter g required to be type Plot")
	}
}

func containsInd([]Individual, *Individual) bool {

	return true // we'll add this in later!
}

// tiles don't need x or y, they are single points on the map
// they also will not register as part of the Gridded interface, which is good because they are not Grid holders
func (t Tile) Enter(target Individual) bool {
	isInd := false
	switch ind := target.(type) {
		case Mob: isInd = true
		case Obj: isInd = true
		default: isInd = false
	}
	if !isInd { return false }
	return ind // this is not finished, we're going to use containsInd() to see if the item is in the slice
	// if it is, we return true (though it would be a bug to try to enter a tile you're already in)
	// if they aren't, we add them in and return true on success
}

func (t Tile) Entered(target Display) {

}

func (t Tile) Exit(target Display) bool {
	return true
}

func (t Tile) Exited(target Display) {

}
