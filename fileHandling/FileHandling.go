package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// to work with files like .txt
// also there's some special libraries for using .pdf .csv files we're not covering it now

func main() {
	fmt.Println("Lets begin with file handling")
	content := "This file is created using FileHandling.go operation"
	// above will be some content which we want to be in the file
	// os.Create will create file
	// err var is used to handle err
	// file var is we using to operating on newly created file
	file, err := os.Create("./createdfile.txt")

	checkNilErr(err)
	// write string will create file and add content in that file
	length, err := io.WriteString(file, content)
	// length var is for operating on newly created file

	checkNilErr(err)

	fmt.Println("Length is:", length)
	readFile("./createdfil.txt")
	defer file.Close()

}

// now reading file which is separate method
func readFile(filename string) {
	// we use ioutil.ReadFile to read file
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("data inside the file is \n", string(databyte))
	// you will notice
	//[84 104 105 115 32 102 105 108 101 32 105 115 32 99 114 101 97 116 101 100 32 117 115 105 110 103 32 70 105 108 101 72 97 110 100 108 105 110 103 46 103 111 32 111 112 101 114 97 116 105 111 110]
	// this kind of data format printed in
	// now how to handle this use string(readfile var name)
	//
}

// instead of writing 3 statements of checking nil err every where we can
// craete a sep func for them
func checkNilErr(err error) {
	if err != nil {
		fmt.Println("There's some error")
		panic(err)

	}
}
