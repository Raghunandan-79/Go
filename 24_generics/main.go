package main

import "fmt"

func add[T int | float64](a T, b T) T {
	return a + b
}

func printSlice[T any](slice []T) {
	for _, v := range slice {
		fmt.Println(v)
	}
}

type Box[T any] struct {
	value T
}

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(add(2.5, 3.1))	

	printSlice([]int{1, 2, 3})
	printSlice([]string{"go", "is", "fast"})

	intBox := Box[int]{value: 10}
	strBox := Box[string]{value: "Go"}
	fmt.Println(intBox.value)
	fmt.Println(strBox.value)
}