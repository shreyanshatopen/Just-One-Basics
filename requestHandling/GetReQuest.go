package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to get method")
	performgetRequuest()
}
func performgetRequuest() {
	const myurl = "https://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		fmt.Println("Error Occured")
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)

}
