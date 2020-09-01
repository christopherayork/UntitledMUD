package UntitledMUD

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	// do any other things like world loading
	// as the scope expands, player client and server would be seperated into different applications
	game := true
	for game { game := GameLoop() }
	//fmt.Scanln()
	//fmt.Println("Done")
}


func GameLoop() bool {
	playing := true
	scanner := bufio.NewScanner(os.Stdin)
	for playing {
		scanner.Scan()
		input := scanner.Text()
		if len(input) < 1 { continue }

	}
	return false
}
