package main

import (
	"fmt"
	"sync"
)

type Account struct {
	balance int
	mu      sync.Mutex
}

func (a *Account) Deposit(amount int) {
	a.mu.Lock()
	a.balance += amount
	a.mu.Unlock()
}

func (a *Account) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func main() {
	acc := Account{balance: 10}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			acc.Deposit(1)
		}()
	}

	wg.Wait()
	fmt.Println("Final Balance :", acc.Balance())
}
