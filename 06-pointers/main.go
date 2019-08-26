package main

import (
	"errors"
	"fmt"
)

//Wallet is a wallet
type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

//Deposit adds to the balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Withdraw subtracts from the balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

//Balance gets balance of wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func main() {}
