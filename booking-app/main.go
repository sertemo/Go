package main

import (
	"fmt"
	"strings"
) // format

// Vamos a hacer un programa en CLI para reservar tickets de conferencias

func main() {
	var conferenceName = "Go Conference" // o conferenceName := "Go Conference"
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = conferenceTickets
	var bookings []string

	fmt.Printf("conferenceTickets is %T and remainingTickets is %T\n", conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking app!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint8
		// ask the user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
			firstName, lastName, userTickets, email)
		fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("Bookings are: %v\n", firstNames)

		var noTicketsRemaining bool = remainingTickets == 0

		if noTicketsRemaining {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}

	}

}
