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


func (m MapGenerator) Generate() (*Map, bool) {
	Grid := NewGrid()
	Map, _ := NewMap()
	Grid.Enter(Map, 1,1) // we'll only have a single map for now, but later abstractions will call for reading the map coords out of the map file
	var mapData map[string]interface{}
	errjs := ReadJSON("map1.json", &mapData)
	if errjs != nil { fmt.Println(errjs) }
	var legend map[string]interface{}
	//fmt.Println(mapData)
	if lgn, okl := mapData["legend"].(map[string]interface{}); okl {
		legend = lgn
	} else {
		fmt.Println(fmt.Sprintf("%v, %v", lgn, okl))
		fmt.Println("error: MapGenerator.Generate(), legend failed to load from mapdata")
	}
	//return Map, true
	if maps, okm := mapData["map"].(map[string]interface{}); okm {
		for kregion, vregion := range maps {
			// these are the keys for regions
			// split key into it's coordinates
			if zones, okr := vregion.(map[string]interface{}); okr {
				xr, yr := GetCoords(kregion)
				Region, _ := NewRegion(Map, *Grid, xr, yr)
				for kzone, vzone := range zones {
					if areas, okz := vzone.(map[string]interface{}); okz {
						xz, yz := GetCoords(kzone)
						Zone, _ := NewZone(Region, xz, yz)
						for karea, varea := range areas {
							if plots, oka := varea.(map[string]interface{}); oka {
								xa, ya := GetCoords(karea)
								Area, _ := NewArea(Zone, xa, ya)
								for kplot, vplot := range plots {
									if tiles, okp := vplot.(map[string]interface{}); okp {
										xp, yp := GetCoords(kplot)
										Plot, _ := NewPlot(Area, xp, yp)
										for ktile, vtile := range tiles {
											if tile, okt := vtile.(string); okt {
												xt, yt := GetCoords(ktile)
												Tile, _ := NewTile(Plot, xt, yt)
												if tilevals, tvok := legend[tile].(map[string]string); tvok {
													// tile holds a string with a reference id for legend
													Tile.name = tilevals["Name"]
													Tile.description = tilevals["Description"]
												}
											} // should we fail the whole thing on individual tile fail?
										}
									} else {
										fmt.Println("error: MapGenerator.Generate(), Plot failed to be loaded from mapdata")
										continue
									}
								}
							} else {
								fmt.Println("error: MapGenerator.Generate(), Area failed to be loaded from mapdata")
								continue

							}
						}
					} else {
						fmt.Println("error: MapGenerator.Generate(), Zone failed to be loaded from mapdata")
						continue
					}
				}
			} else {
				fmt.Println("error: MapGenerator.Generate(), Region failed to be loaded from mapdata")
				continue
			}
		}
	} else {
		fmt.Println("error: MapGenerator.Generate(), Map failed to be loaded from mapdata")
		return nil, false
	}
	fmt.Println("Map loaded successfully!")
	return Map, true
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

func NewMapGen(x, y int) (*Map, bool) {
	defer panicRecover()
	generator := MapGenerator{width: x, height: y}
	return generator.Generate()
}

func panicRecover() {
	recover()
}
