package lld

import "fmt"

/*
Design and implement an In Memory Loan EMI Calculator.

The code should have functionality to create users. Users can be either customer or admin. All users will have a unique identifier: username.
Admin can create a Loan in the system for a customer.
While creating a loan, admin_username, customer_username, principal amount, interest rate and time (loan tenure) need to be taken as input.
The interest for the loan is calculated by I = P*N*R where P is the principal amount, N is the number of years and R is the rate of interest. The total amount to repay will be A = P + I The amount should be paid back monthly in the form of EMIs. Each EMI = A/(N*12)
Users should be able to make EMI payments for their loan only.
Users should be able to fetch loan info for their loans only. Fetching a loan should return the loan info along with all the emi payments done and EMIs remaining[optional].
Admin should be able to fetch all loans for all customers.
All the functions should take username as one of the arguments, and all user level validation should happen against this username
*/
type Address struct {
	street_name string
	state       string
	country     string
	city        string
	pincode     string
}

type User struct {
	name      string
	user_name string
	address   Address
}

type Customer struct {
	details User
	loan    Loan
}

type Admin struct {
	details User
	system  *System
}

type System struct {
	customers map[string]Customer
	admins    map[string]Admin
	loans     []Loan
}

func (s *System) createCustomer(customer Customer) {
	s.customers[customer.details.user_name] = customer
}

func (s *System) createAdmin(admin Admin) {
	s.admins[admin.details.user_name] = admin
}

type Loan struct {
	admin_username    string
	customer_username string
	principal         int
	rate              int
	tenure            int
	interest          int
	installments      []EMI
}

type EMI struct {
	amount  int
	duedate int
	status  bool
}

func (a *Admin) createLoan(admin_username, customer_username string, principal int, rate int, tenure int) {
	if a.ValidateAdmin(admin_username) {
		fmt.Println("Invalid admin")
		return
	}

	if a.ValidateCustomer(customer_username) {
		fmt.Println("Invalid customer")
		return
	}

	loan := Loan{
		admin_username:    admin_username,
		customer_username: customer_username,
		principal:         principal,
		rate:              rate,
		tenure:            tenure,
		interest:          principal * rate * tenure,
	}

	loan.installments = make([]EMI, tenure)
	emi := (principal + loan.interest) / (tenure * 12)

	for i := 0; i < tenure; i++ {
		loan.installments[i] = EMI{
			amount:  emi,
			duedate: i,
		}
	}

	customer := a.system.customers[customer_username]

	customer.loan = loan
	a.system.loans = append(a.system.loans, loan)
}

func (c *Customer) MakePayment(user_name string) {
	if c.details.user_name != user_name {
		fmt.Println("Invalid customer")
		return
	}

	installment_index := -1
	for i := 0; i < c.loan.tenure; i++ {
		if !c.loan.installments[i].status {
			installment_index = i
		}
	}

	if installment_index == -1 {
		fmt.Println("loan already paid")
		return
	}

	c.loan.installments[installment_index].status = true
}

func (c *Customer) FetchLoanInfo(user_name string) Loan {
	return c.loan
}

func (a *Admin) FetchLoanInfo() []Loan {
	return a.system.loans
}

func (a *Admin) ValidateCustomer(user_name string) bool {
	if _, ok := a.system.customers[user_name]; !ok {
		fmt.Println("Invalid customer")
		return false
	}

	return true
}

func (a *Admin) ValidateAdmin(user_name string) bool {
	if a.details.user_name != user_name {
		fmt.Println("Invalid admin")
		return false
	}

	return true
}
