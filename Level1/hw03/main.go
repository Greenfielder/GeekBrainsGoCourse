package main

import (
	"fmt"
	"os"
)

func main() {

	var operation string
	fmt.Scan(&operation)
	fmt.Println("Введите два операнда через пробел: ")
	var operand1, operand2 float64
	fmt.Scan(&operand1, &operand2) // validate

	var result float64
	switch operation {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		result = operand1 / operand2 //validate /0
	default:
		fmt.Println("Введена неизвестная операция.")
		os.Exit(1)
	}
	fmt.Printf("Результат = %5.2f\n", result)
}
