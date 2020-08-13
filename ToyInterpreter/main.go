package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	strconv "strconv"
	"strings"
)

// map that takes the variables, each key is the variable name with a string value
var m = make(map[string]string)

// function for calculations
func calc(s []string) int {

	// first check if there are any variables already declared in the slice given
	// if there are then replace them with their associated value

	if contains(m, s[0]) {
		s[0] = m[s[0]]
	}

	if contains(m, s[2]) {
		s[2] = m[s[2]]
	}

	// convert the elements either side of the operator to integer values
	val1, _ := strconv.Atoi(s[0])
	val2, _ := strconv.Atoi(s[2])

	// switch case for operator
	switch s[1] {

	case "+":
		return val1 + val2

	case "-":
		return val1 - val2

	case "*":
		return val1 * val2

	case "/":
		return val1 / val2

	default:
		return 0
	}

}

// function to return boolean of whether key is in the map
func contains(m map[string]string, s string) bool {
	if _, ok := m[s]; ok {
		return true
	}

	return false
}

// function to read the slice of strings given from each line of the txt file
func readLines(s []string) {

	// if there is only one element in the slice
	if len(s) == 1 {

		// check if it's a variable, then print value
		if contains(m, s[0]) {

			fmt.Println(m[s[0]])
		} else {
			fmt.Println(s[0])
		}

		// if it's longer than one element, it will be declaration
	} else if s[1] == "=" {

		// if the length is more than three
		if len(s) > 3 {

			// for loop over slice, converting any variables into their stored value
			for i := 2; i < len(s); i++ {
				if contains(m, s[i]) {
					s[i] = m[s[i]]
				}
			}

			// then set the value of variable at start of slice to result of calculation
			m[s[0]] = strconv.Itoa(calc(s[2:]))

		} else {

			// otherwise set the value of the variable to value on right hand side
			// making sure any variables on right hand side are given their stored value
			if contains(m, s[2]) {

				m[s[0]] = m[s[2]]

			} else {
				m[s[0]] = s[2]
			}

		}

	}

}

func main() {

	// get file path from os arguments
	filePath := os.Args[1]

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// set up new scanner for file
	scanner := bufio.NewScanner(file)

	// declare string variable to take each line
	var line string

	// loop over file line by line
	for scanner.Scan() {
		line = scanner.Text()

		// create slice of strings by separating by space
		s := strings.Split(line, " ")

		// run slice through readLines function
		readLines(s)

	}
}
