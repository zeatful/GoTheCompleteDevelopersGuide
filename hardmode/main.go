package main

import (
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	file, _ := os.Open(fileName)
	io.Copy(os.Stdout, file)
}
