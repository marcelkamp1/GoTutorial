package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(encrypt("dich"))
}

func encrypt(s string) string {

	// generate slice of strings
	slice := []string{}

	// iterate through length of string
	// append single letters to slice
	for i := 0; i < len(s); i++ {
		slice = append(slice, string(s[i]))
	}

	// swap 2nd and last letter
	slice[1], slice[len(slice)-1] = slice[len(slice)-1], slice[1]

	// turn s[0] into []byte and initialize bytes.Buffer
	ascii := []byte{s[0]}
	buf := bytes.NewBuffer(ascii)

	// convert to int64
	asciiInt, err := binary.ReadUvarint(buf)
	if err != nil {
		return "Error"
	}

	// convert int64 into string
	asciiStr := strconv.FormatInt(int64(asciiInt), 10)

	// edit s[0] to ascii-value
	slice[0] = asciiStr

	// turn slice into string and return it
	return strings.Join([]string(slice), "")

}

/*
Step 1: convert first letter to ascii
Step 2: swap 2nd letter with letter at len(s)-1
Step 3: append both to a slice
Step 4: convert slice to string with strings.join
Step 5: return final string
*/
