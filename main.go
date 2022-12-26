package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avainlable\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var booking [50]string

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
	booking[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", booking)
	fmt.Printf("The first value: %v\n", booking[0])
	fmt.Printf("Array type: %T\n", booking)
	fmt.Printf("Array length: %v\n", len(booking))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, eamil)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
