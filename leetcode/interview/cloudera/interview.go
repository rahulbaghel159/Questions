package cloudera

import (
	"errors"
	"fmt"
)

/*

Implement a ‘bank’ class.

The bank should keep track of the account balances for different accounts designated by numeric identifiers.

It should support the following operations:
	deposit/withdraw to a single account,
	transfer between two accounts, and
	calculate the total amount of money in the bank.

Opening a new account should be done by depositing or transferring money to a previously-unused account ID.
Accounts cannot have negative balances.
Regulations also state that no account may hold more than 10B units.

-> Bank
	-> Accounts
		-> ID
		-> Balance
-> unique account id
-> 0 < balance < 10B USD
-> APIs/Operations
	-> Account Holders/Customer
		-> Deposit
		-> Withdraw
		-> Transfer
	-> Bank Manager/Admin
		-> Total Amount

entities

-> Bank
	-> []Customer
	-> []Admin
	-> deposit(accountID int, amount float)
	-> withdraw(accountID int, amount float)
	-> transfer(fromAccountID int, toAccountID int, amount float)
-> Account
	-> ID
	-> Balance
-> Customer
	-> ID
	-> Account
-> Balance
	-> Amount
	-> Currency
-> Account ID Generator
	-> generateAccountID()
-> Admin
	-> ID
	-> getAmount(adminID int)

*/

type Bank struct {
	id         int
	title      string
	customer   map[int]*Account //mappping person's unique id
	account    map[int]*Account //mapping account id
	admin      map[int]Admin
	balance    Balance
	MaxBalance int
	MinBalance int
	generator  AccountIDGenerator
}

type Account struct {
	id       int
	balance  *Balance
	customer Customer
}

type Balance struct {
	amount   float64
	currency string
}

type Customer struct {
	id     int
	person Person
}

type Person struct {
	name     string
	uniqueID int
}

func (b Bank) createAccount(person Person) (Account, error) {
	if _, ok := b.customer[person.uniqueID]; ok {
		return *b.customer[person.uniqueID], errors.New("Account already exists")
	}

	accountID := b.generator.generateAccountID()

	account := Account{
		id: accountID,
		balance: &Balance{
			amount:   0,
			currency: "USD",
		},
		customer: Customer{
			id:     accountID, //customer id
			person: person,
		},
	}

	b.customer[person.uniqueID] = &account

	return account, nil
}

// not thread safe
func (b Bank) deposit(accountID int, amount float64) error {
	if _, ok := b.account[accountID]; !ok {
		return errors.New("Account does not exist")
	}

	if amount <= float64(b.MinBalance) || amount >= float64(b.MaxBalance) {
		return errors.New("Invalid Amount")
	}

	balance := b.account[accountID].balance

	if balance.amount+amount > float64(b.MaxBalance) {
		return errors.New("Account will have excess balance after deposit")
	}

	// update user balance
	balance.amount += amount
	b.account[accountID].balance = balance

	//update bank balance
	b.balance.amount += amount

	return nil
}

// not thread safe
func (b Bank) withdraw(accountID int, amount float64) error {
	if _, ok := b.account[accountID]; !ok {
		return errors.New("Account does not exist")
	}

	if amount <= float64(b.MinBalance) || amount >= float64(b.MaxBalance) {
		return errors.New("Invalid Amount")
	}

	balance := b.account[accountID].balance

	if amount > balance.amount {
		return errors.New("Not Enough Money")
	}

	//txn

	// update user balance
	balance.amount -= amount
	b.account[accountID].balance = balance

	//update bank balance
	b.balance.amount -= amount

	//commit

	return nil
}

// not thread safe
func (b Bank) transfer(fromAccountID int, toAccountID int, amount float64) error {
	if _, ok := b.account[fromAccountID]; !ok {
		return errors.New("Withdrawal account does not exist")
	}

	if _, ok := b.account[toAccountID]; !ok {
		return errors.New("Deposit account does not exist")
	}

	if amount <= float64(b.MinBalance) || amount >= float64(b.MaxBalance) {
		return errors.New("Invalid Amount")
	}

	fromBalance := b.account[fromAccountID].balance
	toBalance := b.account[toAccountID].balance

	if amount > fromBalance.amount {
		return errors.New("Not Enough Money")
	}

	if toBalance.amount+amount > float64(b.MaxBalance) {
		return errors.New("Account will have excess balance after deposit")
	}

	//start txn

	fromBalance.amount -= amount
	b.account[fromAccountID].balance = fromBalance

	//network

	toBalance.amount += amount
	b.account[toAccountID].balance = toBalance

	//commit

	return nil
}

type AccountIDGenerator struct {
	counter int
}

func (g AccountIDGenerator) generateAccountID() int {
	g.counter++

	return g.counter
}

type Admin struct {
	id int
}

func (b Bank) getAmount(adminID int) (Balance, error) {
	if _, ok := b.admin[adminID]; !ok {
		return Balance{}, errors.New("Invalid User")
	}

	return b.balance, nil
}

func main() {
	a := Admin{}

	b := Bank{
		id:    1,
		title: "Test",
		admin: map[int]Admin{
			1: a,
		},
		customer:   map[int]*Account{},
		account:    map[int]*Account{},
		MaxBalance: 10000000000,
		MinBalance: 0,
	}

	testPerson := Person{
		name:     "Rahul",
		uniqueID: 123,
	}

	testAccount, _ := b.createAccount(testPerson)
	fmt.Println(testAccount.id)

	b.deposit(testAccount.id, 10)
	fmt.Println(b.getAmount(1))
}
