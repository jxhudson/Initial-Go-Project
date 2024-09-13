package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your naame: ")

	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!", name)

	fmt.Printf("\n Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 20 {
		fmt.Println("You can continue")
	} else {
		fmt.Println("Must be over 20 to play")
		return
	}

	score := 0

	fmt.Printf("What is the creators first name? ")
	var answer string
	fmt.Scan(&answer)

	if answer == "Justin" {
		fmt.Println("You are correct")
		score++
	} else if answer == "justin" {
		fmt.Println("You are correct")
		score++
	} else {
		fmt.Println("Try again, next time")
	}

	fmt.Printf("How many siblings does Justin have? ")
	var siblings uint
	fmt.Scan(&siblings)

	if siblings == 1 {
		fmt.Println("You are correct")
		score++
	} else {
		fmt.Println("Try again, next time")
	}

	fmt.Printf("You scored %v out of 2 questions.", score)
}
