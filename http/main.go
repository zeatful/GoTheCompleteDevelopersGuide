package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	// one line replacement for code below
	//io.Copy(os.Stdout, resp.Body)

	/*
		// slice is static and cannot grow
		bs := make([]byte, 99999)
		// reads data into byte slice until it runs out of space
		resp.Body.Read(bs)
		// cast byte slice into string and print it
		fmt.Println(string(bs))
	*/
}

func (logWriter) Write(bs []byte) (int, error) {
	return 1, nil
}
