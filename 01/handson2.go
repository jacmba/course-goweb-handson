package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
}

func (p person) pSpeak() {
	fmt.Printf("Hello, I am %s %s\n", p.fName, p.lName)
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (s secretAgent) saSpeak() {
	fmt.Printf("My name is %s, %s %s\n", s.lName, s.fName, s.lName)
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

	fmt.Println(p.fName)
	p.pSpeak()

	fmt.Println(sa.fName)
	sa.saSpeak()
	sa.pSpeak()
}
