package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type course struct {
	Name     string `json:"courseName"` // -> when we convert to json it would do based on the alias we provider here
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // meaning of "-" in json means that when we are consuming the api this field wont be given
	Tags     []string `json:"tags,omitempty"` // here "Tags" key would be replaced by "tags" and if there is nothing in here it would be omitted and not sent
}

type joke struct {
	Categories []string `json:"categories"`
	Id         int      `json:"id"`
	Content    string   `json:"content"`
}

type incomingRequest struct {
	StatusCode int    `json:"statusCode"`
	Data       joke   `json:"data"`
	Message    string `json:"message"`
	Success    bool   `json:"success"`
}

func main() {
	fmt.Println("Welcome to the world of JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{Name: "ReactJS", Price: 200, Platform: "Udemy", Password: "Hemlo", Tags: []string{"Web Dev", "Javascript"}},
		{Name: "Flutter", Price: 100, Platform: "Udemy", Password: "Hi", Tags: []string{"App Dev", "Android", "iOS", "Linux"}},
		{Name: "HTML", Price: 300, Platform: "Udemy", Password: "Crazy", Tags: nil}, // see in the final output for this Tags wont be there and password is not reflected anywhere
	}

	// we will convert the structs into json using Marshal
	// MarshalIndent is used to do in a more readable format
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	isNilErr(err)

	// we have the final json and we will print it
	fmt.Println("final json -> ", string(finalJson))
}

func DecodeJson() {
	const webUrl string = "https://api.freeapi.app/api/v1/public/randomjokes/joke/random"

	response, err := http.Get(webUrl)
	isNilErr(err)

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	isNilErr(err)

	var request incomingRequest
	var jomke joke

	// now we have byte json in content and we will convert it to readable form
	if json.Valid(content) {
		fmt.Println("Json is valid")
		err := json.Unmarshal(content, &request)
		isNilErr(err)

		jomke = request.Data
		fmt.Printf("%+v \n", jomke)
	} else {
		fmt.Println("Error in getting the content")
	}

	// some cases when we just need to add the json data to a key value pair
	var onlineData map[string]interface{} // here interface means that the value can be anything int string array slice anything
	err = json.Unmarshal(content, &onlineData)

	isNilErr(err)

	// lets print the key value pairs we have
	fmt.Printf("data -> %+v type of data -> %T \n", onlineData , onlineData)
}

func isNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
