package main

import "fmt"

func main() {
	// Print the address of the "x" variable (in other words the pointer of x).
	x := 40
	fmt.Printf("%v\n", &x)
	// Store the address of x to a variable named y. Use y to print the value 40 (basically dereference y).
	y := &x
	fmt.Printf("%v\n", *y)
}
