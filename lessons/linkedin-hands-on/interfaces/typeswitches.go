package main

import "fmt"

func whatTypeAmI(val interface{}) string {
	switch val.(type) {
	case int:
		return fmt.Sprintf("%v is an int", val)
	case string:
		return fmt.Sprintf("%v is a string", val)
	default:
		return fmt.Sprintf("%v is not supported type", val)
	}

}
func main() {
	fmt.Println(whatTypeAmI(1))
	fmt.Println(whatTypeAmI("hello"))
	fmt.Println(whatTypeAmI(true))
}
