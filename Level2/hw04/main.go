//HW04
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	workers()
	exitOnSigterminal()
}

// Task 1
func workers() {
	var workers = make(chan struct{}, 1)
	var count int = 0

	for i := 1; i <= 1000; i++ {
		workers <- struct{}{}

		go func(job int) {
			defer func() {
				<-workers
			}()
			count++
		}(i)
	}

	time.Sleep(2 * time.Millisecond)
	fmt.Println("Final count: ", count)

}

// Task 2
func exitOnSigterminal() {

	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, syscall.SIGTERM)
	exitChan := make(chan int)

	go func() {
		for {
			s := <-signalChanel
			switch s {
			case syscall.SIGTERM:
				fmt.Println("Signal terminte triggered.")
				exitChan <- 0
			default:
				fmt.Println("Unknown signal.")
				exitChan <- 1
			}
		}
	}()
	exitCode := <-exitChan
	os.Exit(exitCode)
}
