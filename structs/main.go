package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	john := person{firstName: "John", lastName: "Doe"}
	fmt.Println(john)
	// prints each field name and value
	fmt.Printf("%+v", john)
}
