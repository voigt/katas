package main

import (
	"fmt"
	"strings"
)

var MORSE_CODE = map[string]string{
	".-":     "a",
	"-...":   "b",
	"-.-.":   "c",
	"-..":    "d",
	".":      "e",
	"..-.":   "f",
	"--.":    "g",
	"....":   "h",
	"..":     "i",
	".---":   "j",
	"-.-":    "k",
	".-..":   "l",
	"--":     "m",
	"-.":     "n",
	"---":    "o",
	".--.":   "p",
	"--.-":   "q",
	".-.":    "r",
	"...":    "s",
	"-":      "t",
	"..-":    "u",
	"...-":   "v",
	".--":    "w",
	"-..-":   "x",
	"-.--":   "y",
	"--..":   "z",
	".-.-":   "ä",
	"---.":   "ö",
	"..--":   "ü",
	"----":   "Ch",
	"-----":  "0",
	".----":  "1",
	"..---":  "2",
	"...--":  "3",
	"....-":  "4",
	".....":  "5",
	"-....":  "6",
	"--...":  "7",
	"---..":  "8",
	"----.":  "9",
	".-.-.-": ".",
	"--..--": ",",
	"..--..": "?",
	"-.-.--": "!",
	"---...": ":",
	".-..-.": "\"",
	".----.": "'",
	"-...-":  "=",
}

// my attempt
func DecodeMorse(morseCode string) (result string) {

	if len(morseCode) <= 0 {
		return
	}

	if strings.Compare(morseCode, " ") == 0 {
		return
	}

	length := len(strings.Split(morseCode, " "))

	for i, char := range strings.Split(morseCode, " ") {
		if char == "" {
			if len(result)-1 >= 0 &&
				result[len(result)-1:] != " " &&
				i+2 < length { // quite ugly
				result += " "
			}
		} else {
			result += MORSE_CODE[char]
		}
	}

	return
}

// better attempt
// I didn't know of strings.TrimSpace. Definitely makes things easier!
func ElegantDecodeMorse(morseCode string) (decoded string) {
	for _, word := range strings.Split(strings.TrimSpace(morseCode), "   ") {
		for _, char := range strings.Split(word, " ") {
			decoded += MORSE_CODE[char]
		}
		decoded += " "
	}
	return strings.TrimSpace(decoded)
}

func main() {
	fmt.Println(DecodeMorse(""))
	fmt.Println(DecodeMorse(".... . -.--   .--- ..- -.. ."))
	fmt.Println(DecodeMorse("...."))
	fmt.Println(DecodeMorse("...---..."))
	fmt.Println(DecodeMorse(" "))
	fmt.Printf("'%s'\n", DecodeMorse(" . "))
	fmt.Printf("'%s'\n", DecodeMorse("   .   . "))
	fmt.Printf("'%s'\n", DecodeMorse("      ...---... -.-.--   - .... .   --.- ..- .. -.-. -.-   -... .-. --- .-- -.   ..-. --- -..-   .--- ..- -- .--. ...   --- ...- . .-.   - .... .   .-.. .- --.. -.--   -.. --- --. .-.-.-  "))

	fmt.Println(ElegantDecodeMorse(""))
	fmt.Println(ElegantDecodeMorse(".... . -.--   .--- ..- -.. ."))
	fmt.Println(ElegantDecodeMorse("...."))
	fmt.Println(ElegantDecodeMorse("...---..."))
	fmt.Println(ElegantDecodeMorse(" "))
	fmt.Printf("'%s'\n", ElegantDecodeMorse(" . "))
	fmt.Printf("'%s'\n", ElegantDecodeMorse("   .   . "))
	fmt.Printf("'%s'\n", ElegantDecodeMorse("      ...---... -.-.--   - .... .   --.- ..- .. -.-. -.-   -... .-. --- .-- -.   ..-. --- -..-   .--- ..- -- .--. ...   --- ...- . .-.   - .... .   .-.. .- --.. -.--   -.. --- --. .-.-.-  "))

}
