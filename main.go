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
}
