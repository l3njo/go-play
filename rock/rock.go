package rock

import (
	"fmt"
	"strings"

	"github.com/l3njo/play/misc"
)

// Get returns a random value from {"rock", "paper", "scissors"}
func Get() string {
	return misc.Choices[misc.Int(len(misc.Choices))]
}


// Play returns a string stating what each player chose and a bool for whether the user won.
func Play(userChoice string) (result string) {
	result = ""
	_, valid := misc.ChoiceMap[userChoice]
	if !valid {
		result = "Error! Invalid choice."
		return
	}
	userChoice = strings.ToLower(userChoice)
	computerChoice := misc.Choices[misc.Int(len(misc.Choices))]
	switch userChoice {
		case computerChoice:
			result = fmt.Sprintf("You drew! You both chose %v.", computerChoice)
		case misc.ChoiceMap[computerChoice]:
			result = fmt.Sprintf("You lost! You chose %v and the computer chose %v.", userChoice, computerChoice)
		default:
			result = fmt.Sprintf("You won! You chose %v and the computer chose %v.", userChoice, computerChoice)
	}
	return
}

// func main() {
// 	var choice string
// 	choices := map[int]int{
// 		1: 0,
// 		0: 2,
// 		2: 1,
// 	}

// 	computerMove := misc.Int(3)
// 	humanMove := -1

// 	fmt.Scanf("%s\n", &choice)
// 	choice = strings.ToLower(choice)

// 	if choice == "rock" {
// 		humanMove = 0
// 	} else if choice == "paper" {
// 		humanMove = 1
// 	} else if choice == "scissors" {
// 		humanMove = 2
// 	}

// 	if humanMove != -1 {
// 		switch humanMove {
// 		case computerMove:
// 			fmt.Println("Tie, Replay!")
// 		case choices[computerMove]:
// 			fmt.Println("Computer Wins!")
// 			break loop
// 		default:
// 			fmt.Println("You Win!")
// 			break loop
// 		}
// 	}
// }
