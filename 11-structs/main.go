package main

import "fmt"

/*
	There is no Inheritance in GoLang
		There is 1 single class that would be there without any inheritance or anything

	With no Inheritance there is no child class or super class or anything
*/

/*
	The Syntax of defining struct is
		type ____(Struct Name) struct
		{
			_____(Variable Name) _____(Data type)
			_____(Variable Name) _____(Data type) --> notice that there is no ,
		}
*/

// capital meaning its a public class as well as public variables
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to Structs")

	// we are using the walrus operator to make a new object with ease
	param := User{"Param", "ParamKansagra@go.dev", true, 20}
	fmt.Println("The object looks like -> ", param)

	/*
		Other way of declaring the object is like ->
			var otherObject User = User{"Param", "ParamKansagra@go.com", true, 20}
	*/

	// To print the object in more detail we can use %+v
	fmt.Printf("Object in more detail -> %+v \n", param)

	// To print any one value from the structure would be like ->
	fmt.Printf("Name of the obejct is like -> %v and the email is -> %v \n", param.Name, param.Email)
}
