package main

import (
	"fmt"
	"sync"
)

func main() {

	// fmt.Println(runtime.NumCPU())

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for c := 'a'; c <= 'z'; c += 1 {
			fmt.Printf("%c", c)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i <= 10; i += 1 {
			fmt.Printf("%d", i)
		}
	}()

	wg.Wait()

}
