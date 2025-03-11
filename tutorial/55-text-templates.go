package main

import (
	"os"
	"text/template"
)

var languages = [...]string{"Go", "Rust", "C#", "Java", "C++"}

func textTemplates() {
	const format1 = "value is \"{{.}}\"\n"
	template1 := template.New("demo1")
	template1, error := template1.Parse(format1)
	if error != nil {
		panic(error)
	}

	template.Must(template1.Parse(format1)) // alternatively, we can use the template.Must function to panic in case Parse returns an error.

	template1.Execute(os.Stdout, "Bob")     // value is "Bob"
	template1.Execute(os.Stdout, 42)        // value is "42"
	template1.Execute(os.Stdout, languages) // value is "[Go Rust C# Java C++]"

	create := func(name, format string) *template.Template { // helper function to create new templates
		return template.Must(template.New(name).Parse(format))
	}

	const format2 = "fancy value is \"{{.}}\"\n"
	template2 := create("demo2", format2)
	template2.Execute(os.Stdout, "Sydney") // fancy value is "Sydney"

	const format3 = "fancy value is \"{{.Name}}\"\n"
	template3 := create("demo3", format3)
	template3.Execute(os.Stdout, struct{ Name string }{Name: "Amy"}) // fancy value is "Amy"
	template3.Execute(os.Stdout, map[string]string{"Name": "Buddy"}) // fancy value is "Buddy"

	const format4 = "{{if . -}} yes {{else -}} no {{end}}\n"
	template4 := create("demo4", format4)
	template4.Execute(os.Stdout, "not empty") // yes
	template4.Execute(os.Stdout, "")          // no

	const format5 = "range: {{range .}}{{.}} {{end}}\n"
	template5 := create("demo5", format5)
	template5.Execute(os.Stdout, languages) // range: Go Rust C# Java C++
}
