package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avainlable\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var eamil string
		var userTickets uint
		// Ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scanln(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scanln(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scanln(&eamil)

		fmt.Println("Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, eamil)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			firstNames = append(firstNames, strings.Split(booking, " ")[0])
		}

		fmt.Printf("Thefirst names of bookings are: %v\n", firstNames)
	}

}
