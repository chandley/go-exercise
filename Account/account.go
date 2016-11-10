package main

import "fmt"

func main() {
	var myAccount = currentAccount{100}
	fmt.Println(myAccount.Deposit(10))
	fmt.Println(myAccount.Withdraw(20))
}

type currentAccount struct {
	balance int
}

func (c currentAccount) Withdraw(amount int) int {
	c.balance -= amount
   return c.balance
}

func (c currentAccount) Deposit(amount int) int {
	c.balance += amount
	return c.balance
}

type Account interface {

// Withdraw deducts amount from account returning the new balance
Withdraw(amount int) (newBalance int)

// Deposit adds amount to the account returning the new balance
Deposit(amount int) (newBalance int)
}
