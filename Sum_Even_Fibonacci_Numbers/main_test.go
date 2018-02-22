package main

import "testing"

var tables = []struct {
	input  int
	expect int
}{
	{8, 10},
	{111111, 60696},
	{8675309, 4613732},
	{96746, 60696},
	{144100000, 82790070},
	{65140000, 82790070},
	{7347000000, 6293134512},
	{10000000000000, 8583840088782},
	{123456789000000, 154030760585064},
	{2, 2},
	{1, 2},
}

func TestSumEvenFibonacci(t *testing.T) {

	for _, table := range tables {
		result := SumEvenFibonacci(table.input)
		if result != table.expect {
			t.Errorf("Even Fibonacci Sum of '%s' was incorrect got: %s, want: %s.", table.input, result, table.expect)
		}
	}
}
func TestIdealSumEvenFibonacci(t *testing.T) {

	for _, table := range tables {
		result := IdealSumEvenFibonacci(table.input)
		if result != table.expect {
			t.Errorf("Even Fibonacci Sum of '%s' was incorrect got: %s, want: %s.", table.input, result, table.expect)
		}
	}
}
