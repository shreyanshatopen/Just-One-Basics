package main

import (
	"html/template"
	"io"
	"log"
	"os"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	templatePath := "package wise/TEXT/Parse.html"
	fileRef, err := os.Open("package wise/TEXT/Parse.html")
	handleError(err)
	// parsing html file by passing its path stored in templatePath
	t, err2 := template.New("Parse.html").ParseFiles(templatePath)
	handleError(err2)
	log.Println(*t)
	// executing(printing html file)
	// err3 := t.Execute(os.Stdout, templatePath)
	// handleError(err3)
	// now instead of printing in terminal we will print it in separate log file

	// Now printing whole html string in a separate log file
	file2getString, _ := os.Create("ValidhtmlString.log")
	// now we will save printed log outputs in "logFile.log"
	log.SetOutput(file2getString)

	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime)
	reader := io.Reader(fileRef)
	buffer := make([]byte, 5000) // buffer is for reading bytes

	n, err2 := reader.Read(buffer)
	log.Println(n, err2)
	log.Println(string(buffer)) // this statement we are not able to print in log file
	// this needs to be parse

}
