package main

import "fmt"

func main() {

	var firstVar string = "First Variable"
	var secondVar uint = 25
	var firstName string
	var lastName string
	var emailAddress string

	someFunc(firstVar)

	fmt.Printf("Welcome to our Go Lang project here is my first variable: %v \n", firstVar)
	fmt.Printf("Let's get started with our second variable: %v \n", secondVar)
	fmt.Printf("The type of my first variable is: %T \n", firstVar)
	fmt.Printf("The type of my second variable is: %T \n", secondVar)

	
	for {
		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email address:")
		fmt.Scan(&emailAddress)

		fmt.Printf(
			"Thank you %v %v. Would you prefer to receive an email at %v? \n",
			firstName,
			lastName,
			emailAddress,
		)

		var arrayOfNames []string

		// arrayOfNames[0] = firstName + " " + lastName
		arrayOfNames = append(arrayOfNames, firstName + " " + lastName)

		fmt.Printf("The whole slice: %v \n", arrayOfNames)
		fmt.Printf("The first element of the slice: %v \n", arrayOfNames[0])
		fmt.Printf("The length of the slice: %v \n", len(arrayOfNames))
		fmt.Printf("The type of the slice is: %T \n", arrayOfNames)
	}
}

func someFunc(someVar string) {
	fmt.Printf("This function just did something! Here is the first var value: %v \n", someVar)
}