package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// slice is static and cannot grow
	bs := make([]byte, 99999)

	// reads data into byte slice until it runs out of space
	resp.Body.Read(bs)

	// cast byte slice into string and print it
	fmt.Println(string(bs))
}
