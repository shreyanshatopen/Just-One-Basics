package main

import (
	"fmt"
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
	fileRef, err := os.Create("Parse.html")

	handleError(err)
	defer fileRef.Close()
	writer := io.Writer(fileRef)

	towrite, err := writer.Write([]byte(`
	<!DOCTYPE html>
	<html lang="en">
	
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	
		<style>
			{
				font-family: 'Roboto', sans-serif;
			}
	
			.page {
				page-break-inside: auto;
			}
	
			.container {
				width: 75%;
				margin: 50px auto;
			}
	
			table {
				border-collapse: collapse;
				width: 100%;
			}
	
			tr {
				border-bottom: solid thin;
			}
			 td {
				padding: 5px 5px;
				font-size: 8px;
			}
		</style>
	
	</head>
	
	<body>
		<div class="page">
			<div class="container">
								
							<h4 style="text-align: left;">PSRS FOLYILMS</h4>
				<h5 style="text-align: center;">Receivables Report</h5>
				<table>
					<tr style="background-color: lightgrey;">
						<td style="text-align: left;"><b>Contact Name</b></td>
						<td style="text-align: left;"><b>Contact Type</b></td>
						<td style="text-align: left;"><b>Email ID</b></td>
						<td style="text-align: left;"><b>Phone No</b></td>
						<td style="text-align: right;"><b>Total Billed</b></td>
						<td style="text-align: right;"><b>Total Outstanding</b></td>
						<td style="text-align: right;"><b>Dso</b></td>
					</tr>
										 <tr>
							<td style="text-align: left;">ANVESH AXISD</td>
							<td style="text-align: left;">Both</td>
							<td style="text-align: left;">shubham.chawla@bankopen.co</td>
							<td style="text-align: left;">8078630803</td>
							<td style="text-align: right;">10539.00</td>
							<td style="text-align: right;">10539.00</td>
							<td style="text-align: right;">0.00</td>
						</tr>
										<tr>
							<td style="text-align: left;">Anvesh NIMPD</td>
							<td style="text-align: left;">Both</td>
							<td style="text-align: left;">anilkumar.reddy@bankopen.co</td>
							<td style="text-align: left;">1234567891</td>
							<td style="text-align: right;">549847.00</td>
							<td style="text-align: right;">549847.00</td>
							<td style="text-align: right;">0.00</td>
						</tr>
										<tr>
							<td style="text-align: left;">Anil</td>
							<td style="text-align: left;">Customer</td>
							<td style="text-align: left;">anil.reddy@gmail.com</td>
							<td style="text-align: left;">8995961144</td>
							<td style="text-align: right;">82095.02</td>
							<td style="text-align: right;">82086.02</td>
							<td style="text-align: right;">0.00</td>
						</tr>
									   <tr>
							<td style="text-align: left;">Jes</td>
							<td style="text-align: left;">Customer</td>
							<td style="text-align: left;">jes@live.in</td>
							<td style="text-align: left;">9744188458</td>
							<td style="text-align: right;">1131.45</td>
							<td style="text-align: right;">1131.45</td>
							<td style="text-align: right;">0.00</td>
						</tr>
										<tr>
							<td style="text-align: left;">Anvesh</td>
							<td style="text-align: left;">Both</td>
							<td style="text-align: left;">anveshtt@yopmail.com</td>
							<td style="text-align: left;">1221122112</td>
							<td style="text-align: right;">40000.00</td>
							<td style="text-align: right;">40000.00</td>
							<td style="text-align: right;">0.00</td>
						</tr>
										<tr>
							<td style="text-align: left;">anjana</td>
							<td style="text-align: left;">Both</td>
							<td style="text-align: left;">anjana.mohan@bankopen.co</td>
							<td style="text-align: left;">9746097228</td>
							<td style="text-align: right;">-2182319.98</td>
							<td style="text-align: right;">-2182319.98</td>
							<td style="text-align: right;">0.00</td>
						</tr>
						 <tr>
							<td style="text-align: left;">DHARMA TEJA</td>
							<td style="text-align: left;">Customer</td>
							<td style="text-align: left;">dharma.rao@bankopen.co</td>
							<td style="text-align: left;">8121241753</td>
							<td style="text-align: right;">29835.23</td>
							<td style="text-align: right;">29826.23</td>
							<td style="text-align: right;">0.00</td>
						</tr>
										<tr>
							<td style="text-align: left;">AMMA</td>
							<td style="text-align: left;">Customer</td>
							<td style="text-align: left;">ammm@yopmail.com</td>
							<td style="text-align: left;">4564565757</td>
							<td style="text-align: right;">61.03</td>
							<td style="text-align: right;">61.03</td>
							<td style="text-align: right;">0.00</td>
						</tr>
										<tr>
							<td style="text-align: left;">Amit</td>
							<td style="text-align: left;">Both</td>
							<td style="text-align: left;">amit.agrahari@bankopen.co</td>
							<td style="text-align: left;">9717552843</td>
							<td style="text-align: right;">-1999691.99</td>
							<td style="text-align: right;">-1999691.99</td>
							<td style="text-align: right;">0.00</td>
						</tr>
								</table>
			</div>
		</div>
	</body>
														 
	`))
	handleError(err)
	fmt.Printf("size(in bytes) : %v, error(if any): %v", towrite, err)

	// Now printing whole html string in a separate log file
	file2getString, _ := os.Create("htmlString.log")
	// now we will save printed log outputs in "logFile.log"
	log.SetOutput(file2getString)

	log.SetFlags(log.Ldate | log.Lshortfile)
	reader := io.Reader(fileRef)
	buffer := make([]byte, 5000) // buffer is for reading bytes

	n, err2 := reader.Read(buffer)
	log.Println(n, err2)
	log.Println(string(buffer)) // this statement we are not able to print in log file
	// this needs to be parse

}
