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

func (m MapGenerator) Generate() (*Map, bool) {
	Map, _ := NewMap()
	var mapData map[string]interface{}
	ReadJSON("map1.json", &mapData)
	var legend map[string]map[string]string
	if lgn, okl := mapData["legend"].(map[string]map[string]string); okl {
		legend = lgn
	}
	if maps, okm := mapData["map"].(map[string]interface{}); okm {
		for kregion, vregion := range maps {
			// this is the keys for regions
			// split key into it's coordinates
			if zones, okr := vregion.(map[string]interface{}); okr {
				xr, yr := GetCoords(kregion)
				Region, _ := NewRegion(Map, xr, yr)
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
												tilevals := legend[tile]
												Tile.name = tilevals["Name"]
												Tile.description = tilevals["Description"]
											} // should we fail the whole thing on individual tile fail?
										}
									} else { return nil, false }
								}
							} else { return nil, false }
						}
					} else { return nil, false }
				}
			} else { return nil, false }
		}
	}
	return Map, true
}

// unloads Unmarshal'd JSON into the result pointer (map of interfaces).
// no return here, just hang onto your result pointer
func ReadJSON(fileName string, result *map[string]interface{}) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error: ReadJSON(), cannot open that file")
		return
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), result)
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
	generator := MapGenerator{width: x, height: y}
	return generator.Generate()
}
