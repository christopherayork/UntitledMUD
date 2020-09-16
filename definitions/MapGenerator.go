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


func PullField(mapData map[string]interface{}, sect string) []map[string]interface{} {
	var result []map[string]interface{}
	if sec, ok := mapData[sect].([]map[string]interface{}); ok { result = sec } else {
		fmt.Println(fmt.Sprintf("%v, %v", sec, ok))
		fmt.Println(fmt.Sprintf("error: MapGenerator.Generate(), %v failed to load from mapdata", sect))
	}
	return result
}

func CreateTypes(targets []map[string]interface{}, g *Grid, callback func(Grid, int, int, [][]int)) {
	// Currently reads a block formation from the resulting interface and creates it's instances from that reading
	// Perhaps a future update could make use of the pairs data, once it's filled out

	for _, v := range targets {
		if val, ok := v["block"]; ok {
			if valmap, okvm := val.(map[string][]int); okvm {
				if xslice, okx := valmap["x"]; okx {
					if yslice, oky := valmap["y"]; oky {
						pvals := make([][]int, 1)
						for x := xslice[0]; x <= xslice[1]; x++ {
							for y := yslice[0]; y <= yslice[1]; y++ {
								set := []int {x, y}
								pvals = append(pvals, set)
							}
						}
						callback(*g, 0, 0, pvals)
						// pass in the full array of coordinate sets that it covers on the grid, so the constructor can use them
					}
				}
			}
		}
	}
}
func CreateTypesTile(targets []map[string]interface{}, g *Grid, callback func(Grid, int, int, string, string), legend map[string]map[string]string) {
	for _, v := range targets {
		if val, ok := v["block"]; ok {
			if valmap, okvm := val.(map[string]int); okvm {
				if d, okd := v["desc"]; okd {
					if key, okkey := d.(string); okkey {
						if proto, okproto := legend[key]; okproto {
							if name, okname := proto["Name"]; okname {
								if desc, okdesc := proto["Description"]; okdesc {
									if x, okx := valmap["x"]; okx {
										if y, oky := valmap["y"]; oky {
											callback(*g, x, y, name, desc)
											// after verifying all this stuff, we can finally use it to create our tile!
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func (m MapGenerator) Generate() (*Grid, bool) {
	NGrid := NewGrid()
	NMap, _ := NewMap()
	NGrid.Enter(NMap, 1,1) // we'll only have a single map for now, but later abstractions will call for reading the map coords out of the map file
	var mapData map[string]interface{}
	errjs := ReadJSON("map1.json", &mapData)
	if errjs != nil { fmt.Println(errjs) }
	var legend map[string]map[string]string
	if lgn, okl := mapData["legend"].(map[string]map[string]string); okl { legend = lgn } else {
		fmt.Println(fmt.Sprintf("%v, %v", lgn, okl))
		fmt.Println("error: MapGenerator.Generate(), legend failed to load from mapdata")
	}
	regions := PullField(mapData, "regions")
	zones := PullField(mapData, "zones")
	areas := PullField(mapData, "areas")
	plots := PullField(mapData, "plots")
	tiles := PullField(mapData, "tiles")
	// NewRegion(Map, Grid, 0, 0, regions)
	// will need to change NewRegion() to support the new map format
	// we need a callback for each, since they return a different type in the same format
	// we can't tell our CreateTypes() function to accept all of these without changing the return types for each constructor, which we don't want
	rcb := func(g Grid, x, y int, coords [][]int) { _, _ = NewRegion(g, x, y, coords) }
	zcb := func(g Grid, x, y int, coords [][]int) { _, _ = NewZone(g, x, y, coords) }
	acb := func(g Grid, x, y int, coords [][]int) { _, _ = NewArea(g, x, y, coords) }
	pcb := func(g Grid, x, y int, coords [][]int) { _, _ = NewPlot(g, x, y, coords) }
	tcb := func(g Grid, x, y int, name, desc string) { _, _ = NewTile(g, x, y, name, desc) }
	CreateTypes(regions, NGrid, rcb)
	CreateTypes(zones, NGrid, zcb)
	CreateTypes(areas, NGrid, acb)
	CreateTypes(plots, NGrid, pcb)
	CreateTypesTile(tiles, NGrid, tcb, legend)
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
