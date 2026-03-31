package main

import "fmt"

func main() {
	if 0 == 7 % 2 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 0 == 8 % 4 {
		fmt.Println("8 is divisible by 4")
	}

	if 0 == 8 % 2 || 0 == 7 % 2 {
		fmt.Println("Either 8 or 7 are even")
	}

	if num := 9; 0 > num {
		fmt.Println(num, "is negative")
	} else if 10 > num {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
