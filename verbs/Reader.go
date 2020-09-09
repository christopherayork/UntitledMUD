package verbs

import (
	"bufio"
	"os"
)

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
