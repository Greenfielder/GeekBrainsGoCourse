package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var operation string
	fmt.Println("Введите нужную операцию (+, -, *, /, ^)")
	fmt.Scan(&operation)

	fmt.Println("Введите два операнда через пробел: ")
	var operand1, operand2 float64
	fmt.Scan(&operand1, &operand2)

	var result float64
	switch operation {

	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Println("Делить на 0 нельзя!")
			os.Exit(1)
		} else {
			result = operand1 / operand2
		}
	case "^":
		result = math.Pow(operand1, operand2)
	default:
		fmt.Println("Введена неизвестная операция.")
		os.Exit(1)
	}
	fmt.Printf("Результат = %5.2f\n", result)
}
