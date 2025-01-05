package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers in GO")

	var intPtr *int // -> this a pointer pointing to a integer value in memory
	// var stringPtr *string // -> this is a pointer pointing to a string value in memory

	fmt.Println("Value of integer pointer is -> ", intPtr) // -> initially if no value is given it has the value <nil>

	var myNumber int = 23
	intPtr = &myNumber                                // -> now intPtr would be pointing to the myNumber value
	fmt.Println("Pointer is pointing to -> ", intPtr) // & -> as always this gives the address of the variable
	fmt.Println("Getting value thru variable -> ", myNumber)
	fmt.Println("Getting value thru pointer -> ", *intPtr)

	*intPtr = *intPtr * 2 // we are changing the value pointed by the pointer by 2
	fmt.Println("new value thru pointer -> ", *intPtr)
	fmt.Println("new value thru variable -> " , myNumber)
}
