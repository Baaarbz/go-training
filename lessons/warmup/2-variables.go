package main

import (
	"fmt"
)

const age = 42
const name = "Fernando Alonso"

func main() {
	var name, age = getPersonalInfo()
	fmt.Printf("%v is %v years old\n", name, age)
	fmt.Printf("%v years old\n", getAge())
}

func getPersonalInfo() (string, uint8) {
	return name, age
}
func getAge() uint8 {
	return age
}
