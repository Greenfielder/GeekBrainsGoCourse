package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

type bankClient struct {
	sync.Mutex
	clientName string
	balance    int
}

type moneyOrder struct {
	amount   int
	writeOff bool
}

func newBankClient(name string, sun int) (*bankClient, error) {
	return &bankClient{clientName: name, balance: sun}, nil
}

func doManeyOrder(m moneyOrder) {

	if m.writeOff {
		fmt.Println("списываем")
	} else {
		fmt.Println("начисляем")
	}

}

func main() {

	trace.Start(os.Stderr)
	defer trace.Stop()

}
