package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
}

func (p person) speak() {
	fmt.Printf("Hello, I am %s %s\n", p.fName, p.lName)
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (s secretAgent) speak() {
	fmt.Printf("My name is %s, %s %s\n", s.lName, s.fName, s.lName)
}

type human interface {
	speak()
}

func say(h human) {
	h.speak()
}

func main() {
	p := person{
		"John",
		"Doe",
	}

	sa := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	say(p)
	say(sa)
	say(sa.person)
}
