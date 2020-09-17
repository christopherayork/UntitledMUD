package definitions

import (
	"encoding/json"
	"fmt"
	"os"
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
	rcb := func(x, y int, coords [][]int) { _, _ = NewRegion(*NGrid, x, y, coords) }
	zcb := func(x, y int, coords [][]int) { _, _ = NewZone(*NGrid, x, y, coords) }
	acb := func(x, y int, coords [][]int) { _, _ = NewArea(*NGrid, x, y, coords) }
	pcb := func(x, y int, coords [][]int) { _, _ = NewPlot(*NGrid, x, y, coords) }
	tcb := func(x, y int, name, desc string) { _, _ = NewTile(*NGrid, x, y, name, desc) }
	mapData.Regions.MakeGroup(rcb)
	mapData.Zones.MakeGroup(zcb)
	mapData.Areas.MakeGroup(acb)
	mapData.Plots.MakeGroup(pcb)
	mapData.Tiles.MakeGroup(tcb, mapData.Legend)
	fmt.Println("Map loaded successfully!")
	return NGrid, true
}

func ReadMapJSON(fileName string, result *Mapfile)  (err error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error: ReadMapJSON(), cannot open that file")
	}
	defer func() {
		err4 := jsonFile.Close()
		if err4 != nil { err = fmt.Errorf("error: ReadMapJSON(), couldn't close that file") }
	}()
	err = json.NewDecoder(jsonFile).Decode(result)
	if err != nil {
		return fmt.Errorf("error: ReadMapJSON(), json.Unmarshal() failed to map binary to result map")
	}
	return nil
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
	X [2]int `json:"x"`
	Y [2]int `json:"y"`
}

func (block MapBlock) SplitX() (int, int) {
	return block.X[0], block.X[1]
}

func (block MapBlock) SplitY() (int, int) {
	return block.Y[0], block.Y[1]
}

type MapBlockTile struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type MapSectionGridGroup []MapSectionGrid

type GridConstructor func(int, int, [][]int)

// Takes in a *Grid and callback which calls a type constructor.
// This supplies all the given data for a multi-coordinate Gridded type to be made and expanded over all locations
func (group MapSectionGridGroup) MakeGroup(callback GridConstructor) bool {
	for _, grid := range group {
		pvals := grid.GetPairs()
		callback(0, 0, pvals)
		return true
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
	pvals := make([][]int, 0, (xupper-xlower)*(yupper-ylower)) // save on reallocations
	for x := xlower; x <= xupper; x++ {
		for y := ylower; y <= yupper; y++ {
			set := []int {x, y}
			pvals = append(pvals, set)
		}
	}
	return pvals
}

type MapSectionTileGroup []MapSectionTile

type TileConstructor func(int, int, string, string)

// Takes in a *Grid and a callback which holds a constructor for a type.
// This pulls all the necessary data for a Tile constructor and passes it through
func (group MapSectionTileGroup) MakeGroup(callback TileConstructor, legend MapLegend) bool {
	success := false
	for _, tile := range group {
		descobj := legend.Get(tile.Desc)
		name, desc := descobj.Name, descobj.Description
		callback(tile.Block.X, tile.Block.Y, name, desc)
		success = true
	}
	return success
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
