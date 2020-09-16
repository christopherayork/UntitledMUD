package definitions

/*
String this together so that when an Individual enters a Tile, if the Tile isn't within the same Plot that the Tile it left was in
Plot.Enter() gets run on the Individual.
This goes all the way up the chain so that If the player Exits a plot and Enters a new one with a different Area, Area.Enter()
Is called on the Individual.
If they leave an Area and enter a new one that isn't of the same Zone, Zone.Enter() gets called on the individual
Leaving a Zone and entering a new one in a different Region calls Region.Enter() on the individual as well.
*/

type Tile struct {
	Tangible
	grid *Grid
}

// Creates a new Tile{} and returns it if it can successfully Enter() a parent Plot{}.
// The dirs parameter is optional and takes a map expected to hold *Tile pointers with the options of n, s, e, w for keys.
func NewTile(grid Grid, x, y int) (*Tile, error) {
	tile := Tile{}
	tile.x = x
	tile.y = y
	tile.grid = &grid
	return &tile, nil
}

func containsInd(inds []*Individual, ind *Individual) bool {
	for _, v := range inds {
		if v == ind { return true }
	}
	return false // we'll add this in later!
}

func removeInd(inds []*Individual, ind *Individual) (bool, []*Individual) {
	for i, v := range inds {
		if v == ind {
			if i == 0 && len(inds) > 1 { // ind is the first element and the slice has more after
				return true, inds[1:len(inds)]
			} else if i == 0 { // ind is the first and only element
				return true, make([]*Individual, 0, 1)
			} else if i == len(inds) { // ind is at the end of the slice with other elements before it
				return true, inds[0:len(inds)-1]
			} else { // ind is somewhere in the middle of the slice
				return true, append(inds[0:i], inds[i+1:len(inds)]...)
			}
		}
	}
	return false, inds
}

func (t Tile) GetDir(dir string) *Tangible {
	switch dir {
		case "north":
			return t.grid.GetValue("tiles", t.x, t.y + 1)
		case "south":
			return t.grid.GetValue("tiles", t.x, t.y - 1)
		case "east":
			return t.grid.GetValue("tiles", t.x + 1, t.y)
		case "west":
			return t.grid.GetValue("tiles", t.x - 1, t.y)
	}
	return nil
}

// tiles don't need x or y, they are single points on the map
// they also will not register as part of the Gridded interface, which is good because they are not Grid holders
func (t Tile) Enter(target Individual) bool {
	exists := containsInd(t.contents, &target)
	if exists { return true
	} else {
		tloc := *t.loc
		// currently, a tile must be in a location for it to be valid for entering
		if _, ok := tloc.(Gridded); ok {
			parentPerms := tloc.Enter(target, t.x, t.y)
			if !parentPerms { return false }
			t.contents = append(t.contents, &target)
			defer t.Entered(target)
			return true
		}
	}
	return false
	// this is not finished, we're going to use containsInd() to see if the item is in the slice
	// if it is, we return true (though it would be a bug to try to enter a tile you're already in)
	// if they aren't, we add them in and return true on success
}

func (t Tile) Entered(target Individual) {
	// we currently have no overloading to add to the process with this
	// it's merely a placeholder for additional effects in future release canvases
	if d, ok := target.(Display); ok {
		t.Apply(d)
	}
}

func (t Tile) Exit(target Individual) bool {
	removed, newContents := removeInd(t.contents, &target)
	if removed {
		defer t.Exited(target)
		t.contents = newContents
		return true
	} else {
		return true
	}
	// currently there are no mechanics to bind someone to a location
	// perhaps I should tie the Exit call to whether or not the next tile allows the Enter
	// or perhaps Exit is only called after the next tile allows entry
	// decide more on that later
}

func (t Tile) Exited(target Individual) {
	// placeholder for future release canvases
}
