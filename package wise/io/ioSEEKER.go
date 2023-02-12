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
	// to fix reader so that it can read twice we use seeker
	reader := io.Reader(fileRef)

	buffer, err := io.ReadAll(reader)

	fmt.Printf("%v is data, %v is error\n", string(buffer), err)

	// we will convert reader into seeker interface
	seeker := reader.(io.Seeker)
	// seekerEnd means from last it will read data from last
	seeker.Seek(-4, io.SeekEnd)
	buffer2, err2 := io.ReadAll(reader)
	// now operation is valid
	fmt.Printf("%v is data, %v is error\n", string(buffer2), err2)

	// seeker current means from current byte till mentioned in argument
	// currently seeker is at the end of the file so -4 positions means prev -
	seeker.Seek(-4, io.SeekCurrent)
	buffer3, err3 := io.ReadAll(reader)
	fmt.Printf("%v is data, %v is error\n", string(buffer3), err3)

}
