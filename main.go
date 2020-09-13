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

	fmt.Println("------------------------------------")
	fmt.Println("------------------------------------")
	fmt.Println(" Car and Bus Fuel")
	var carFuel gallons
	var busFuel liters
	carFuel = gallons(10.0)
	busFuel = liters(240.0)
	fmt.Printf("Car Fuel : %.2f\n", carFuel)
	fmt.Printf("Bus Fuel : %0.2f\n\n", busFuel)

	fmt.Println("------------------------------------")
	fmt.Println(" Car 2 and Bus 2 Fuel")
	// Short Definitions
	carFuel2 := gallons(5.55)
	busFuel2 := liters(440.0)
	fmt.Printf("Car Fuel 2: %.2f\n", carFuel2)
	fmt.Printf("Bus Fuel 2: %0.2f\n\n", busFuel2)

	fmt.Println("------------------------------------")
	fmt.Println(" Car 3 and Bus 3 Fuel")
	// Addinid structures
	carFuel3 := carFuel + carFuel2
	busFuel3 := busFuel + busFuel2
	fmt.Printf("Car Fuel 3: %.2f\n", carFuel3)
	fmt.Printf("Bus Fuel 3: %0.2f\n\n", busFuel3)

	// . These are current errors
	// carFuel4 := gallons(1.2)
	// busFuel4 := liters(2.5)
	// carFuel5 := gallons(8.0)
	// busFuel5 := liters(30.0)

	fmt.Println("------------------------------------")
	fmt.Println(" Car 4 and Bus 4 Fuel")
	carFuel4 := gallons(1.2)
	busFuel4 := liters(4.5)
	carFuel4 += toGallons(liters(40.0))
	busFuel4 += toLiters(gallons(30.0))
	fmt.Printf("Car Fuel 4: %.2f\n", carFuel4)
	fmt.Printf("Bus Fuel 4: %0.2f\n\n", busFuel4)

	fmt.Println("------------------------------------")
	fmt.Println(" Method type ")
	value2 := myType("Lisa McClay")
	value2.sayHi()

	value3 := myType("Angelo McClay")
	value3.sayHi()

	fmt.Println("------------------------------------")
	fmt.Println(" number double ")
	aNum := number(4)
	fmt.Printf("number : %d\n", aNum)
	aNum.double()
	fmt.Printf("number : %d\n", aNum)
	aNum.double()
	fmt.Printf("number : %d\n", aNum)
	aNum.double()
	fmt.Printf("number : %d\n", aNum)
	aNum.double()
	fmt.Printf("number : %d\n", aNum)
	aNum.double()
	fmt.Printf("number : %d\n", aNum)

	fmt.Println("------------------------------------")
	fmt.Println(" Date Struct ")
	aDate := date{year: 2020, month: 9, day: 13}
	fmt.Println(aDate)

	aDate.setYear(1969)
	aDate.setMonth(19)
	aDate.setDay(69)
	fmt.Println(aDate)

	err = aDate.setYear(2034)
	if err != nil {
		log.Fatal(err)
	}
	err = aDate.setMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = aDate.setDay(31)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(aDate)

}
