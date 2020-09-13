package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//	"bufio"
	//	"os"
)

func GetStrings(filename string) ([]string, error) {
	fmt.Println(filename)
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func GetMapStrings(filename string) (map[string]int, error) {

	votes := make(map[string]int)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		votes[line]++
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return votes, nil
}

func main() {

	lines, err := GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	votes, err := GetMapStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(votes)

	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Println("------------------------------------")
	fmt.Println("myStruct 1 ")
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)

	myStruct.number = 21
	myStruct.word = "Tony"
	myStruct.toggle = true

	fmt.Println("------------------------------------")
	fmt.Println("myStruct 2 ")
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)

	fmt.Println("------------------------------------")
	fmt.Println("Subscriber")
	var mySubscriber subscriber
	mySubscriber.rate = 12.32
	applyDiscount(&mySubscriber)
	fmt.Println(mySubscriber)

	fmt.Println("------------------------------------")
	fmt.Println("pointer")
	var value int = 23
	var pointer *int = &value
	fmt.Println(*pointer)

	fmt.Println("------------------------------------")
	fmt.Println("pointer 2")
	var mySubscriber2 subscriber
	mySubscriber2.rate = 12.23
	var mySubscriber2Pointer *subscriber = &mySubscriber2
	// Must wrap pointer in parenthese
	(*mySubscriber2Pointer).rate = 45.34
	fmt.Println((*mySubscriber2Pointer).rate)

	mySubscriber2Pointer.rate = 99.33
	fmt.Println((*mySubscriber2Pointer).rate)

	fmt.Println("------------------------------------")
	fmt.Println("------------------------------------")
	fmt.Println("Default Subscriber")
	tonySub := defaultSubscriber("Tony McClay")
	applyDiscount(tonySub)
	printInfo(tonySub)

	fmt.Println("------------------------------------")
	maxineSub := defaultSubscriber("Maxine Seals")
	applyDiscount(maxineSub)
	printInfo(maxineSub)

	fmt.Println("------------------------------------")
	fmt.Println("------------------------------------")
	fmt.Println("Struct Literals ")
	kristenSub := subscriber{name: "Kristen Smith", rate: 45.3,
		active: false}
	printInfo(&kristenSub)

	angeloSub := subscriber{name: "Angelo McClay", rate: 2.32,
		active: true}
	printInfo(&angeloSub)

	fmt.Println("------------------------------------")
	fmt.Println("------------------------------------")
	fmt.Println("Address Struct Literals ")
	myAddress := address{street: "2893 Avon Ct", city: "Palm Harbor",
		state: "FL", postalCode: "34684"}
	fmt.Printf("current Address\n%v\n\n", myAddress)

	fmt.Println("------------------------------------")
	fmt.Println("------------------------------------")
	fmt.Println("Full kristenSub Subscriber ")
	kristenSub.homeAddress = myAddress

	fmt.Printf("Full Kristin Sub\n%v\n\n", kristenSub)

	fmt.Println("------------------------------------")
	fmt.Println("------------------------------------")
	fmt.Println("Employee Tony ")
	tonyEmployee := employee{name: "Tony McClay"}
	tonyEmployee.rate = 99.99
	tonyEmployee.active = false
	// anonymous fields - address from struct
	tonyEmployee.address.street = "14425 Vincennes"
	tonyEmployee.address.city = "Harvey"
	tonyEmployee.address.state = "IL"
	tonyEmployee.address.postalCode = "60426"
	fmt.Printf("Full Tony Employee\n%v\n\n", tonyEmployee)
	// anonymous fields - address from struct -
	// don't specify the address sub structure
	tonyEmployee.street = "2393 Avon Ct"
	tonyEmployee.city = "Palm Harbor"
	tonyEmployee.state = "FL"
	tonyEmployee.postalCode = "34684"
	fmt.Printf("Full Tony Employee 2\n%v\n\n", tonyEmployee)
}
