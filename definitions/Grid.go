package definitions

type Grid struct {
	grid map[string]map[string]*Tangible
	parent *Gridded
}

func NewGrid(p interface{}) *Grid {
	if val, ok := p.(Gridded); ok {
		return &Grid{grid: make(map[string]map[string]*Tangible), parent: &val}
	} else { return &Grid{} }
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
