package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var welcomeMessage string = "Welcome to user input"
	fmt.Println(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza := ")

	// it is called comma,ok || err,ok things (when we are not using the variable we can just do _)
	inputRating, _ := reader.ReadString('\n') // so we are going to read till we will get a \n

	fmt.Println("Thanks for the rating -> ", inputRating)
	fmt.Printf("Type of the rating is %T \n", inputRating)
}
