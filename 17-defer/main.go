package main

import "fmt"

/*
	So defer statements are the statements which are executed just before the function ends
		i.e. agar we have the function called pypy and isme ek defer statement hai to fir
			just before the function pypy ends the defer statement will execute then it will return

		Now if there are multiple defer statements then they would be executed in LIFO(last in first out order)
			Jo defer statement pehle add kia hai vo sabse last me execute hoga
			and jo statement sab se last me add hua hai vo sabse pehle execute hoga
*/

func myDeferFunction() {
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}

	// here because we are defering all the statements
	// LIFO stack would be like -> 0 1 2 3 4 5 6 7 8 9
	// thus it would be printed like -> 9 8 7 6 5 4 3 2 1 0
}

func main() {
	fmt.Println("Hello welcome to Defer statements")
	defer fmt.Println("Byeee")

	fmt.Println("This is after the defer statement")

	myDeferFunction()
}
