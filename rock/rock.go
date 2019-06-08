package rock

import (
	"fmt"
	"strings"

	"github.com/l3njo/play/misc"
)

func main() {
	var choice string
	choices := map[int]int{
		1: 0,
		0: 2,
		2: 1,
	}

	computerMove := misc.Int(3)
	humanMove := -1

	fmt.Scanf("%s\n", &choice)
	choice = strings.ToLower(choice)

	if choice == "rock" {
		humanMove = 0
	} else if choice == "paper" {
		humanMove = 1
	} else if choice == "scissors" {
		humanMove = 2
	}

	if humanMove != -1 {
		switch humanMove {
		case computerMove:
			fmt.Println("Tie, Replay!")
		case choices[computerMove]:
			fmt.Println("Computer Wins!")
			break loop
		default:
			fmt.Println("You Win!")
			break loop
		}
	}
}
