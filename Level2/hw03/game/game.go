package game

import (
	"fmt"
	"math/rand"
	"time"
)

// function randomInt is a random int generator
func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min > max {
		return min
	} else {
		return rand.Intn(max-min) + min
	}
}

// StartGame function, this is the logic of "guest number game"
func StartGame() {
	count := 0
	guess := -1
	number := randomInt(1, 11)

	for count < 3 && guess != number {
		fmt.Println("Отгадайте число в промежутке (0..10): ")
		fmt.Scanln(&guess)
		if guess != number {

			if guess > number {
				fmt.Println("Загаданное число меньше")
			} else if guess < number {
				fmt.Println("Загаданное число больше")
			}
			count++

		} else {
			fmt.Println("Ты выиграл! :)")
		}

		if count == 3 {
			fmt.Println("Ты проиграл... :(")
		}
	}
}
