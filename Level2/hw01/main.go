package main

import (
	"fmt"
	"green/files"
	"green/panics"
	"strconv"
)

func main() {
	var filesNumber int = 100000

	panics.ExamplePanicAndHandle()

	for i := 0; i < filesNumber; i++ {
		err := files.CreateFile("temp/hello.txt" + strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
		}
	}
}
