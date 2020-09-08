package definitions


// Generates a Map{} with designated params
// Currently, it takes no seed, since our target map is fixed
type MapGenerator struct {
	// seed string
	width int
	height int

}

func (m MapGenerator) Generate() (*Map, bool) {
	Map, _ := NewMap()
	Region, err := NewRegion(Map, 1, 1)
	if _, ok := err.(error); ok { return nil, false }
	Zone, _ := NewZone(Region, 1, 1)
	areas := make([]*Area, 2)
	Area1, _ := NewArea(Zone, 1, 1)
	areas[0] = Area1
	Area2, _ := NewArea(Zone, 1, 2)
	areas[1] = Area2
	plots := make(map[string]map[string]*Plot)
	for _, v := range areas {
		for x := 1; x < 3; x++ {
			for y := 1; y < 3; y++ {
				Plot, _ := NewPlot(v, x, y)
				plots[string(x)][string(y)] = Plot
				if x == 2 && y == 2 { // upper-right corner is going to be a special plot
					Plot.description = "A dark and putrid hole in the cave wall leads to this menacing room, with the light sounds of drumming and low clamoring of fiends within playing outwards."
				}
			}
		}
	} // plots created, now lets make tiles
	// very simple square algorithm
	for _, v := range plots {
		for _, v2 := range v {
			for x := 1; x < 3; x++ {
				for y := 1; y < 3; y++ {
					Tile, _ := NewTile(v2, x, y)
					Tile.description = "Boop"
				}
			}
		}
	}
	return Map, true
}

func NewMapGen(x, y int) (*Map, bool) {
	generator := MapGenerator{width: x, height: y}
	return generator.Generate()
}

var mockup = map[string]map[string]map[string]map[string]map[string]string{
	"1,1": map[string]map[string]map[string]map[string]string{
		"1,1": map[string]map[string]map[string]string{
			"1,1": map[string]map[string]string{
				"1,1": map[string]string{
					"1,1": "Beep1",
					"1,2": "Beep2",
					"2,1": "Beep3",
					"2,2": "Beep4",
				},
				"1,2": map[string]string{
					"1,1": "Beep1",
					"1,2": "Beep2",
					"2,1": "Beep3",
					"2,2": "Beep4",
				},
				"2,1": map[string]string{
					"1,1": "Beep1",
					"1,2": "Beep2",
					"2,1": "Beep3",
					"2,2": "Beep4",
				},
				"2,2": map[string]string{
					"1,1": "Beep1",
					"1,2": "Beep2",
					"2,1": "Beep3",
					"2,2": "Beep4",
				},
			},
			"1,2": map[string]map[string]string{
				"1,1": map[string]string{
					"1,1": "Beep1",
					"1,2": "Beep2",
					"2,1": "Beep3",
					"2,2": "Beep4",
				},
				"1,2": map[string]string{
					"1,1": "Beep1",
					"1,2": "Beep2",
					"2,1": "Beep3",
					"2,2": "Beep4",
				},
				"2,1": map[string]string{
					"1,1": "Beep1",
					"1,2": "Beep2",
					"2,1": "Beep3",
					"2,2": "Beep4",
				},
				"2,2": map[string]string{
					"1,1": "Beep1",
					"1,2": "Beep2",
					"2,1": "Beep3",
					"2,2": "Beep4",
				},
			},
		},
	},
}
