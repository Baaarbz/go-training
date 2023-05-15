package main

import "fmt"

// create a numeric interface with a type set
type numeric interface {
	~int | ~float64
	grow()
}

// update sumNumerics function to use a numeric interface with a type set
func sumNumerics[T numeric](a, b T) T {
	return a + b
}

type specialInteger int

func (si specialInteger) grow() {}

func main() {
	one := specialInteger(1)
	two := specialInteger(2)
	fmt.Println(sumNumerics(one, two))
}
