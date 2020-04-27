// style1 separates data from behavior, so the AccountService is completely
// responsible for defining how to credit and debit, in addition to transferring
// between two accounts.

package style1

import (
	"errors"
)

type Amount int

type Balance struct {
	Amount Amount
}

type Account struct {
	balance Balance
	number  string
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

func (a CheckingAccount) Number() string {
	return a.number
}

func (a SavingsAccount) Number() string {
	return a.number
}

func (a *CheckingAccount) UpdateBalance(amount Amount) {
	a.balance.Amount += amount
}

func (a *SavingsAccount) UpdateBalance(amount Amount) {
	a.balance.Amount += amount
}

type Exchanger interface {
	Balance() Amount
	UpdateBalance(amount Amount)
	Number() string
}

type AccountTransferer interface {
	Credit(a Exchanger, amount Amount) Exchanger
	Debit(a Exchanger, amount Amount) (Exchanger, error)
	Transfer(from, to Exchanger, amount Amount) (Exchanger, Exchanger, error)
}

type AccountService struct{}

func (as AccountService) Credit(a Exchanger, amount Amount) Exchanger {
	a.UpdateBalance(amount)

	return a
}

func (as AccountService) Debit(a Exchanger, amount Amount) (Exchanger, error) {
	if a.Balance() < amount {
		return a, errors.New("Insufficient balance in account")
	}

	a.UpdateBalance(-amount)

	return a, nil
}

func (as AccountService) Transfer(from, to Exchanger, amount Amount) (Exchanger, Exchanger, error) {
	d, err := as.Debit(from, amount)
	if err != nil {
		return from, to, err
	}
	c := as.Credit(to, amount)
	if err != nil {
		return from, to, err
	}

	return d, c, nil
}

func NewCheckingAccount(balance Amount, name string) CheckingAccount {
	return CheckingAccount{
		Account: Account{
			balance: Balance{Amount: balance},
			number:  "a1",
		},
	}
}

func NewSavingsAccount(balance Amount, name string) SavingsAccount {
	return SavingsAccount{
		Account: Account{
			balance: Balance{Amount: balance},
			number:  "a1",
		},
	}
}
