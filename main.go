package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickers is %T, conferenceName is %T\n", conferenceTicket, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avainlable\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
