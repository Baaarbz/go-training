package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func GuardInputs(userName string, email string, userTickets uint) (bool, bool) {
	isValidInput := len(userName) > 2 && strings.Contains(email, "@")
	isValidTickets := userTickets > 0

	return isValidInput, isValidTickets
}

func AskUserInputs() (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var userTickets uint
	// Ask user for their inputs
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstname)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastname)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("How many tickets do you want: ")
	fmt.Scan(&userTickets)

	return firstname, lastname, email, userTickets
}

func CreateMapOfBooking(firstname string, lastname string, email string, userTickets uint) map[string]string {
	mapUserData := make(map[string]string)
	mapUserData["name"] = firstname + " " + lastname
	mapUserData["email"] = email
	mapUserData["tickets"] = strconv.FormatUint(uint64(userTickets), 10)

	return mapUserData
}

func BookingMapOperations(bookingsMap []map[string]string) {
	var firstNames []string
	for _, booking := range bookingsMap {
		fields := strings.Fields(booking["name"])
		firstNames = append(firstNames, fields[0])
	}

	fmt.Printf("Emails that had made a booking:%v\n", firstNames)
}
