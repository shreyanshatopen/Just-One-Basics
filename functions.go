package main

import "fmt"

// pointer as an argument
func print(greeting, name *string) {
	var hell *string = greeting
	// I am checking whether reinitialising name will work or not
	*name = "Shanus"
	fmt.Println("Name ptr: ", *name)
	fmt.Println(*hell)
}

// add two no.s
func adder(a int, b int) int {
	return a + b
}

// add values altogether without passing it in an array
func proAdder(values ...int) (int, string) {
	var add int = 0
	for _, val := range values {
		add += val
	}
	return add, "This is the message"
}
func main() {

	// function inside func is not allowed
	greeting := "Hello!"
	name := "Shreyansh"
	print(&greeting, &name)
	fmt.Println("Name is", name)

	ans := adder(1, 2)
	fmt.Println("twoSum : ", ans)

	added, myMessage := proAdder(1, 2, 3, 4, 5)
	fmt.Println("MultipleSum :", added)
	fmt.Println("Second return value : ", myMessage)

}
