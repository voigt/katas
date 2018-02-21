package main

import (
	"fmt"
	"strconv"
)

// my attempt
func century(year int) int {

	if year%100 == 0 {
		return (year / 100)
	}

	return (year / 100) + 1
}

// very clever attempt
func cleverCentury(year int) int {
	return (year + 99) / 100
}

func main() {
	fmt.Println(strconv.Itoa(century(1705)))
	fmt.Println(strconv.Itoa(century(1900)))
	fmt.Println(strconv.Itoa(century(1601)))
	fmt.Println(strconv.Itoa(century(2000)))

	fmt.Println(strconv.Itoa(cleverCentury(1705)))
	fmt.Println(strconv.Itoa(cleverCentury(1900)))
	fmt.Println(strconv.Itoa(cleverCentury(1601)))
	fmt.Println(strconv.Itoa(cleverCentury(2000)))
}
