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


func PullField(mapData map[string]interface{}, sect string) []interface{} {
	var result []interface{}
	if section, oksect := mapData[sect]; oksect {
		if sec, ok := section.([]interface{}); ok { result = sec } else {
			fmt.Println(fmt.Sprintf("%v, %v", sec, ok))
			fmt.Println(fmt.Sprintf("error: MapGenerator.Generate(), %v failed to load from mapdata", sect))
		}
	}
	return result
}

func CreateTypes(targets []interface{}, g *Grid, callback func(Grid, int, int, [][]int)) {
	// Currently reads a block formation from the resulting interface and creates it's instances from that reading
	// Perhaps a future update could make use of the pairs data, once it's filled out

	for _, target := range targets {
		if v, okv := target.(map[string]interface{}); okv {
			if val, ok := v["block"]; ok {
				if valmap, okvm := val.(map[string]interface{}); okvm {
					//fmt.Println("Reached line 43")
					if xsliceint, okx := valmap["x"]; okx {
						if ysliceint, oky := valmap["y"]; oky {
							if xslice, okxs := xsliceint.([]interface{}); okxs {
								if yslice, okys := ysliceint.([]interface{}); okys {
									if xlower, xlowok := xslice[0].(float64); xlowok {
										if xupper, xupok := xslice[1].(float64); xupok {
											if ylower, ylowok := yslice[0].(float64); ylowok {
												if yupper, yupok := yslice[1].(float64); yupok {
													//return
													pvals := make([][]int, 0)
													for x := int(xlower); x <= int(xupper); x++ {
														for y := int(ylower); y <= int(yupper); y++ {
															set := []int {int(x), int(y)}
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
						}
					}
				}
			}
		}
	}
}
func CreateTypesTile(targets []interface{}, g *Grid, callback func(Grid, int, int, string, string), legend map[string]interface{}) {
	for _, target := range targets {
		if v, okv := target.(map[string]interface{}); okv {
			if val, ok := v["block"]; ok {
				if valmap, okvm := val.(map[string]interface{}); okvm {
					if d, okd := v["desc"]; okd {
						if key, okkey := d.(string); okkey {
							if expanded, okexpanded := legend[key]; okexpanded {
								if proto, okproto := expanded.(map[string]interface{}); okproto {
									if nameint, okname := proto["Name"]; okname {
										if descint, okdesc := proto["Description"]; okdesc {
											if xinterface, okx := valmap["x"]; okx {
												if x, okx2 := xinterface.(float64); okx2 {
													if yinterface, oky := valmap["y"]; oky {
														if y, oky2 := yinterface.(float64); oky2 {
															if name, okname2 := nameint.(string); okname2 {
																if desc, okdesc2 := descint.(string); okdesc2 {
																	callback(*g, int(x), int(y), name, desc)
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
							} else {
								fmt.Println(fmt.Sprintf("Failed to load legend[key](legend[%v])", key))
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
	//fmt.Println(mapData)
	var legend map[string]interface{}
	if leg, okleg := mapData["legend"]; okleg {
		if lgn, okl := leg.(map[string]interface{}); okl { legend = lgn } else {
			fmt.Println(fmt.Sprintf("%v, %v", lgn, okl))
			fmt.Println("error: MapGenerator.Generate(), legend failed to load from mapdata")
		}
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
	// figure out why the other types arent loading into the grid properly
	// region appears to be loading fine
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
