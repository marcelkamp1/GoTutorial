package main

// TODO ONGOING

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isValidIP("192.168.0.1"))
}

func isValidIP(ip string) bool {

	splits := strings.Split(ip, ".")
	newInts := []int{}

	if len(splits) == 4 {
		for _, elem := range splits {
			newInt, err := strconv.ParseInt(elem, 10, 64)
			if err != nil {
				fmt.Println("Error", err)
				return false
			}
			if newInt >= 0 && newInt <= 255 {
				newInts = append(newInts, int(newInt))
			}
		}

		if len(newInts) == 4 {
			fmt.Println(newInts)
			fmt.Println(len(newInts))
			return true
		}
	}
	return false
}
