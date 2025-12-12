package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50 //uint aceita apenas numeros positivos

	fmt.Printf("Welcome to our %s booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here.")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //Scan() pega input do usuario e coloca no endereço de memoria de userName

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("\nEnter the number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)

	fmt.Printf("There are %v tickets left.\n", remainingTickets)

}
