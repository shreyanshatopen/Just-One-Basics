package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// io.Writer

	fileRef, _ := os.Create("file.txt")
	defer fileRef.Close()

	writer := io.Writer(fileRef)

	n, err := writer.Write([]byte("Hello"))
	fmt.Println(n, err) // it will print size of []byte and err(if not then nil)

	// to add into files
	changedSize, err2 := io.WriteString(writer, "!")

	fmt.Println(changedSize, err2)
}
