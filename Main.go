package UntitledMUD



func main() {
	// do any other things like world loading
	// as the scope expands, player client and server would be seperated into different applications
	game := true
	for game { game = GameLoop() }
	// we don't want to automatically exit the program just because the game loop ended
	// we might want to auto-reconnect on connection drop
	// or we might want to start a new game by returning true on the loop
	//fmt.Scanln()
	//fmt.Println("Done")
}



