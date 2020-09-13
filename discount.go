package main

import "fmt"

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func printInfo(s *subscriber) {
	fmt.Printf("Name   : %s\n", s.name)
	fmt.Printf("Rate   : %0.2f\n", s.rate)
	fmt.Printf("Active : %t\n\n", s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}
