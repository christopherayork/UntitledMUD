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

func (g Grid) GetValue(x, y int) *Tangible {
	if v, ok := g.grid[string(x)]; ok {
		if v2, ok2 := v[string(y)]; ok2 {
			return v2
		}
	}
	return nil
}

func (g Grid) Enter(target interface{}, x int, y int) bool {
	success := false
	if _, ok := g.grid[string(x)]; !ok {
		g.grid[string(x)] = make(map[string]*Tangible)
	}
	if tan, ok2 := target.(Tangible); ok2 {
		g.grid[string(x)][string(y)] = &tan
		tan.loc = *g.parent
		tan.x = x
		tan.y = y
		success = true
		defer g.Entered(&tan)
		return success
	} else { return false }
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
