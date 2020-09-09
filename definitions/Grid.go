package definitions
import "strconv"

type Grid struct {
	grid map[string]map[string]*Tangible
	parent *Gridded
}

func (g Grid) String() string {
	return "Placeholder -> create visual grid from items in grid.map"
}


// make a function to link all tangibles

func NewGrid(p interface{}) *Grid {
	if val, ok := p.(Gridded); ok {
		return &Grid{grid: make(map[string]map[string]*Tangible), parent: &val}
	} else { return &Grid{} }
}

func (g Grid) GetValue(x, y int) *Tangible {
	if v, ok := g.grid[strconv.Itoa(x)]; ok {
		if v2, ok2 := v[strconv.Itoa(y)]; ok2 {
			return v2
		}
	}
	return nil
}

func (g Grid) Enter(target interface{}, x int, y int) bool {
	success := false
	if _, ok := g.grid[strconv.Itoa(x)]; !ok {
		g.grid[strconv.Itoa(x)] = make(map[string]*Tangible)
	}
	if tan, ok2 := target.(Tangible); ok2 {
		g.grid[strconv.Itoa(x)][strconv.Itoa(y)] = &tan
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
	if _, ok := g.grid[strconv.Itoa(target.x)]; ok {
		if _, ok2 := g.grid[strconv.Itoa(target.x)][strconv.Itoa(target.y)]; ok2 {
			g.grid[strconv.Itoa(target.x)][strconv.Itoa(target.y)] = nil
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
