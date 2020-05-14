package main

import "fmt"

func main() {

	// var colors map[string]string

	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	delete(colors, "red")

	/*
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
	}*/



	fmt.Println(colors)
}
