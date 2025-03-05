package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
	rwmu    sync.RWMutex
}

func NewAccount(balance int) *BankAccount {
	return &BankAccount{balance: balance}
}
func (b *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
	b.mu.Lock()
	defer func() {
		b.mu.Unlock()
		wg.Done()
	}()
	b.balance += amount
	fmt.Println("Depositing", amount)
}
func (b *BankAccount) Balance(wg *sync.WaitGroup) {

	b.rwmu.RLock()
	fmt.Println("Balance", b.balance)
	b.rwmu.RUnlock()
	wg.Done()

}
func (b *BankAccount) Withdraw(amount int, wg *sync.WaitGroup) {
	b.mu.Lock()
	if b.balance >= amount {
		b.balance -= amount
	} else {
		fmt.Println("Failed Withdrawing Insufficient balance", b.balance)
	}
	b.mu.Unlock()
	wg.Done()
}
func main() {
	bank := NewAccount(100)
	var wg sync.WaitGroup
	wg.Add(4)
	go bank.Deposit(10, &wg)
	go bank.Balance(&wg)
	go bank.Withdraw(100, &wg)
	go bank.Withdraw(200, &wg)
	wg.Wait()
	fmt.Println(bank.balance)
}
