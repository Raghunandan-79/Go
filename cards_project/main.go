package main

import "fmt"

func main() {
	// Creating a vairable
	var card string = "Ace of Spades"
	fmt.Println(card)

	/*
	Breakdown of above code:
		var -> We're about to create a new variable
		card -> The name of the variable will be 'card'
		string -> Only a "string" will be assigned to this variable
		"Ace of Spades" -> Assign the value "Ace of Spades" 
						   to this variable
	*/

	/*
	Basic datatypes in Go:
	
		- bool: true or false
		- string: "Hi!", "Hows it going?"
		- int: 0, -10000, 99999
		- float64: 10.000001, 0.00009, -100.003
	*/

	// we only us this := when we are first time declaring variable
	card1 := "Ace of Spades"
	fmt.Println(card1)

	card1 = "Five of Diamonds"
	fmt.Println(card1)
}