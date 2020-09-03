package definitions


import "errors"


// smallest grid available
// add extra handlers to support adding tiles to it's grid
type Plot struct {
	Tangible
	grid *Grid
}

func NewPlot(g Gridded) (*Plot, error) {
	if _, ok := g.(Area); !ok {
		return &Plot{}, errors.New("error: NewPlot(), argument for parameter g must be of type Area")
	}
	plot := Plot{}
	plot.grid = NewGrid(plot)
	plot.loc = g
	return &plot, nil
}

func (p Plot) Enter(target interface{}, x, y int) bool {
	if val, ok := target.(Tile); ok {
		gridSuccess := p.grid.Enter(&val, x, y)
		return gridSuccess
	} else { return false }
}

func (p Plot) Entered(target *Tangible) {

}

func (p Plot) Exit(target *Tangible, x, y int) bool {
	return true
}

func (p Plot) Exited(target *Tangible) {

}