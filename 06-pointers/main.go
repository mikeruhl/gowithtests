package main

//Wallet is a wallet
type Wallet struct {
	balance int
}

//Deposit adds to the balance
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

//Balance gets balance of wallet
func (w *Wallet) Balance() int {
	return w.balance
}

func main() {}
