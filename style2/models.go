// style2 defines behavior that acts on a single type on that type and only
// separates behavior that works on different or multiple types.

package style2

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

func (a *CheckingAccount) UpdateBalance(amount Amount) {
	a.balance.Amount += amount
}

func (a *SavingsAccount) UpdateBalance(amount Amount) {
	a.balance.Amount += amount
}

type AccountService struct{}

type Exchanger interface {
	Credit(amount Amount) Exchanger
	Debit(amount Amount) (Exchanger, error)
	Balance() Amount
	UpdateBalance(amount Amount)
}

func (a CheckingAccount) Credit(amount Amount) Exchanger {
	a.UpdateBalance(amount)

	return &a
	//return CheckingAccount{Account: Account{Balance{a.Balance() + amount}, a.Number}}, nil
}

func (a CheckingAccount) Debit(amount Amount) (Exchanger, error) {
	if a.Balance() < amount {
		return &a, errors.New("Insufficient balance in account")
	}

	a.UpdateBalance(-amount)

	return &a, nil
	//return CheckingAccount{Account: Account{Balance{a.Balance() - amount}, a.Number}}, nil
}

func (a SavingsAccount) Credit(amount Amount) Exchanger {
	a.UpdateBalance(amount)

	return &a
	//return SavingsAccount{Account: Account{Balance{a.Balance() + amount}, a.Number}}, nil
}

func (a SavingsAccount) Debit(amount Amount) (Exchanger, error) {
	if a.Balance() < amount {
		return &a, errors.New("Insufficient balance in account")
	}

	a.UpdateBalance(-amount)

	return &a, nil
	//return SavingsAccount{Account: Account{Balance{a.Balance() - amount}, a.Number}}, nil
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
