package main

import (
	"fmt"
	"sync"
)

var value int = 0

//Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”
func main() {
	wg := &sync.WaitGroup{}
	counter := 400
	wg.Add(counter)

	for i := 0; i < counter; i++ {
		go func() {
			defer wg.Done()
			value++
		}()
	}
	wg.Wait()
	fmt.Println(value)
}
