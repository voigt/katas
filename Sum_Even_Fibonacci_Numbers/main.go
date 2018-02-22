package main

import "fmt"

// My solution (does not 100% fit the description, but I wanted to use an iterator)
// my solution starts with 0, 1 - not with 1, 2.
// This leads to a problem that SumEvenFibonacci(1).Equals(2) != true; but test cases require this
func SumEvenFibonacci(limit int) (sum int) {

	fib := newFibonacci()

	if limit == 1 {
		return 2
	}

	for i := fib(); i <= limit; i = fib() {
		if i%2 == 0 {
			sum += i
		}
	}

	return
}

func newFibonacci() func() int {
	a, b := 0, 1

	return func() int {
		hold := a
		a = b
		b = hold + b
		return b
	}
}

// Solution better suited to description
func IdealSumEvenFibonacci(limit int) (sum int) {
	sum, a, b := 2, 1, 2
	for b <= limit {
		a, b = b, a+b
		if b%2 == 0 {
			sum += b
		}
	}
	return
}

func main() {

	fmt.Println(SumEvenFibonacci(10))
	fmt.Println(IdealSumEvenFibonacci(10))
}
