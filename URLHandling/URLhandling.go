package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghjb456ghbj"

func main() {
	// handling URL means taking data from URL especially parameter
	fmt.Println("Handling URLs")

	// lets parse url (parsing means extracting info from URL)
	answer, _ := url.Parse(myurl)

	fmt.Println("URL is: ", myurl)
	fmt.Println("Port is: ", answer.Port())
	fmt.Println("Scheme is ", answer.Host)
	fmt.Println("Path is: ", answer.Path)
	fmt.Println("Scheme is ", answer.Scheme)

	// now you will print parsed details ex: components of URL

}
