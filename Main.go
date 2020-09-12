package main

//import def "./definitions"
import (
	def "./definitions"
	"os"
	"bufio"
	"fmt"
	"time"
)

//var MAP_GEN_SEED = "boopyboop"
var MAP_GEN_WIDTH = 50
var MAP_GEN_HEIGHT = 50

func main() {
	fmt.Println("Welcome to UntitledMUD!")
	Grid, ok := def.NewMapGen(MAP_GEN_WIDTH, MAP_GEN_HEIGHT)
	// do any other things like world loading
	// as the scope expands, player client and server would be seperated into different applications
	if !ok {
		fmt.Println("MapGen failed...")
		time.Sleep(1000)
		return
	} // we can't play without a map
	world = NewWorld(Grid)
	// when we use the MapGenerator, if we ever implement saving we will need much additional functionality to detect if a map has been generated
	// we need to create a world that holds our map, and is globally accessible to the gameloop and other game methods
	fmt.Println(Grid)
	fmt.Println("Enter a command!")
	scanner := bufio.NewScanner(os.Stdin)
	input := ""
	game := true
	for game {
		//game = verbs.GameLoop()
		scanner.Scan()
		input = scanner.Text()
		fmt.Println(input) // placeholder
	}
	// we don't want to automatically exit the program just because the game loop ended
	// we might want to auto-reconnect on connection drop
	// or we might want to start a new game by returning true on the loop
	//fmt.Scanln()
	//fmt.Println("Done")
}

type World struct {
	space *def.Grid // if we ever expand into multiple worlds, this will hold a Map{}, and Grid{} will be expanded to hold multiple Worlds
	// we could fill in a lot of extra useful utility things for the World struct like time, etc
}

func NewWorld(m *def.Grid) World {
	w := World{}
	w.space = m
	return w
}

var world World



