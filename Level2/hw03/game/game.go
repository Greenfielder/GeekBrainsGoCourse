package game

import (
	"fmt"
	"math/rand"
)

func StartGame() {
	count := 0
	guess := -1
	number := rand.Intn(10)

	for count < 3 && guess != number {
		fmt.Print("Отгадайте число в промежутке (0..9): ")
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
