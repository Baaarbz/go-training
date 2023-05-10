package main

import (
	"fmt"
	"go-training/booking-app/helper"
)

const ConferenceName = "Go Conference"
const ConferenceTickets = 50

var bookings []UserData

var bookingsMap []map[string]string

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	var remainingTickets uint = ConferenceTickets
	printVariableType()

	fmt.Printf("Welcome to %v booking application!\n", ConferenceName)
	fmt.Printf("We have total of %v and there are %v available.\n", ConferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		firstName, lastName, email, userTickets := helper.AskUserInputs()
		isValidInput, isValidTickets := helper.GuardInputs(firstName, email, userTickets)

		if isValidInput && isValidTickets {
			if userTickets > remainingTickets {
				fmt.Printf("We only have %v tickets remaining, You can not book %v tickets\n", remainingTickets, userTickets)
				continue
			}
			mapOperations(firstName, lastName, email, userTickets)
			structOperations(firstName, lastName, email, userTickets)
			remainingTickets = remainingTickets - userTickets

			fmt.Printf("User %v booked %v tickets.\nYou will receive a confirmation email to: %v\n", firstName+" "+lastName, userTickets, email)
			fmt.Printf("Remaining tickets: %v\n", remainingTickets)

			if remainingTickets == 0 {
				fmt.Println("Conference is booked out. Come back next year!")
				break
			}
		} else {
			if !isValidTickets {
				fmt.Println("You can not book less than 1 ticket")
			}
			if !isValidInput {
				fmt.Println("Check you inputs! Email or name it is wrong")
			}
		}
	}
	printArrayInfo()
}

func mapOperations(firstName string, lastName string, email string, userTickets uint) {
	bookingMap := helper.CreateMapOfBooking(firstName, lastName, email, userTickets)
	bookingsMap = append(bookingsMap, bookingMap)
	helper.BookingMapOperations(bookingsMap)
}

func structOperations(firstName string, lastName string, email string, userTickets uint) {
	bookingStruct := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, bookingStruct)
}

func printVariableType() {
	fmt.Printf("conferenceTickets is %T\n", ConferenceTickets)
	fmt.Printf("conferenceName is %T\n", ConferenceName)
}

func printArrayInfo() {
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("First array value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array lenght: %v\n", len(bookings))
}
