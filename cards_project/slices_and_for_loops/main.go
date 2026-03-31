package main

import "fmt"

/*
	Array: Fixed length list of things of same type
	Slice: An array that can grow or shrink but of same type
*/

func main() {
	// arrays
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	for i, val := range arr {
		fmt.Println(i, val)
	}

	// slices
	cards := []string{"Ace of Diamonds", newCard()}
	fmt.Println(cards)
	cards = append(cards, "King of Hearts")
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
