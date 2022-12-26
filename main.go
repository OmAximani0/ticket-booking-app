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

	greetUsers(conferenceName, conferenceTicket, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, remainingTickets, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(conferenceName, remainingTickets, userTickets, firstName, lastName, email, bookings)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name and last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers(conferenceName string, conferenceTicket int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, strings.Split(booking, " ")[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, remainingTickets uint, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// Ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(conferenceName string, remainingTickets uint, userTickets uint, firstName string, lastName string, email string, bookings []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
