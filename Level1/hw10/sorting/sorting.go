package sorting

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SortByInsert() {
	var numbers []int
	fmt.Println("Введите числа через пробел: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbers = stringToInt(scanner.Text())

	sortingByInsert(numbers)
	fmt.Println(numbers)

}

func stringToInt(s string) []int {
	var sliceOfNumbers []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			sliceOfNumbers = append(sliceOfNumbers, i)
		}
	}
	return sliceOfNumbers
}

func sortingByInsert(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		number := numbers[i]
		j := i
		for ; j >= 1 && numbers[j-1] > number; j-- {
			numbers[j] = numbers[j-1]
		}
		numbers[j] = number
	}
	return numbers
}
