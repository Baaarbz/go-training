package main

import "fmt"

// create a number interface with a type set
type number interface {
	~int | ~float64
	grow()
}

// update sumNumber function to use a number interface with a type set
func sumNumber[T number](a, b T) T {
	return a + b
}

type specialNumber int

func (s specialNumber) grow() {}

// equal returns true if a and b are equal.
func equal[T comparable](a, b T) bool {
	return a == b
}

func main() {
	// invoke equal with comparable types
	fmt.Println("equal(1, 1):", equal(1, 1))
	fmt.Println("equal(\"one\", \"two\"):", equal("one", "two"))

	// invoke equal with a custom type
	type c struct{ f string }
	fmt.Println("equal(c{f: \"a\"}, c{f: \"a\"}):", equal(c{f: "a"}, c{f: "a"}))
	fmt.Println("equal(c{f: \"a\"}, c{f: \"b\"}):", equal(c{f: "a"}, c{f: "b"}))
}
