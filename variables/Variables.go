package main

import "fmt"

func main() {
	// type 1 declaration
	var varname int = 10
	fmt.Printf("%v", varname)
	// type 2 declaration(senses datatype automatically)
	varname2 := 11
	fmt.Println(varname2)
	stringvar := "Shreyansh"
	fmt.Println(stringvar)

	// using float var(it will trim trailing zeroes ex 10.6 will be printed)

	var floatWala float32 = 10.60
	fmt.Println(floatWala)

}
