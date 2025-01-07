package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to different web requests")
	const getUrlToCall = "https://api.freeapi.app/api/v1/users/current-user"
	const postUrlToCall = "https://api.freeapi.app/api/v1/users/register"

	// PerformGetRequest(getUrlToCall)
	// PerformPostRequest(postUrlToCall)
}

func PerformGetRequest(urlToCall string) {
	// lets do a get request to this url
	response, err := http.Get(urlToCall)
	isNilErr(err)

	defer response.Body.Close() // close the request as needed

	fmt.Println("Status code of this reponse -> ", response.StatusCode)

	content, err := io.ReadAll(response.Body)

	if response.StatusCode == 200 && err == nil {
		// we will just print out the content

		var reponseString strings.Builder
		byteCount, err := reponseString.Write(content)

		isNilErr(err)

		fmt.Println("Number of bytes in content -> ", byteCount)
		fmt.Println("Data in response -> ", reponseString.String())

		/*
			Another way of doing it would be
			content := string(content)
			fmt.Println("Content -> ", content)
		*/
	} else {
		// we dont have a 200 code so we have an error
		if response.StatusCode != 200 {
			fmt.Printf("Err status code %d recived \n", response.StatusCode)
		} else {
			fmt.Println("Error in reading content", err)
		}
	}
}

func PerformPostRequest(urlToCall string) {
	// lets make a request body
	requestBody := strings.NewReader(`
		{
			"email": "aditya.gadhvi@domain.com",
			"password": "test@123",
			"role": "ADMIN",
			"username": "aditya_gadhvi"
		}
	`)

	response, err := http.Post(urlToCall, "application/json", requestBody)
	isNilErr(err)

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Status code -> ", response.StatusCode)

	if response.StatusCode == 201 {
		fmt.Println("The content returned by server -> ", string(content))
	} else {
		fmt.Println("Error recived from server")
		fmt.Println("error -> ", err)
		fmt.Println("Content recived -> ", string(content))
	}

}

func PerformPostRequestWithFormData(urlToCall string) {
	// lets write the form data
	data := url.Values{}

	data.Add("firstName", "Param")
	data.Add("lastName", "Kansagra")
	data.Add("email", "paramkansagra@go.dev")

	response, err := http.PostForm(urlToCall, data)
	isNilErr(err)

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if response.StatusCode == 200 && err == nil {
		fmt.Println("Form data successfully posted")
		fmt.Println("Content -> ", string(content))
	} else {
		fmt.Println("Error -> ", err)
		fmt.Println("Content recived from website -> ", content)
	}
}

func isNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
