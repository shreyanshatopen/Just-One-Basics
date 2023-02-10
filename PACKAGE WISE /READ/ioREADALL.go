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

	buffer, err := io.ReadAll(reader)
	if err == nil {
		fmt.Println("message got conv in bytes is:", string(buffer))
	} else {
		fmt.Println(err)
	}

	// if you will try to read twice then nothing gonna happen and string(buffer) will
	// return invalid address
	// as reader has already reached to end of func
	buffer2, err2 := io.ReadAll(reader)
	fmt.Printf("message got conv in bytes is:%v & err is: %v", string(buffer2), err2.Error())

}
