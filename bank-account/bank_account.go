package account

import "sync"

// Define the Account type here.
type Account struct {
	mu     sync.Mutex
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
	a.mu.Lock()
	if a.Open {
		a.mu.Unlock()
		return a.Amount, a.Open
	} else {
		a.mu.Unlock()
		return 0, false
	}
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if a.Open {
		a.mu.Lock()
		a.Amount += amount
		if a.Amount < 0 {
			a.Amount = 0
			a.mu.Unlock()
			return 0, false
		}
		a.mu.Unlock()
		return a.Amount, a.Open
	} else {
		return 0, false
	}
}
func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	// If account is Open -> close it
	// else return false for the operation and leave it closed.
	if a.Open {
		a.Open = false
		a.mu.Unlock()
		return a.Amount, !a.Open
	} else {
		a.Amount = 0
		a.mu.Unlock()
		return a.Amount, a.Open
	}
}
