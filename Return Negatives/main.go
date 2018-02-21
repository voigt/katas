package main

import (
	"fmt"
	"strconv"
)

func MakeNegative(x int) int {
	if x > 0 {
		x *= -1
	}
	return x
}

func main() {
	fmt.Println(strconv.Itoa(MakeNegative(1)))
	fmt.Println(strconv.Itoa(MakeNegative(-5)))
	fmt.Println(strconv.Itoa(MakeNegative(0)))
}
