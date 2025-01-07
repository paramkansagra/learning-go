package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "http://api.open-notify.org/iss-now.json?place=India&time=220"

func main() {
	fmt.Println("Welcome to handling urls in GO")
	fmt.Println("The current url ->", myUrl)

	// lets get the different chunks of the url
	parsedUrl, err := url.Parse(myUrl)
	checkNil(err)

	fmt.Println("url scheme ->", parsedUrl.Scheme)
	fmt.Println("host ->", parsedUrl.Host)
	fmt.Println("path ->", parsedUrl.Path)
	fmt.Println("raw query ->", parsedUrl.RawQuery)
	fmt.Println("port number ->", parsedUrl.Port())

	queryParams := parsedUrl.Query()
	fmt.Printf("The type of query params is -> %T \n", queryParams) // url.Values -> they are key value pairs in a very fancy way

	// lets get some query parameters of the url
	fmt.Println("They place is -> ", queryParams["place"])
	fmt.Println("The time is -> ", queryParams["time"])

	// lets print out all of them
	for key, value := range queryParams {
		fmt.Println(key, value)
	}

	// lets construct a url using different parts
	partsOfUrl := &url.URL{ // remeber that we have to give &
		Scheme: "https",
		Host:   "api.open-notify.org",
		Path:   "iss-now.json",
	}

	finalUrl := partsOfUrl.String() // string of the final url
	fmt.Println("final url -> ", finalUrl)
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
