package rock

import (
	"fmt"
	"strings"

	"github.com/l3njo/play/misc"
)

// Get returns a slice of random values from {"rock", "paper", "scissors"}
func Get(num int) (resultSet []string, err string) {
	resultSet = []string{}
	err = ""
	for len(resultSet) < num{
		resultSet = append(resultSet, misc.Choices[misc.Int(len(misc.Choices))])
	}
	return
}


// Play returns a slice of strings stating what each player chose and whether the user won.
func Play(userChoices ...string) (resultSet []string, err string) {
	resultSet = []string{}
	err = ""
	for i, userChoice := range userChoices{
		_, valid := misc.ChoiceMap[userChoice]
		if !valid {
			err = err + fmt.Sprintf("Error on input %v! Invalid choice.\n", i)
			continue
		}
		userChoice = strings.ToLower(userChoice)
		computerChoice := misc.Choices[misc.Int(len(misc.Choices))]
		switch userChoice {
			case computerChoice:
				resultSet = append(resultSet, fmt.Sprintf("%v: You drew! You both chose %v.\n", i, computerChoice))
			case misc.ChoiceMap[computerChoice]:
				resultSet = append(resultSet, fmt.Sprintf("%v: You lost! You chose %v and the computer chose %v.\n", i, userChoice, computerChoice))
			default:
				resultSet = append(resultSet, fmt.Sprintf("%v: You won! You chose %v and the computer chose %v.\n", i, userChoice, computerChoice))
		}
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
