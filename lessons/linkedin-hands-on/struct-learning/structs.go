package main

import "fmt"

type author struct {
	first string
	last  string
}

func (a author) fullName() string {
	return a.first + " " + a.last
}

func (a *author) changeName(first, last string) {
	a.first = first
	a.last = last
}

func main() {
	a := author{
		first: "James",
		last:  "Baldwin",
	}
	fmt.Printf("%#v\n", a)
	fmt.Println(a.fullName())

	a.changeName("Eduardo", "Barbosa")
	fmt.Println(a.fullName())
}
