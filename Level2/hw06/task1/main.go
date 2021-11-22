package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
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

// Добавляем необходимую сумму клиенту
func (bc *bankClient) add(amount int) {
	bc.Lock()
	bc.balance += amount
	time.Sleep(time.Second)
	bc.Unlock()
}

// Списываем необходимую сумму у клиента
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

// Создаем денежный перевод
func doManeyOrder(bc *bankClient, m moneyOrder, wg *sync.WaitGroup) {
	if m.writeOff {
		bc.debit(m.amount)
	} else {
		bc.add(m.amount)
	}
	wg.Done()
}

// Задание 1. Написать программу, которая использует мьютекс для безопасного доступа к данным из нескольких потоков.
//Выполните трассировку программы.
func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	client, err := newBankClient("John Morrison", 100000)
	if err != nil {
		log.Fatalln(err)
	}

	moneyOrders := []moneyOrder{
		{23000, true},
		{11000, true},
		{3000, true},
		{25000, false},
		{7000, true},
		{2000, true},
		{15000, true},
		{20000, false},
		{5000, true},
		{17000, true},
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(moneyOrders))
	for _, m := range moneyOrders {
		go doManeyOrder(client, m, wg)
	}

	wg.Wait()
	fmt.Println("Our client named", client.clientName, "Have: ", client.getBalance())

}
