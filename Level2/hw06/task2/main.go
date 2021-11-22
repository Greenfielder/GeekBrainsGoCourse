package main

import (
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func doSomething(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Microsecond)
}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	numberOfGorutine := 3
	wg := &sync.WaitGroup{}
	wg.Add(numberOfGorutine)

	for i := 0; i < numberOfGorutine; i++ {
		go doSomething(wg)
		runtime.Gosched()
	}
	runtime.Gosched()
	time.Sleep(100 * time.Nanosecond)
}
