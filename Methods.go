package main

import "fmt"

type user struct {
	name   string
	email  string
	status bool
	age    int
}

// methods are like func which waits for its execution
// only diff is methods contains a receiver argument
// it can be of struct n non struct type
func (u user) getMail() {
	fmt.Println(u.email)

}
func (u user) getName() {
	fmt.Println(u.name)

}
func (u user) getStatus() {
	fmt.Println(u.status)

}
func (u user) getAge() {
	fmt.Println(u.age)
}

func main() {
	Shreyansh := user{"Shreyansh", "shreyansh.tiwari@bankopen.co", true, 22}
	Shreyansh.getName()
	Shreyansh.getMail()
	Shreyansh.getStatus()
	Shreyansh.getAge()
	fmt.Println(Shreyansh)
}
