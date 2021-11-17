package main

import (
	"fmt"
	"green/mymutex"
	"log"
	"sync"
	"time"
)

func doConcurrent(n int) {
	wg := &sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, number int) {
			defer wg.Done()
			time.Sleep(time.Second)
			log.Printf("Goroutine #%d is done", number)
		}(wg, i)
	}
	wg.Wait()
}

func main() {

	// 1. Напишите программу, которая запускает n потоков и дожидается завершения их всех
	doConcurrent(100)

	// 2. Реализуйте функцию для разблокировки мьютекса с помощью defer
	var ms mymutex.Kit
	ms = mymutex.NewRWMutexKit()
	ms.Add(17)
	fmt.Println("Set has 7?", ms.Has(7))
	fmt.Println("Set has 17?", ms.Has(17))
}
