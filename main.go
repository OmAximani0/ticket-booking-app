package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTicket, "tickets and", remainingTickets, "are still avainlable")
	fmt.Println("Get your tickets here to attend")
}
