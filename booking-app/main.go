package main

import (
// 참조 하는 폴더까지 다 적어야 한다.
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferneceName = "Go Conference"
var remainigTicketes uint = 50
var bookings = []string{}

func main() {

	gretUsers()

	for {

		firtstName, lastName, email, UserTickets := getUserInput()
		// 마찬가지로 helper.ValidateUserInput << 여기도 참조하는 helper.V 로 변경 fmt도 마찬가지지만 참조 하는 것은 fmt.P 이런식으로 대문자가 되어야함
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firtstName, lastName, email, UserTickets, remainigTicketes)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(UserTickets, firtstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookins are: %v\n", firstNames)

			if remainigTicketes == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}

func gretUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferneceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable\n", conferenceTickets, remainigTicketes)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter you first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter you last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter you email address:")
	fmt.Scan(&email)

	fmt.Println("Enter you number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainigTicketes = remainigTicketes - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainigTicketes, conferneceName)
}
