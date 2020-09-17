package definitions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Generates a Map{} with designated params
// Currently, it takes no seed, since our target map is fixed
type MapGenerator struct {
	// seed string
	width int
	height int

}

// change Generate to return a Grid instead of a Map

func (m MapGenerator) Generate() (*Grid, bool) {
	NGrid := NewGrid()
	NMap, _ := NewMap()
	NGrid.Enter(NMap, 1,1) // we'll only have a single map for now, but later abstractions will call for reading the map coords out of the map file
	mapData := Mapfile{}
	errjs := ReadMapJSON("map1.json", &mapData)
	if errjs != nil { fmt.Println(errjs) }
	// we need a callback for each, since they return a different type in the same format
	rcb := func(g Grid, x, y int, coords [][]int) { _, _ = NewRegion(g, x, y, coords) }
	zcb := func(g Grid, x, y int, coords [][]int) { _, _ = NewZone(g, x, y, coords) }
	acb := func(g Grid, x, y int, coords [][]int) { _, _ = NewArea(g, x, y, coords) }
	pcb := func(g Grid, x, y int, coords [][]int) { _, _ = NewPlot(g, x, y, coords) }
	tcb := func(g Grid, x, y int, name, desc string) { _, _ = NewTile(g, x, y, name, desc) }
	mapData.Regions.MakeGroup(NGrid, rcb)
	mapData.Zones.MakeGroup(NGrid, zcb)
	mapData.Areas.MakeGroup(NGrid, acb)
	mapData.Plots.MakeGroup(NGrid, pcb)
	mapData.Tiles.MakeGroup(NGrid, tcb, mapData.Legend)
	fmt.Println("Map loaded successfully!")
	return NGrid, true
}

// unloads Unmarshal'd JSON into the result pointer (map of interfaces).
// no return here, just hang onto your result pointer
func ReadJSON(fileName string, result *map[string]interface{}) error {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error: ReadJSON(), cannot open that file")
	}
	defer func() {
		err4 := jsonFile.Close()
		if err4 != nil { fmt.Println(fmt.Errorf("error: ReadJSON(), couldn't close that file")) }
	}()
	byteValue, err2 := ioutil.ReadAll(jsonFile)
	if err2 != nil {
		return fmt.Errorf("error: ReadJSON(), cannot read jsonFile into byteValue array")
	}
	err3 := json.Unmarshal([]byte(byteValue), result)
	if err3 != nil {
		return fmt.Errorf("error: ReadJSON(), json.Unmarshal() failed to map binary to result map")
	}
	return nil
}

func ReadMapJSON(fileName string, result *Mapfile) error {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error: ReadMapJSON(), cannot open that file")
	}
	defer func() {
		err4 := jsonFile.Close()
		if err4 != nil { fmt.Println(fmt.Errorf("error: ReadMapJSON(), couldn't close that file")) }
	}()
	byteValue, err2 := ioutil.ReadAll(jsonFile)
	if err2 != nil {
		return fmt.Errorf("error: ReadMapJSON(), cannot read jsonFile into byteValue array")
	}
	err3 := json.Unmarshal([]byte(byteValue), result)
	if err3 != nil {
		return fmt.Errorf("error: ReadMapJSON(), json.Unmarshal() failed to map binary to result map")
	}
	return nil
}

func GetCoords(merged string) (int, int) {
	arr := strings.Split(merged, ",")
	if len(arr) == 1 {
		firstVal, _ := strconv.Atoi(arr[0])
		return firstVal, 0
	} else if len(arr) < 1 {
		return 0, 0
	} else {
		firstVal, _ := strconv.Atoi(arr[0])
		secondVal, _ := strconv.Atoi(arr[1])
		return firstVal, secondVal
	}
}

func NewMapGen(x, y int) (*Grid, bool) {
	defer panicRecover()
	generator := MapGenerator{width: x, height: y}
	return generator.Generate()
}

func panicRecover() {
	recover()
}

type MapBlock struct {
	X []int `json:"x"`
	Y []int `json:"y"`
}

func (block MapBlock) SplitX() (int, int) {
	if len(block.X) >= 2 { return block.X[0], block.X[1] }
	return 0, 0
}

func (block MapBlock) SplitY() (int, int) {
	if len(block.Y) >= 2 { return block.Y[0], block.Y[1] }
	return 0, 0
}

type MapBlockTile struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type MapSectionGridGroup []MapSectionGrid

// Takes in a *Grid and callback which calls a type constructor.
// This supplies all the given data for a multi-coordinate Gridded type to be made and expanded over all locations
func (group MapSectionGridGroup) MakeGroup(g *Grid, callback func(Grid, int, int, [][]int)) bool {
	for _, grid := range group {
		if len(grid.Block.X) >= 2 && len(grid.Block.Y) >= 2 {
			pvals := grid.GetPairs()
			callback(*g, 0, 0, pvals)
			return true
		}
	}
	return false
}

type MapSectionGrid struct {
	Pairs [][]int `json:"pairs"`
	Block MapBlock `json:"block"`
}

func (grid MapSectionGrid) GetPairs() [][]int {
	xlower, xupper := grid.Block.SplitX()
	ylower, yupper := grid.Block.SplitY()
	pvals := make([][]int, 0)
	for x := xlower; x <= xupper; x++ {
		for y := ylower; y <= yupper; y++ {
			set := []int {x, y}
			pvals = append(pvals, set)
		}
	}
	return pvals
}

type MapSectionTileGroup []MapSectionTile

// Takes in a *Grid and a callback which holds a constructor for a type.
// This pulls all the necessary data for a Tile constructor and passes it through
func (group MapSectionTileGroup) MakeGroup(g *Grid, callback func(Grid, int, int, string, string), legend MapLegend) bool {
	for _, tile := range group {
		descobj := legend.Get(tile.Desc)
		name, desc := descobj.Name, descobj.Description
		callback(*g, tile.Block.X, tile.Block.Y, name, desc)
		return true
	}
	return false
}

type MapSectionTile struct {
	Block MapBlockTile `json:"block"`
	Desc string `json:"desc"`
}

type MapDescriber struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type MapLegend map[string]MapDescriber

// Get a MapDescriber by id
// Returns an empty MapDescriber instance if none are found
func (legend MapLegend) Get(id string) MapDescriber {
	if val, ok := legend[id]; ok { return val }
	return MapDescriber{}
}

type Mapfile struct {
	Legend MapLegend `json:"legend"`
	Regions MapSectionGridGroup `json:"regions"`
	Zones MapSectionGridGroup `json:"zones"`
	Areas MapSectionGridGroup `json:"areas"`
	Plots MapSectionGridGroup `json:"plots"`
	Tiles MapSectionTileGroup `json:"tiles"`
}
