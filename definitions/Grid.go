package definitions
import "strconv"

type Grid struct {
	grid map[string]map[string]*Tangible
	parent *Gridded
}

/*
	Grid rework!
	Currently, each individual grid has to be pieced together when generating a more cohesive map
	also, when auto-connecting tiles to each other, (and every other non movable tangible), it will
	simplify the process by placing everything in it's proper position with no duplicates possible
	as is, there can be duplicate tangibles within a grid since they are separate
	the following change will appear as

	grid: {
		maps: {
			z: map {...},
			z: map {...}
		},
		regions: {
			x,y: region {...},
			x,y: region {...}
		},
		zones: {
			x,y: zone {...},
			x,y: zone {...}
		},
		areas: {
			x,y: area {...},
			x,y: area {...}
		},
		plots: {
			x,y: plot {...},
			x,y: plot {...}
		},
		tiles: {
			x,y: tile {...},
			x,y: tile {...}
		}
	}

	other things to consider are, with this change, how will each parent grouping spread over the map
	will they be forced into rectangular formations?
	will they take up many keys (x,y pairs) within their respective grid?
	idea:
	group: {
		x,y: &group1,
		x,y: &group1,
		x,y: &group2,
		x,y: &group3
	}
	the idea is the same as relational databases; but in this case the prospect of persistence is ignored
	since we're using pointers, they will be invalid to use in reloading a map, additional measures will be required if the system ever moves towards it

	one other thing to finalize is the idea of the z axis, but for the sake of mvp, we can leave that out until future updates
 */

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
