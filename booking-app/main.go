package main

import (
	"booking-app/helper"
	"fmt" // format
)

// Vamos a hacer un programa en CLI para reservar tickets de conferencias

// las que estan aqui son "package variables"
var conferenceName = "Go Conference"

const conferenceTickets uint8 = 50

var remainingTickets uint8 = conferenceTickets
var bookings []UserData

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint8
}

func main() {

	greetUser() // Printea el saludo de bienvenida

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketnumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		if isValidName && isValidEmail && isValidTicketnumber {
			bookTicket(firstName, lastName, email, userTickets)

		} else {
			if !isValidName {
				fmt.Println("Your name is invalid, try again")
			}

			if !isValidEmail {
				fmt.Println("Your email is invalid, try again")
			}

			fmt.Println("Your input is invalid, try again")
			continue
		}
		firstNames := getFirstNames()
		fmt.Printf("Bookings are: %v\n", firstNames)
	}
	fmt.Println("Thank you for booking. See you next time!")
}

func greetUser() {
	// Printea el saludo de bienvenida
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint8) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(
	firstName string,
	lastName string,
	email string,
	userTickets uint8,
) {

	// create a map for a user
	var userData UserData
	userData.firstName = firstName
	userData.lastName = lastName
	userData.email = email
	userData.numberOfTickets = userTickets
	/*
		o tambien
		var userData = UserData{
			firstName:       firstName,
			lastName:        lastName,
			email:           email,
			numberOfTickets: userTickets,
		}
	*/

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)
}
