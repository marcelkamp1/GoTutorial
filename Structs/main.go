package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// Option 1 for structs:
	alex := person{"Alex", "Anderson", contactInfo{"alex.a@google.com", 53773}}

	// Option 2 structs:
	john := person{firstName: "John", lastName: "Doe", contactInfo: contactInfo{email: "john.doe@google.com", zipCode: 53773}}

	// Option 3 structs:
	var bob person
	bob.firstName = "Bob"
	bob.lastName = "der Baumeister"
	bob.email = "bob@baumeister.com"
	bob.zipCode = 53773

	// Declare a Pointer
	alex.updateName("Alexander")
	alex.print()


	john.print()
	bob.print()
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

// Pointer Method
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
