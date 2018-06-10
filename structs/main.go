package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// embedded struct, don't have to give a name, can just use the type
	//contact contactInfo
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000, // with nested structs, every line MUST have a comma
		},
	}

	//fmt.Printf("%+v", jim)

	// give me the memory address of jim
	//jimPointer := &jim // convert jim to a pointer

	// call updateName on this pointer
	//jimPointer.updateName("jimmy")
	jim.updateName("jimmy")
	jim.print()

	//john := person{firstName: "John", lastName: "Doe"}
	//fmt.Println(john)
	// prints each field name and value
	//fmt.Printf("%+v", john)
}

// receivers with structs
func (p person) print() {
	fmt.Printf("%+v", p)
}

/*
	// * here is a type description - takes a pointer to something
	func( pointerToSomething *something) method() {
		// * here is an operator - converting pointer to a something value
		*pointerToSomething
	}
*/

// pointer that points at a person
func (p *person) updateName(newFirstName string) {
	// converts person to a value
	(*p).firstName = newFirstName

	// Question, how is the below different? Didn't need to
	// use & above in main to get pointer receiver either?
	//p.firstName = newFirstName
}

/* pointer and reference conversions:

Turn |address| into |value|   with |*address|
Turn |value|   into |address| with |&value|
*/

// ********* GO gotcha Slices do not pass by value ********

/*
	----------------------------------
	 Passed By Value (need pointers)
	----------------------------------
	int
	float
	string
	bool
	struct

	------------------------------------------
	 Passed By Reference (don't need pointers)
	------------------------------------------
	slice
	map
	channel
	pointer (duh?)
	function

*/
