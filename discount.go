package main

import (
	"errors"
	"fmt"
)

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

func toGallons(l liters) gallons {
	return gallons(l * 0.264)
}

func toLiters(g gallons) liters {
	return liters(g * 3.785)
}

func (m myType) sayHi() {
	fmt.Println("say hi, ", m)
}

func (n *number) double() {
	*n *= 2
}

// setter methods
func (d *date) setYear(year int) error {
	if year < 1 {
		return errors.New("Invalid Year")
	}
	d.year = year
	return nil
}

func (d *date) setMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid month")
	}
	d.month = month
	return nil
}

func (d *date) setDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid day")
	}
	d.day = day
	return nil
}
