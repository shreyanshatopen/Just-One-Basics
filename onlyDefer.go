package main

import (
	"fmt"
)

func main() {
	num := 10
	defer fmt.Println(num + 20)
	fmt.Println(num)
	defer fmt.Println(num + 10)
	fmt.Println(num + 30)
	// non defer will be printed as in order top to bottom
	// defer order is like stack it follows LIFO order
	// last defer will be invoked first
}
