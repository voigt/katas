package main

import "testing"

var tables = []struct {
	encodedStr string
	expectStr  string
}{
	{".-", "a"},
	{".", "e"},
	{"..", "i"},
	{". .", "ee"},
	{".   .", "e e"},
	{"...---...", ""},
	{"... --- ...", "sos"},
	{"...   ---   ...", "s o s"},

	{" . ", "e"},
	{"   .   . ", "e e"},
	{"      ... --- ... -.-.--   - .... .   --.- ..- .. -.-. -.-   -... .-. --- .-- -.   ..-. --- -..-   .--- ..- -- .--. ...   --- ...- . .-.   - .... .   .-.. .- --.. -.--   -.. --- --. .-.-.-  ", "sos! the quick brown fox jumps over the lazy dog."},
}

func TestDecodeMorse(t *testing.T) {

	for _, table := range tables {
		decodedStr := DecodeMorse(table.encodedStr)
		if decodedStr != table.expectStr {
			t.Errorf("Decoding of '%s' was incorrect got: %s, want: %s.", table.encodedStr, decodedStr, table.expectStr)
		}
	}
}

func TestElegantDecodeMorse(t *testing.T) {
	for _, table := range tables {
		decodedStr := ElegantDecodeMorse(table.encodedStr)
		if decodedStr != table.expectStr {
			t.Errorf("Decoding of '%s' was incorrect got: %s, want: %s.", table.encodedStr, decodedStr, table.expectStr)
		}
	}
}
