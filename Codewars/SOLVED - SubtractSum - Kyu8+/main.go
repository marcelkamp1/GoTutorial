package main

import (
	"fmt"
	"strconv"
)

// Kata-Solution is -> return "apple"
// subtraction always ends on a key with the value apple

func main() {
	fruits := map[int] string{
		1: "kiwi",
		2: "pear",
		3: "kiwi",
		4: "banana",
		5: "melon",
		6: "banana",
		7: "melon",
		8: "pineapple",
		9: "apple",
		10: "pineapple",
	}
	n := subtractSum(77)

	if f, key := fruits[n]; key {
		fmt.Println(f)
	}
}

func subtractSum(n int) int {
	for n > 10 {
		result := 0
		numbers := []int{}
		target := strconv.Itoa(n)
		first := target[:1]
		second := target[1:2]
		third := target[2:]

		firstInt, _ := strconv.Atoi(first)
		secondInt, _ := strconv.Atoi(second)
		thirdInt, _ := strconv.Atoi(third)

		numbers = append(numbers, firstInt, secondInt, thirdInt)
		for _, elem := range numbers {
			result = result + elem
		}
		n = n - result
	}
	return n
}
