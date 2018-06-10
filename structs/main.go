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
	jimPointer := &jim // convert jim to a pointer

	// call updateName on this address
	jimPointer.updateName("jimmy")
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
	//* infront of a type means it is a pointer to an object of that type
	func( pointerToSomething *something) method() {
		// * in front of a variable means the value of that object
		*pointerToSomething
	}
*/

// pointer that points at a person
func (p *person) updateName(newFirstName string) {
	// converts person to a value
	(*p).firstName = newFirstName
}
