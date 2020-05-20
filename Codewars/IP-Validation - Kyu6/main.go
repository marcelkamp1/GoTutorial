package main


// TODO ONGOING


import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(is_valid_ip("258.168.2.2"))
}

func is_valid_ip(ip string) bool {

	splits := strings.Split(ip, ".")
	newInts := []int{}

	if len(splits) == 4 {
		stringToInt(splits, newInts)
	}
	return true
}

func stringToInt(splits []string, newInts []int) {
	for _, elem := range splits {
		newInt, err := strconv.ParseInt(elem, 10, 64)
		if err != nil {
			fmt.Println("Error", err)
		}
		newInts = append(newInts, int(newInt))
		if len(newInts) == 4 {
			for _, elem := range newInts {
				if elem >= 0 && elem <= 255{
				}
			}
		}
	}
}