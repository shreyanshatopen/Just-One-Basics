package main

import (
	"fmt"
	"log"
)

func main() {
	// MODULE DEFER PANIC RECOVER
	fmt.Println("start")
	// func
	panicker()
	// after panicking
	fmt.Println("end")

}
func panicker() {

	fmt.Println("about to Panic")
	// just like throw func in c++  we keep func which is prone to errors in defer func
	defer func() {

		// code is that simple
		// if err is nill then no recovery
		if err := recover(); err != nil {
			// recover err msg here
			log.Println("Error :", err)
		}
	}()
	panic("sth bad happened")
	fmt.Println("done panicking")
}
