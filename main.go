package main

import (
	"errors"
	"fmt"
)

type Account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
}

// deposit method
func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("cannot deposit zero or negative amount")
	}
	acc.Balance += amount
	fmt.Printf("Deposited  %.2f. New balance is %.2f to  %s\n", amount, acc.Balance, acc.OwnerName)
	return nil
}

// withdrawal method
func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("cannot withdraw zero or negative amount")
	}

	if acc.Balance < amount {
		return fmt.Errorf("Insufficient balance! You tried to withdraw %.2f. But you remaining balance was %.2f\n", amount, acc.Balance)
	}

	acc.Balance -= amount
	fmt.Printf("Withdrew %.2f remaining balance of %.2f to %s\n", amount, acc.Balance, acc.OwnerName)
	return nil
}

// get Balance method
func (acc *Account) GetBalance() float64 {
	return acc.Balance
}

// Stringer Interface
func (acc *Account) String() string {
	return fmt.Sprintf("Account: [%s], Owner : [%s], Balance: [%.2f]", acc.AccountNumber, acc.OwnerName, acc.Balance)
}

// Savings Account
type SavingsAccount struct {
	Account
	InterestRate float64
}

func (sa *SavingsAccount) AddInterest() {
	interest := sa.InterestRate / 100 * sa.Balance //access promoted balance field
	fmt.Printf("Adding interest: %.2f to Savings Account: %s", interest, sa.AccountNumber)
	err := sa.Deposit(interest)
	if err != nil {
		fmt.Printf("AddInterest : Error depositing interest %.2f to Savings Account: %s\n", interest, sa.AccountNumber)
	}
}

// overdraft account
type OverdraftAccount struct {
	Account
	OverdraftLimit float64
}

func (oa *OverdraftAccount) WithDraw(amount float64) error {
	if amount <= 0 {
		return errors.New("cannot withdraw zero or negative overdraft amount")
	}
	//Allow overdraft up to balance and overdraft limit
	if (oa.Balance + oa.OverdraftLimit) < amount {
		return fmt.Errorf("Withdrawal of %.2f, exceeds the limit of %.2f", amount, oa.OverdraftLimit)
	}
	//allow overdraft
	oa.Balance -= amount
	fmt.Printf("Withdrew %.2f from overdraft account account %s, New Balance: %.2f\n", amount, oa.AccountNumber, oa.Balance)
	return nil
}

func main() {

	fmt.Printf("----Savings Account----\n")
	account := SavingsAccount{
		Account: Account{
			AccountNumber: "SAVINGS",
			Balance:       3500,
			OwnerName:     "Joel",
		},
		InterestRate: 11.0,
	}

	err := account.Deposit(200)
	if err != nil {
		fmt.Printf("Error depositing savings account: %s\n", err)
		return
	}

	account.AddInterest()

	err = account.Withdraw(200)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("------ Overdraft  Account------\n")
	account1 := OverdraftAccount{
		Account: Account{
			AccountNumber: "OVERDRAFT",
			Balance:       3500,
			OwnerName:     "Joel",
		},
		OverdraftLimit: 20,
	}

	err = account1.WithDraw(5100)
	if err != nil {
		fmt.Printf("Error %s\n", err)
		return
	}

}
