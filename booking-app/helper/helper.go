// helper functions of the main application

package helper

import "strings"

func ValidateUserInput( // Para decirle a go que es exportable, poner mayusculas a la primera letra
	firstName string,
	lastName string,
	email string,
	userTickets uint8,
	remainingTickets uint8,
) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 3 && len(lastName) >= 3
	var isValidEmail bool = strings.Contains(email, "@")
	isValidTicketnumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketnumber
}
