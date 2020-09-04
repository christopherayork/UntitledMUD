package definitions


import "errors"


// smallest grid available
// add extra handlers to support adding tiles to it's grid
type Plot struct {
	Tangible
	grid *Grid
}

// func (t Tangible) Apply(target Display) {}

func NewPlot(g Gridded) (*Plot, error) {
	if _, ok := g.(Area); !ok {
		return &Plot{}, errors.New("error: NewPlot(), argument for parameter g must be of type Area")
	}
	plot := Plot{}
	plot.grid = NewGrid(plot)
	plot.loc = g
	return &plot, nil
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