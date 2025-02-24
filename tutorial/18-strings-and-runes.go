package main

import (
	"fmt"
	"unicode/utf8"
)

func stringsAndRunes() {

	const text = "สวัสดี"

	fmt.Println("Text value is:", text)
	fmt.Println("Length of text is:", len(text))

	for i := 0; i < len(text); i++ {
		fmt.Printf("%x ", text[i])
	}
	fmt.Println("\nRune count:", utf8.RuneCountInString(text))

	for index, _rune := range text {
		fmt.Printf("%#U starts at %d\n", _rune, index)
	}

	for index, increment := 0, 0; index < len(text); index += increment {
		_rune, width := utf8.DecodeRuneInString(text[index:])
		fmt.Printf("%#U starts at %d\n", _rune, index)
		increment = width

		examineRune(_rune)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
