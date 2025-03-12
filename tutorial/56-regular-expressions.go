package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func regularExpressions() {
	regex := "p([a-z]+)ch"
	match1, _ := regexp.MatchString(regex, "peach")
	fmt.Println(match1) // true

	_regexp, _ := regexp.Compile(regex)
	match2 := _regexp.MatchString("peach")
	fmt.Println(match2) // true

	firstMatch := _regexp.FindString("punch peach")
	fmt.Println(firstMatch)                                                  // punch
	fmt.Println("index: ", _regexp.FindStringIndex("peach punch"))           // index:  [0 5]
	fmt.Println(_regexp.FindStringSubmatch("peach punch"))                   // [peach ea]
	fmt.Println(_regexp.FindStringSubmatchIndex("peach punch"))              // [0 5 1 3]
	fmt.Println(_regexp.FindAllString("peach punch pinch", -1))              // [peach punch pinch]
	fmt.Println(_regexp.FindAllString("peach punch pinch", 2))               // [peach punch]
	fmt.Println(_regexp.FindAllStringSubmatchIndex("peach punch pinch", -1)) // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	fmt.Println(_regexp.Match([]byte("peach"))) // true

	_regexp = regexp.MustCompile(regex) // when creating global variables with regular expressions you can use the MustCompile variation of Compile. MustCompile panics instead of returning an error, which makes it safer to use for global variables.
	fmt.Println("regexp: ", _regexp)    // regexp:  p([a-z]+)ch

	replacedString := _regexp.ReplaceAllString("peach pinch", "bob")
	fmt.Println(replacedString) // bob bob

	in := []byte("a peach")
	out := _regexp.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) // a PEACH
}
