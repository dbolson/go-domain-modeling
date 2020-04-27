package main

import (
	"fmt"
	"scratch-go/style1"
	"scratch-go/style2"
	"scratch-go/style3"
)

func style1Version() {
	as := style1.AccountService{}
	a1 := style1.NewCheckingAccount(100, "a1")
	a2 := style1.NewSavingsAccount(100, "a2")

	d, c, err := as.Transfer(&a1, &a2, 100)
	if err != nil {
		fmt.Println("error transferring", err.Error())
	}
	fmt.Println("debited", d.Balance())
	fmt.Println("credited", c.Balance())
}

func style2Version() {
	as := style2.AccountService{}
	a1 := style2.NewCheckingAccount(100, "a3")
	a2 := style2.NewSavingsAccount(100, "a4")

	d, c, err := as.Transfer(&a1, &a2, 100)
	if err != nil {
		fmt.Println("error transferring", err.Error())
	}
	fmt.Println("debited", d.Balance())
	fmt.Println("credited", c.Balance())
}

func style3Version() {
	as := style3.AccountService{}
	a1 := style3.NewCheckingAccount(100, "a3")
	a2 := style3.NewSavingsAccount(100, "a4")

	d, c, err := as.Transfer(a1, a2, 100)
	if err != nil {
		fmt.Println("error transferring", err.Error())
	}
	fmt.Println("debited", d.Balance())
	fmt.Println("credited", c.Balance())
}

func main() {
	style1Version()
	fmt.Println("----------")
	style2Version()
	fmt.Println("----------")
	style3Version()
}
