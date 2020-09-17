package definitions

import (
	"fmt"
	"strconv"
)

type Grid struct {
	grid map[string]map[string]map[string]*Mapped
	maxx int
	maxy int
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
			x: { y: map {...} },
			x: { y: map {...} }
		},
		regions: {
			x: { y: region {...} },
			x: { y: region {...} }
		},
		zones: {
			x: { y: zone {...} },
			x: { y: zone {...} }
		},
		areas: {
			x: { y: area {...} },
			x: { y: area {...} }
		},
		plots: {
			x: { y: plot {...} },
			x: { y: plot {...} }
		},
		tiles: {
			x: { y: tile {...} },
			x: { y: tile {...} }
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
	//return fmt.Sprintf("Grid: %v", g.grid)
	str := ""
	str += g.Grid2String(g.grid["region"], "Region")
	str += g.Grid2String(g.grid["zone"], "Zone")
	str += g.Grid2String(g.grid["area"], "Area")
	str += g.Grid2String(g.grid["plot"], "Plot")
	str += g.Grid2String(g.grid["tile"], "Tile")

	fmt.Println(g.grid["region"], g.grid["zone"], g.grid["area"], g.grid["plot"], g.grid["tile"])
	return str
}

func (g Grid) Grid2String(sect map[string]map[string]*Mapped, mtype string) string {
	str := fmt.Sprintf("%v", mtype)
	str += "\n\n"
	for y := g.maxy; y >= 1; y-- { // top to bottom
		for x := 1; x < g.maxx; x++ { // left to right
			fmt.Print("running ")
			if xmap, okxm := sect[strconv.Itoa(x)]; okxm {
				if _, okyv := xmap[strconv.Itoa(y)]; okyv {
					str += "[*]"
				} else {
					str += "[ ]"
				}
			}
		}
		str += "\n"
	}
	str += "\n"
	return str
}


// make a function to link all tangibles

func NewGrid() *Grid {
	return &Grid{grid: make(map[string]map[string]map[string]*Mapped), maxx: 0, maxy: 0}
}

/* Gets a value out of the Grid's map
Takes in values such that g.grid["region"]["1"]["1"] returns a *Mapped.
Valid call would be g.GetValue("zone", "3", "4").
If a *Tangible exists within that location, it will be returned
*/
func (g Grid) GetValue(sect string, x, y int) *Mapped {
	if section, ok := g.grid[sect]; ok {
		if v, ok := section[strconv.Itoa(x)]; ok {
			if v2, ok2 := v[strconv.Itoa(y)]; ok2 {
				return v2
			}
		}
	}
	return nil
}

// Generates a key out of a target type for the grid categories.
// If it returns an empty string, the target does not map to any valid categories in the grid
func GetGridCat(target interface{}) string {
	key := ""
	switch v := target.(type) {
		case Map: key = "map"
		case Region: key = "region"
		case Zone: key = "zone"
		case Area: key = "area"
		case Plot: key = "plot"
		case Tile: key = "tile"
		default: fmt.Println(fmt.Sprintf("error: GetGridCat(), cannot operate on type %v", v))
	}
	return key
}

// To pass a single item into the grid, supply an x and y.
// To pass multiple items into the grid, supply a map[string]map[string]string as an arg at the end
// Takes the format {x: {y: true}}
func (g Grid) Enter(target interface{}, x int, y int, coords ...map[string]map[string]bool) bool {
	success := false
	//return true
	key := GetGridCat(target)
	if len(key) == 0 { return false } // we can't enter a non valid type!
	if x > 0 && y > 0 {
		success = g.Add(target, strconv.Itoa(x), strconv.Itoa(y), key)
	} else if len(coords) > 0 {
		mappings := coords[0]
		for xc, v := range mappings {
			for yc, v2 := range v {
				if v2 {
					success = g.Add(target, xc, yc, key)
					fmt.Println("Grid.Enter() coords success")
				}
			}
		}
	}
	return success
}

func (g Grid) Add(target interface{}, x, y string, key string) bool {
	// return true
	if _, ok := g.grid[key]; !ok { // make sure the section is added if it doesn't already exist
		g.grid[key] = make(map[string]map[string]*Mapped)
	}
	tmap := g.grid[key]
	if _, ok := tmap[x]; !ok {
		tmap[x] = make(map[string]*Mapped)
	}
	if grd, ok2 := target.(Mapped); ok2 {
		tmap[x][y] = &grd
		xint, xerr := strconv.Atoi(x)
		yint, yerr := strconv.Atoi(y)
		if g.maxx < xint { g.maxx = xint } // make sure these are getting set
		if g.maxy < yint { g.maxy = yint }
		// now that grids are centralized, it needs manually synced on updates
		// i need to figure out an easy way for tangibles to be containerized into their parents
		//tan.loc = *g.parent
		if xerr == nil && yerr == nil { grd.SetCoords(xint, yint) }
		defer g.Entered(&grd)
		return true
	}
	return false
}

func (g Grid) Entered(target interface{}) {

}

func (g Grid) Exit(target interface{}) bool {
	// consider deleting the keys when they're empty, instead of leaving nil entries
	success := false
	key := GetGridCat(target)
	tmap := g.grid[key]
	if tan, oktan := target.(Tangible); oktan {
		if _, ok := tmap[strconv.Itoa(tan.x)]; ok {
			if _, ok2 := tmap[strconv.Itoa(tan.x)][strconv.Itoa(tan.y)]; ok2 {
				tmap[strconv.Itoa(tan.x)][strconv.Itoa(tan.y)] = nil
			}
		}
		success = true
		tan.loc = nil
		tan.x = 0
		tan.y = 0
		defer g.Exited(&tan)
	}
	return success
}

func (g Grid) Exited(target interface{}) {

}
