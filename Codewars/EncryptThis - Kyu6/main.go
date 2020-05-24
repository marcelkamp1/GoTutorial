package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(encrypt("Hello"))
}

func encrypt(s string) string {

	ascii, err := strconv.Atoi(s)
	if err == nil {
		return "Error"
	}
	fmt.Println(ascii)
	return string(ascii)
}

/*
Step 1: convert first letter to ascii
Step 2: swap 2nd letter with letter at len(s)-1
Step 3: append both to a slice
Step 4: convert slice to string with strings.join
Step 5: return final string
 */