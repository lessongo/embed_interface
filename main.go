package main

import "fmt"

var (
	mb = fmt.Println
)

type animal interface {
	breathe()
	walk()
}

type human interface {
	animal
	speak()
}

type employee struct {
	name string
}

func (e employee) breathe() {
	mb("Employee breathes")
}

func (e employee) walk() {
	mb("Employee walk")
}

func (e employee) speak() {
	mb("Employee speak")
}

func main() {
	var h human
	h = employee{name: "Recai"}
	h.breathe()
	h.walk()
	h.speak()
}
