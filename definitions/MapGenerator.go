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
	NGrid.Enter(*NMap, 1,1) // we'll only have a single map for now, but later abstractions will call for reading the map coords out of the map file
	mapData := Mapfile{}
	errjs := ReadMapJSON("map1.json", &mapData)
	if errjs != nil { fmt.Println(errjs) }
	// we need a callback for each, since they return a different type in the same format
	rcb := func(x, y int, coords [][]int) { _, _ = NewRegion(*NGrid, x, y, coords) }
	zcb := func(x, y int, coords [][]int) { _, _ = NewZone(*NGrid, x, y, coords) }
	acb := func(x, y int, coords [][]int) { _, _ = NewArea(*NGrid, x, y, coords) }
	pcb := func(x, y int, coords [][]int) { _, _ = NewPlot(*NGrid, x, y, coords) }
	tcb := func(x, y int, name, desc string) { _, _ = NewTile(*NGrid, x, y, name, desc) }
	fmt.Println("Making regions")
	mapData.Regions.MakeGroup(rcb)
	fmt.Println("Making zones")
	mapData.Zones.MakeGroup(zcb)
	fmt.Println("Making areas")
	mapData.Areas.MakeGroup(acb)
	fmt.Println("Making plots")
	mapData.Plots.MakeGroup(pcb)
	fmt.Println("Making tiles")
	mapData.Tiles.MakeGroup(tcb, mapData.Legend)
	fmt.Println("Map loaded successfully!")
	fmt.Println(mapData)
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
	fmt.Println(group)
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
	fmt.Println(grid.Block)
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

func (mf Mapfile) String() string {
	output := "Mapfile {\n"
	output += "    Legend {\n"
	for key, val := range mf.Legend {
		output += fmt.Sprintf("        %v: {\n", key)
		output += fmt.Sprintf("            Name: %v\n", val.Name)
		output += fmt.Sprintf("            Description: %v\n", val.Description)
		output += "        }\n"
	}
	output += "    }\n"
	msggStringer := func(msgg MapSectionGridGroup, name string) string {
		oVal := fmt.Sprintf("    %v: [\n", name)
		for i, v := range msgg {
			oVal += fmt.Sprintf("        %v: {\n", i)
			oVal += "            Pairs: [\n"
			oVal += "                " // add the space for the first row
			for pairIndex, pairVal := range v.Pairs {
				oVal += fmt.Sprintf("[ %v, %v ] ", pairVal[0], pairVal[1])
				if pairIndex % 4 == 0 {
					oVal += "\n                "
				}
			}
			oVal += "            \n]\n" // close the pairs bracket
			oVal += "            Block: {\n"
			X1, X2 := v.Block.SplitX()
			Y1, Y2 := v.Block.SplitY()
			oVal += fmt.Sprintf("                X: [ %v, %v ]\n", X1, X2)
			oVal += fmt.Sprintf("                Y: [ %v, %v ]", Y1, Y2)
			oVal += "            }\n" // close the block brace
			oVal += "        }\n" // close the loop opened brace
		}
		oVal += "    ]\n"
		return oVal
	}
	output += msggStringer(mf.Regions, "Regions")
	output += msggStringer(mf.Zones, "Zones")
	output += msggStringer(mf.Zones, "Areas")
	output += msggStringer(mf.Zones, "Plots")
	mstgStringer := func(mstg MapSectionTileGroup) string {
		oVal := "    Tiles: [\n"
		for tIndex, tVal := range mstg {
			oVal += fmt.Sprintf("        %v: {\n", tIndex)
			oVal += fmt.Sprintf("            Block: { X: %v, Y: %v }\n", tVal.Block.X, tVal.Block.Y)
			oVal += fmt.Sprintf("            Desc: %v\n", tVal.Desc)
			oVal += "        }\n"
		}
		oVal += "    ]\n"
		return oVal
	}
	output += mstgStringer(mf.Tiles)
	output += "}" // all done!
	return output
}
