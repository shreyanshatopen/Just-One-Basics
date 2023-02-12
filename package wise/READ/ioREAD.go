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
	buffer := make([]byte, 10) // buffer is for reading bytes

	n, err := reader.Read(buffer) // n is the number of bytes read
	fmt.Println(n, err)
	fmt.Println(buffer)         // info in byte terms
	fmt.Println(string(buffer)) // byte conv in string

}
