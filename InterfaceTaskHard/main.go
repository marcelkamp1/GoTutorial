package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Cannot read the file")
		os.Exit(1)
	}
	fmt.Println(string(file))
}


