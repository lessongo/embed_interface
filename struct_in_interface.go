package main

import "fmt"

var (
	mb = fmt.Println
	mf = fmt.Printf
)

type animal interface {
	breathe()
	walk()
}

type dog struct {
	age int
}

func (d dog) breathe() {
	fmt.Println("Dog breathes")
}

func (d dog) walk() {
	fmt.Println("Dog walk")
}

type pet1 struct {
	animal
	name string
}

type pet2 struct {
	animal
	name string
}

func Execute() {
	d := dog{age: 5}
	p1 := pet1{name: "Milo", animal: d}

	mb(p1.name)
	p1.walk()
	p1.breathe()

	p2 := pet2{name: "Oscar", animal: d}
	mb(p2.name)
	p2.breathe()
	p2.walk()
	p1.breathe()
	p1.walk()
}
