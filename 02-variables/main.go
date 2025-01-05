package main

import "fmt"

// we can also declare variables here
// := (walrus) operator is not allowed outside

// var jwtToken string = "" // this is still allowed
const LoginToken string = "fasdafasdfa" // constant variable
// capital L means that it is a public variable that anyone can even outside this file
// so be careful while naming the variables and all

func main() {

	// for a variable declaration we can write as -> var(keyword) ____(name of the variable) ____(data type)
	var username string = "Param Kansagra"
	fmt.Println(username)
	fmt.Printf("variable is of the type -> %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of the type -> %T \n", isLoggedIn)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("variable is of the type -> %T \n", smallInt)

	var smallFloat float32 = 255.2
	fmt.Println(smallFloat)
	fmt.Printf("variable is of the type -> %T \n", smallFloat)

	// default values and alias
	var anotherDefaultVariable int
	fmt.Printf("The default value of an %T is -> %d \n", anotherDefaultVariable, anotherDefaultVariable)

	// another way of declaring a variable is

	// implicitly declaring the variable completly skipping the type of the variable
	// this would be declared implicitly by the lexer based on the value we are passing
	var website = "hello.com" // once declared in this way we cannot change the data type
	fmt.Println(website)

	// another way is "no var style" using :=
	numberOfUsers := 2000 // using the walrus operator
	fmt.Println(numberOfUsers)
}
