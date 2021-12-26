package main

import (
	"fmt"
	"time"
)

const conferenceTickets int = 50

var conferneceName = "Go Conference"
var remainigTicketes uint = 50

// map 사용시
//var bookings = make([]map[string]string, 0)

// structure 사용시
var bookings = make([]UserDdata, 0)

//var bookings = []string{}

// map을 사용하는 것보다 한줄로 깔금하게 나오는거 같음

type UserDdata struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

//var wg = sync.WaitGroup{}

func main() {

	gretUsers()

	for {

		firtstName, lastName, email, UserTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firtstName, lastName, email, UserTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(UserTickets, firtstName, lastName, email)

			//wg.Add(1)
			// 메일을 보내려고 10초 기다리게 했는데 이 부분을 따로 넘기고 그다음을 바로 진행 을 위해 설정
			go sendTicket(int(UserTickets), firtstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookins are: %v\n", firstNames)

			if remainigTicketes == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				//	break
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
	//wg.Wait()
}

func gretUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferneceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable\n", conferenceTickets, remainigTicketes)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//map 사용시
		//firstNames = append(firstNames, booking["firstName"])

		//structure 사용시
		firstNames = append(firstNames, booking.firstName)
		// slice만 사용시
		//var names = strings.Fields(booking)
		//firstNames = append(firstNames, names[0])
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

	// create a map for a user
	var userData = UserDdata{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	/*
		var userData = make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	*/
	//bookings = append(bookings, firstName+" "+lastName)

	//map , slice 동일
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainigTicketes, conferneceName)
}
func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("############################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("############################")
	//wg.Done()
}
