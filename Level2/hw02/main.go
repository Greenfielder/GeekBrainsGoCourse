// This is simple console programm to get the fibonachi number.
// At startup, the program waits for input from the user.
// After entering the number, the result is displayed.
// To exit the program, you must enter: 111.

package main

import (
	"fmt"
	"green/fibonachi"
	"os"
)

func main() {
	var number uint

	for {
		fmt.Println("Введите число дял получения числа фибоначчи (для выхода введите 111): ")
		fmt.Scan(&number)

		if number == 111 {
			os.Exit(1)
		} else {
			fmt.Println("Для числа ", number, " число фибоначчи: ", fibonachi.FibonachiCached(number))
		}
	}
}
