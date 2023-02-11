//Package log implements simple struct methods and functions to log program

package main

import (
	"log"
	"os"
)

// Package log is safe from concurrent goroutines
// it can attach info directly like date, time,filePath

func main() {
	file, _ := os.Create("logFile.log")
	// now we will save printed log outputs in "logFile.log"
	log.SetOutput(file)
	log.Println("Printed Using Log package")
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("Printed Using Log package")
	// log.Panic("Started Panicking...")
	// // immediately shuts program after panicking

	// log.Fatal()

}
