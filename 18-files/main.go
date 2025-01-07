package main

import (
	"fmt"
	"io"
	"os"
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	// to read the files we can use the io util files
	dataBytes, err := os.ReadFile(fileName)
	checkNilErr(err)

	// converting the data bytes into string
	dataString := string(dataBytes)

	// printing the string data
	fmt.Println("Text data in the files -> ", dataString)
}

func writeFile(filepath string, contents string) {
	// first we will try to create the file
	file, err := os.Create(filepath)
	checkNilErr(err)

	// once we have the file open
	// when the function is finished we have to close the file so using defer
	defer file.Close()

	// now we have the file and we can write using that
	lengthOfData, err := io.WriteString(file, contents)
	checkNilErr(err)

	fmt.Println("Length of data written -> ", lengthOfData)
}

func main() {
	fmt.Println("Welcome to File IO in GO")

	var content string = "This data must go to a file"
	var filePath string = "./myGoFile.txt"

	/*
		One method to write data is using OS
		1. Make a file of our choice in the preffered dir we want to
		2. There might be a possibility that the dir we are trying to write it might be reserved dir
			so it might throw an error
		3. after check if there is any error or not we will proceed to writing data into the file
	*/

	writeFile(filePath, content)
	readFile(filePath)
}
