package main

import (
	"fmt"
	"log"
	"sync"
	"time"
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

func (bc *bankClient) add(amount int) {
	bc.Lock()
	bc.balance += amount
	time.Sleep(time.Second)
	bc.Unlock()
}

func (bc *bankClient) debit(amount int) {
	bc.Lock()
	bc.balance -= amount
	time.Sleep(time.Second)
	bc.Unlock()
}

func (bc *bankClient) getBalance() int {
	bc.Lock()
	defer bc.Unlock()

	return bc.balance
}

func doManeyOrder(bc *bankClient, m moneyOrder) {
	if m.writeOff {
		bc.debit(m.amount)
	} else {
		bc.add(m.amount)
	}
}

func main() {

	// trace.Start(os.Stderr)
	// defer trace.Stop()

	client, err := newBankClient("John", 100000)
	if err != nil {
		log.Fatalln(err)
	}

	order := moneyOrder{23000, true}
	doManeyOrder(client, order)

	fmt.Println(client.getBalance())

}
