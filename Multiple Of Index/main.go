package main

import (
	"fmt"
)

// My solution
func multipleOfIndex(ints []int) (multiples []int) {

	for i, num := range ints {
		if (i > 0) && (num%i == 0) {
			multiples = append(multiples, num)
		}
	}

	return
}

// Doing with a for instead of range
func OtherMultipleOfIndex(ints []int) []int {
	// good luck
	ret := make([]int, 0)

	for i := 1; i < len(ints); i++ {
		if ints[i]%i == 0 {
			ret = append(ret, ints[i])
		}
	}
	return ret
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	printSlice(multipleOfIndex([]int{22, -6, 32, 82, 9, 25}))
	printSlice(multipleOfIndex([]int{68, -1, 1, -7, 10, 10}))
	printSlice(multipleOfIndex([]int{11, -11}))
	printSlice(multipleOfIndex([]int{-56, -85, 72, -26, -14, 76, -27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68}))
	printSlice(multipleOfIndex([]int{28, 38, -44, -99, -13, -54, 77, -51}))
	printSlice(multipleOfIndex([]int{-1, -49, -1, 67, 8, -60, 39, 35}))
	printSlice(multipleOfIndex([]int{22, -6, 32, 82, 9, 25}))

	printSlice(OtherMultipleOfIndex([]int{22, -6, 32, 82, 9, 25}))
	printSlice(OtherMultipleOfIndex([]int{68, -1, 1, -7, 10, 10}))
	printSlice(OtherMultipleOfIndex([]int{11, -11}))
	printSlice(OtherMultipleOfIndex([]int{-56, -85, 72, -26, -14, 76, -27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68}))
	printSlice(OtherMultipleOfIndex([]int{28, 38, -44, -99, -13, -54, 77, -51}))
	printSlice(OtherMultipleOfIndex([]int{-1, -49, -1, 67, 8, -60, 39, 35}))
	printSlice(OtherMultipleOfIndex([]int{22, -6, 32, 82, 9, 25}))

}
