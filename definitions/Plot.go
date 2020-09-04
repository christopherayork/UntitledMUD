package definitions


import "errors"


// smallest grid available
// add extra handlers to support adding tiles to it's grid
type Plot struct {
	Tangible
	grid *Grid
	north *Plot
	south *Plot
	east *Plot
	west *Plot
}

// func (t Tangible) Apply(target Display) {}

func NewPlot(g Gridded, dirs ...map[string]*Plot) (*Plot, error) {
	if _, ok := g.(Area); !ok {
		return &Plot{}, errors.New("error: NewPlot(), argument for parameter g must be of type Area")
	}
	plot := Plot{}
	plot.grid = NewGrid(plot)
	plot.loc = g
	for i := range dirs {
		if i > 0 { break }
		if v, ok := dirs[i]["n"]; ok { plot.north = v }
		if v, ok := dirs[i]["s"]; ok { plot.south = v }
		if v, ok := dirs[i]["e"]; ok { plot.east = v }
		if v, ok := dirs[i]["w"]; ok { plot.west = v }
	}
	return &plot, nil
}

func (t Plot) SetDirection(d string, v *Plot) bool {
	switch d {
		case "n":
			t.north = v
			v.south = &t // doubly linked
		case "s":
			t.south = v
			v.north = &t
		case "e":
			t.east = v
			v.west = &t
		case "w":
			t.west = v
			v.east = &t
		default: return false
	}
	return true
}

func (p Plot) Enter(target Display, x, y int) bool {
	if tan, ok := target.(Tile); ok {
		gridSuccess := p.grid.Enter(&tan, x, y)
		if gridSuccess { defer p.Entered(tan) }
		return gridSuccess
	}
	return false
	// add additional clauses to allow for this to be called on Individuals for the purpose of entering rooms
	// and having more centric messages or effects be sent/applied to them
}

func (p Plot) Entered(target Display) {
	if _, ok := target.(Individual); ok{
		p.Apply(target)
	}
}

func (p Plot) Exit(target Display, x, y int) bool {
	return true
}

func (p Plot) Exited(target Display) {

}