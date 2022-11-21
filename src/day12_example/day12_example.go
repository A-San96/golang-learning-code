package main

import "fmt"

func main() {
	challenge := "#90DaysOfDevOps"
	var remainingDays uint = 90
	const daystotal int = 90

	fmt.Printf("Welcome to %v\n", challenge)
	fmt.Printf("This is a %v challenge\n", daystotal)

	var TwitterName string
	var DaysComplete uint

	TwitterName = "@asanjaata"
	DaysComplete = 12

	fmt.Println("Enter your Twitter Handle: ")
	fmt.Scanln(&TwitterName)

	fmt.Println("How many days have you completed: ")
	fmt.Scanln(&DaysComplete)

	remainingDays = remainingDays - DaysComplete

	fmt.Printf("%v has completed %v days of the challenge\n", TwitterName, DaysComplete)
	fmt.Printf("You have %v days remaining for the %v challenge", remainingDays, challenge)
	fmt.Printf("Good luck !")
}
