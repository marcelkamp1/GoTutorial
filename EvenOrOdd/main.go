package main

import "fmt"

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range ints {
		if ints[i]%2 != 0 {
			fmt.Printf("%v is odd \n", ints[i])
		}

		if ints[i]%2 == 0 {
			fmt.Printf("%v is even \n", ints[i])
		}
	}
}
