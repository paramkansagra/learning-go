package main

import "fmt"

/*
	Methods are just like functions but the main difference is
		when the functions are attached to structs they become methods

	Usual syntax of the method is like ->
		func (_____{which struct we are passing in}) _____{method name} (___{any arguments if any}){
			_____ (main function body)
		}
*/

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active?", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is ->", u.Email)
}

func main() {
	var param User = User{"Param Kansagra", "paramkansagra@go.dev", true, 21}
	// now lets try printing the whole user object
	fmt.Printf("user details -> %+v \n", param)

	// now lets call the method as well
	param.GetStatus()

	// when we are passing this we are making a copy and send it to the method
	param.NewMail()

	// thus this wont change the previous data
	fmt.Println("Email of this user is ->", param.Email)
}
