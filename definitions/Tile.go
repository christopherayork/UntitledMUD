package definitions

import "errors"


type Tile struct {
	Tangible
}

func (t Tile) String() string {
	return t.description
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

// tiles don't need x or y, they are single points on the map
// they also will not register as part of the Gridded interface, which is good because they are not Grid holders
func (t Tile) Enter(target *Tangible) bool {
	return true
}

func (t Tile) Entered(target *Tangible) {

}

func (t Tile) Exit(target *Tangible) bool {
	return true
}

func (t Tile) Exited(target *Tangible) {

}
