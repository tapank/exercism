package account

import "sync"

type Account struct {
	Mu  sync.Mutex
	Bal int64
	Ok  bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{Bal: amount, Ok: true}
}

func (a *Account) Balance() (int64, bool) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	if a.Ok {
		return a.Bal, a.Ok
	}
	return 0, false
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	if !a.Ok {
		return 0, false
	}
	bal := a.Bal + amount
	if bal < 0 {
		return 0, false
	}
	a.Bal = bal
	return bal, true
}

func (a *Account) Close() (int64, bool) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	if a.Ok {
		amt := a.Bal
		a.Bal, a.Ok = 0, false
		return amt, true
	}
	return 0, false
}
