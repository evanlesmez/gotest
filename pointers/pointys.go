package main

// On refactor
// https://github.com/quii/learn-go-with-tests/blob/master/pointers-and-errors.md
import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Depost(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
