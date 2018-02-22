package main

import (
	"fmt"
)

func Multiple3And5(number int) (m int) {

	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			m += i
		}
	}

	return
}

func main() {
	fmt.Println(Multiple3And5(10))
	fmt.Println(Multiple3And5(42))
	fmt.Println(Multiple3And5(123))
	fmt.Println(Multiple3And5(1050))
}
