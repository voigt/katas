package main

import (
	"bytes"
	"fmt"
	"strings"
)

// my attempt
func RepeatStr(repititions int, value string) (result string) {

	for j := 0; j < repititions; j++ {
		result += value
	}

	return
}

// very clever attempt
func CleverRepeatStr(repititions int, value string) string {
	// interesting: check how strings.Repeat works!
	return strings.Repeat(value, repititions)
}

// quite cool and performant way
// read http://herman.asia/efficient-string-concatenation-in-go
// or   http://golang-examples.tumblr.com/post/86169510884/fastest-string-contatenation
func CoolRepeatStr(repititions int, value string) string {
	var buffer bytes.Buffer

	for i := 0; i < repititions; i++ {
		buffer.WriteString(value)
	}
	return buffer.String()
}

func main() {
	fmt.Println(RepeatStr(4, "a"))
	fmt.Println(RepeatStr(3, "hello "))
	fmt.Println(RepeatStr(2, "abc"))

	fmt.Println(CleverRepeatStr(4, "a"))
	fmt.Println(CleverRepeatStr(3, "hello "))
	fmt.Println(CleverRepeatStr(2, "abc"))

	fmt.Println(CoolRepeatStr(4, "a"))
	fmt.Println(CoolRepeatStr(3, "hello "))
	fmt.Println(CoolRepeatStr(2, "abc"))

}
