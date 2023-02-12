package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://open.money/"

func main() {
	fmt.Printf("URLs resopnse ")

	//http request, get method , it requires only one parameter
	response, err := http.Get(url)

	handleNilErr(err)
	fmt.Printf("OPEN's response: ", response)

	//closing the connection
	defer response.Body.Close()
	// reading the data from response body
	databyte, err := ioutil.ReadAll(response.Body)

	handleNilErr(err)
	// conv databyte to string
	content := string(databyte)
	fmt.Println("Databyte says: ", content)

}
func handleNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
