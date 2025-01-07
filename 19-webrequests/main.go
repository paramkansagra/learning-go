package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "http://api.open-notify.org/iss-now.json"

func main() {
	fmt.Println("Welcome to web requests")

	responce, err := http.Get(url) // trying to do a get request to the url
	checkNilErr(err)               // Now we will check if there is any error or not

	fmt.Printf("Responce is of the type -> %T \n", responce) // the type of responce is *http.Responce
	defer responce.Body.Close()                              // We are doing this because its our responsibility to close the connection

	responceBytes, err := io.ReadAll(responce.Body)                                        // lets read the responce using io and we will read the body of the get request
	checkNilErr(err)                                                                       // checking if there is error or not
	responceString := string(responceBytes)                                                // convert the byte data into string data
	fmt.Println("The data contained in the body of the get request is ->", responceString) // its some json data about location of iss right now

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
