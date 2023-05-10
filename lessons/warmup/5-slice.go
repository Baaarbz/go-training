package main

import "fmt"

func main() {
	weekdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	weekdays = append(weekdays, "Saturday")
	weekdays = append(weekdays, "Sunday")

	for index, weekday := range weekdays {
		fmt.Println(index)
		if index > 4 {
			fmt.Println("It's weekend!")
		} else {
			fmt.Printf("Current day is %s\n", weekday)
		}
	}
}
