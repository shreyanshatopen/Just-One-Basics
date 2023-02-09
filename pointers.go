package main

import "fmt"

type myDetails struct {
	name  string
	age   int
	phone int64
}

func main() {
	var num int = 10
	var ptr *int = &num

	// pointers are used to avoid creating copies of a var when we pass it to function
	// and thus original value is being passed

	fmt.Println(num, *ptr)
	fmt.Println(num, ptr)

	num = 20
	fmt.Println(num, *ptr)

	*ptr = 30
	fmt.Println(num, *(ptr))

	// pointer to an array
	arr := [...]int{1, 2, 3, 4}
	b := &arr[0]
	c := &arr[1]
	fmt.Printf("%v %v %v \n", arr, *b, *c)

	// pointer to struct

	var newPtr *myDetails   // pointer to struct creation
	newPtr = new(myDetails) // object creation

	newPtr.name = "Shreyansh" // initialising data members
	newPtr.age = 22
	newPtr.phone = 9907408787
	fmt.Println(*newPtr) // printing by dereferencing

}
