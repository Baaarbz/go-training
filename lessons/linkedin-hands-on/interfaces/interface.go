package main

import "fmt"

type humanoid interface {
	speak()
	walk()
}

type person struct {
	name string
}

func (p person) speak() {
	fmt.Printf("%s speaking...\n", p.name)
}

func (p person) walk() {
	fmt.Printf("%s walking...\n", p.name)
}

type dog struct {
}

func (d dog) walk() {
	fmt.Println("Dog speaking...")
}

func main() {
	// invoke with a person
	p := person{
		name: "Nelson",
	}
	doHumanThings(p)

	// can we invoke a dog?
	// doHumanThings(dog{}) NO we cant because it does not implement all the methods needed

	fmt.Println(p)
}

func doHumanThings(h humanoid) {
	h.speak()
	h.walk()
}
