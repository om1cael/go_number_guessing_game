package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var difficultyLevels map[string]int = map[string]int{
	"Easy":   10,
	"Medium": 5,
	"Hard":   3,
}

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")
	fmt.Println("")

	for {
		playGame()

		playAgain := "y"
		fmt.Print("Do you want to play again? [Y/n]: ")
		fmt.Scanln(&playAgain)

		if strings.ToLower(playAgain) != "y" {
			break
		}
	}
}

func playGame() {
	chances, err := getDifficultyLevelInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	var guess int
	currentChances := chances
	attempts := 1
	randomNumber := rand.Intn(101)

	start := time.Now()

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		attempts++

		if guess < randomNumber {
			fmt.Printf("Incorrect! The number is greater than %v\n", guess)
		} else if guess > randomNumber {
			fmt.Printf("Incorrect! The number is less than %v\n", guess)
		} else {
			fmt.Printf("Congratulations! You guessed the number in %v attempts\n", attempts)

			end := time.Now()
			fmt.Printf("You took %v seconds to complete the game!\n", strconv.FormatFloat(end.Sub(start).Abs().Seconds(), 'f', 1, 64))

			break
		}

		currentChances--

		if currentChances == 0 {
			fmt.Println("You didn't guess the number!")
			break
		}
	}
}

func getDifficultyLevelInput() (int, error) {
	var choice int

	fmt.Println("Please select the difficulty level:")

	fmt.Printf("1. Easy (%v chances)\n", difficultyLevels["Easy"])
	fmt.Printf("2. Medium (%v chances)\n", difficultyLevels["Medium"])
	fmt.Printf("3. Hard (%v chances)\n", difficultyLevels["Hard"])

	fmt.Print("\nEnter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		return difficultyLevels["Easy"], nil
	case 2:
		return difficultyLevels["Medium"], nil
	case 3:
		return difficultyLevels["Hard"], nil
	default:
		return 0, errors.New("Invalid choice.")
	}
}
