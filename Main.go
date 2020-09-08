package UntitledMUD

import def "./definitions"

//var MAP_GEN_SEED = "boopyboop"
var MAP_GEN_WIDTH = 50
var MAP_GEN_HEIGHT = 50

func main() {
	Map, ok := def.NewMapGen(MAP_GEN_WIDTH, MAP_GEN_HEIGHT)
	// do any other things like world loading
	// as the scope expands, player client and server would be seperated into different applications
	if !ok { return } // we can't play without a map
	world = NewWorld(Map)
	// when we use the MapGenerator, if we ever implement saving we will need much additional functionality to detect if a map has been generated
	// we need to create a world that holds our map, and is globally accessible to the gameloop and other game methods
	game := true
	for game { game = GameLoop() }
	// we don't want to automatically exit the program just because the game loop ended
	// we might want to auto-reconnect on connection drop
	// or we might want to start a new game by returning true on the loop
	//fmt.Scanln()
	//fmt.Println("Done")
}

type World struct {
	space *def.Map
	// we could fill in a lot of extra useful utility things for the World struct like time, etc
}

func NewWorld(m *def.Map) World {
	w := World{}
	w.space = m
	return w
}

var world World



