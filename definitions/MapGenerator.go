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

func (m MapGenerator) Generate() (*Grid, bool) {
	Grid := NewGrid()
	Map, _ := NewMap()
	Grid.Enter(Map, 1,1) // we'll only have a single map for now, but later abstractions will call for reading the map coords out of the map file
	var mapData map[string]interface{}
	errjs := ReadJSON("map1.json", &mapData)
	if errjs != nil { fmt.Println(errjs) }
	var legend map[string]interface{}
	if lgn, okl := mapData["legend"].(map[string]interface{}); okl { legend = lgn } else {
		fmt.Println(fmt.Sprintf("%v, %v", lgn, okl))
		fmt.Println("error: MapGenerator.Generate(), legend failed to load from mapdata")
	}
	regions := PullField(mapData, "regions")
	zones := PullField(mapData, "zones")
	areas := PullField(mapData, "areas")
	plots := PullField(mapData, "plots")
	// NewRegion(Map, Grid, 0, 0, regions)
	// will need to change NewRegion() to support the new map format

	fmt.Println("Map loaded successfully!")
	return Grid, true
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
