package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// io.Writer

	fileRef, _ := os.Open("file.txt")

	defer fileRef.Close()

	reader := io.Reader(fileRef)
	minBuffer := make([]byte, 1) // this will read only one byte

	for {
		n, err := reader.Read(minBuffer)

		fmt.Printf("Current Byte is %v, no. byte read: %v, Error(if not then nill): %v\n", string(minBuffer), n, err)
		if err != nil {
			break
		}

	}

}
