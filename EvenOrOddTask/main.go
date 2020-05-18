package main

import "fmt"

func main() {
	ints := []int{}
	dynamicRange := 100

	for i := 0; i <= dynamicRange; i++ {
		ints = append(ints, i)
	}

	for _, n := range ints {
		if ints[n]%2 != 0 {
			fmt.Printf("%v is odd \n", ints[n])
		}

		if ints[n]%2 == 0 {
			fmt.Printf("%v is even \n", ints[n])
		}
	}
}
