package main

import "fmt"

type runner interface{ Run() string }
type jumper interface{ Jump() string }
type athlete struct{ name string }

func (a *athlete) Run() string     { return "runs" }
func (a *athlete) GetName() string { return a.name }
func (a athlete) Jump() string     { return "jumps" }

func typeAssertion() {
	print := fmt.Println

	var i interface{} = 42
	q := i.(int)
	print(q)

	var j interface{} = "42"
	k, ok := j.(int)
	if ok {
		print(k)
	} else {
		print("not an integer")
	}

	var _athlete interface{} = athlete{"matilda"}
	print(_athlete.(athlete).name, "rests")

	var _runner runner = &athlete{"bob"}
	print(_runner.(*athlete).GetName(), _runner.Run())

	var _jumper jumper = athlete{"sydney"}
	print(_jumper.(athlete).name, _jumper.Jump())
}
