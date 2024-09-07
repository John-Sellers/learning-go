package main

import "fmt"

func main() {

	var firstVar string = "First Variable"
	var secondVar uint = 25
	var firstName string
	var lastName string
	var emailAddress string
	var phoneNumber string


	fmt.Printf("Welcome to our Go Lang project here is my first variable: %v \n", firstVar)
	fmt.Printf("Let's get started with our second variable: %v \n", secondVar)
	fmt.Printf("The type of my first variable is: %T \n", firstVar)
	fmt.Printf("The type of my second variable is: %T \n", secondVar)

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address:")
	fmt.Scan(&emailAddress)

	fmt.Println("Please enter your phone number:")
	fmt.Scan(&phoneNumber)

	fmt.Printf(
		"Thank you %v %v. Would you prefer to receive an email at %v or a phone call at %v? \n",
		firstName,
		lastName,
		emailAddress, 
		phoneNumber,
	)



}