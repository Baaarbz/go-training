package main

import "fmt"

func whatAmI(val interface{}) string {
	return fmt.Sprintf("%v, %T", val, val)
}

func main() {
	// Invoke whatIAm() with an int
	fmt.Println(whatAmI(1))
}
