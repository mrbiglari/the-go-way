package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func stringFunctions() {
	p(s.Contains("test", "t"))
	p(s.Count("test", "t"))
	p(s.HasPrefix("test", "te"))
	p(s.HasSuffix("test", "st"))
	p(s.Index("test", "s"))
	p(s.Join([]string{"a", "b", "c"}, "-"))
	p(s.Repeat("abc", 5))
	p(s.Replace("book", "o", "0", -1))
	p(s.Replace("book", "o", "0", 1))
	p(s.Split("a,b,c,d,e,f", ","))
	p(s.ToLower("TEST"))
	p(s.ToUpper("test"))
}
