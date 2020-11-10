package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(val Bitcoin) {
	w.balance += val
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}


var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(val Bitcoin) error {

	if val > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= val
	return nil
}
