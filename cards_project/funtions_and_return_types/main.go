package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
	fmt.Println(sum(2, 3))
	fmt.Println(isTrue())
	fmt.Println(radius(5))
}

func newCard() string {
	return "Five of Diamonds"
}

/*
	newCard -> defines a function called 'new card'
	string -> when executed, this function will return a value of type
			  'string'
*/

func sum(a int, b int) int {
	return a + b
}

func isTrue() bool {
	return true
}

func radius(a int) float64 {
	return float64(a) * 96.9
}