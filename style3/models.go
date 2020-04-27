// style3 usess functions that return new values instead of using pointer
// receivers to mutate state.
package style3

import "errors"

type Amount int

type Balance struct {
	Amount Amount
}

type Account struct {
	balance Balance
	Number  string
}

type CheckingAccount struct {
	Account
}

type SavingsAccount struct {
	Account
}

func (a CheckingAccount) Balance() Amount {
	return a.balance.Amount
}

func (a SavingsAccount) Balance() Amount {
	return a.balance.Amount
}

type AccountService struct{}

type Exchanger interface {
	Credit(amount Amount) Exchanger
	Debit(amount Amount) (Exchanger, error)
	Balance() Amount
}

func (a CheckingAccount) Credit(amount Amount) Exchanger {
	return CheckingAccount{Account: Account{Balance{a.Balance() + amount}, a.Number}}
}

func (a CheckingAccount) Debit(amount Amount) (Exchanger, error) {
	if a.Balance() < amount {
		return &a, errors.New("Insufficient balance in account")
	}

	return CheckingAccount{Account: Account{Balance{a.Balance() - amount}, a.Number}}, nil
}

func (a SavingsAccount) Credit(amount Amount) Exchanger {
	return SavingsAccount{Account: Account{Balance{a.Balance() + amount}, a.Number}}
}

func (a SavingsAccount) Debit(amount Amount) (Exchanger, error) {
	if a.Balance() < amount {
		return &a, errors.New("Insufficient balance in account")
	}

	return SavingsAccount{Account: Account{Balance{a.Balance() - amount}, a.Number}}, nil
}

func (as AccountService) Transfer(from Exchanger, to Exchanger, amount Amount) (Exchanger, Exchanger, error) {
	d, err := from.Debit(amount)
	if err != nil {
		return from, to, err
	}
	c := to.Credit(amount)

	return d, c, nil
}

func NewCheckingAccount(balance Amount, name string) CheckingAccount {
	return CheckingAccount{
		Account: Account{
			balance: Balance{Amount: balance},
			Number:  "a1",
		},
	}
}

func NewSavingsAccount(balance Amount, name string) SavingsAccount {
	return SavingsAccount{
		Account: Account{
			balance: Balance{Amount: balance},
			Number:  "a1",
		},
	}
}
