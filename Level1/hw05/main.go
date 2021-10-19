package main

import (
	"fmt"
	"os"
)

// Home work 5
// Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
// Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе

func fibonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonachi(n-1) + fibonachi(n-2)
}

func main() {
	var number uint
	fibonachiNumbers := make(map[uint]uint)

	for {
		fmt.Println("Введите число (для выхода введите 111): ")
		fmt.Scan(&number)

		if number == 111 {
			os.Exit(1)
		} else if _, ok := fibonachiNumbers[number]; ok {
			fmt.Println("Для числа ", number, " число фибоначчи: ", fibonachiNumbers[number], " (взято из памяти)")
		} else {
			fibonachiNumbers[number] = fibonachi(number)
			fmt.Println("Для числа ", number, " число фибоначчи: ", fibonachi(number))
		}
	}
}
