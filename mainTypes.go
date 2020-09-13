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
