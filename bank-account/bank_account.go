package account

import "sync"

type Account struct {
	sync.Mutex
	Amount int64
	Open   bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{Amount: amount, Open: true}
}

func (a *Account) Balance() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.Open {
		return a.Amount, a.Open
	} else {
		return 0, false
	}
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if a.Open {
		a.Lock()
		defer a.Unlock()
		a.Amount += amount
		if a.Amount < 0 {
			a.Amount = 0
			return 0, false
		}
		return a.Amount, a.Open
	} else {
		return 0, false
	}
}
func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	// If account is Open -> close it
	// else return false for the operation and leave it closed.
	if a.Open {
		a.Open = false
		return a.Amount, !a.Open
	} else {
		a.Amount = 0
		return a.Amount, a.Open
	}
}
