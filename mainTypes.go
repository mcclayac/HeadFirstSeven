package main

type subscriber struct {
	rate        float64
	name        string
	active      bool
	homeAddress address
}

// anonymous fields - address
type employee struct {
	rate   float64
	name   string
	active bool
	address
}

type address struct {
	street     string
	city       string
	state      string
	postalCode string
}

type liters float64
type gallons float64

type myType string
type number int
