package lld

type Address struct {
	street_name string
	city        string
	country     string
	pincode     string
}

type Person struct {
	name    string
	address Address
}

type Account struct {
	user_name string
	password  string
}

type Discount struct {
	title      string
	percentage int
	amount     int
}

type Admin struct {
	person  Person
	account Account
}

type Customer struct {
	person   Person
	account  Account
	discount []Discount
}

type Manager struct {
	person  Person
	account Account
}

type City struct {
	name    string
	cinemas []Cinema
}

type Cinema struct {
	name   string
	halls  []Hall
	movies []Movie
}

type Hall struct {
	name string
	show []Show
}

// seat status
const (
	AVAILABLE = iota
	RESERVED
)

// seat type/class
const (
	STANDARD     = "standard"
	DELUXE       = "deluxe"
	SUPER_DELUXE = "super"
)

type Seats struct {
	id     string
	status int
	class  string
	cost   int
}

type Show struct {
	seats    []Seats
	cost     int
	movie    Movie
	discount []Discount
}

type Movie struct {
	title        string
	language     string
	genre        string
	release_date string
}

type Ticket struct {
	customer     Customer
	movie        Movie
	show         Show
	seats        []Seats
	amount       int
	discount     int
	final_amount int
}

type Payment struct {
	ticket        Ticket
	status        int
	transactionID string
}

const (
	PENDING = iota
	AUTHORISED
	REJECTED
	SETTLED
)

type CreditCard struct {
	details string //hashed
}

type Transaction struct {
	id         string
	payment    Payment
	status     int
	created_at string
	updated_at string
}
