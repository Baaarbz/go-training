package main

import (
	"fmt"
	"math/rand"
)

type person struct {
	name string
	age  uint
}

func (p person) printNameAndAge() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func (p *person) randomizeAge() {
	p.age = uint(rand.Uint32())
}

func main() {
	p := initPersonStruct()
	p.printNameAndAge()
	p.randomizeAge()
	p.printNameAndAge()
	fmt.Printf("%v\b %T\n", p, p)
}

func initPersonStruct() person {
	return person{
		name: "Juan Ibars",
		age:  59,
	}
}
