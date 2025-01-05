package main

import "fmt"

func main() {
	fmt.Println("Welcome to If-Else in GoLang")

	var loginCount int = 10
	var resultMessage string

	// It is a normal if else without the brackets
	if loginCount < 10 {
		resultMessage = "Regular User"
	} else if loginCount > 10 {
		resultMessage = "Watch out"
	} else {
		resultMessage = "Excatly 10 login counts"
	}

	fmt.Println("Result message -> ", resultMessage)

	//  we can also do some on the go arethmetic and do the checking as well
	if loginCount%2 == 0 {
		fmt.Println("Login count is even")
	} else {
		fmt.Println("Login count is odd")
	}

	// we can also do on the go assignment and then checking as well for web requests
	if num := 3; num < 3 {
		fmt.Println("On the go assigning and checking and less than 3")
	} else {
		fmt.Println("On the go assigning and checking and more than 3")
	}
}
