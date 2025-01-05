package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}

	fmt.Println("The days list is like -> ", days)

	// lets loop over it and check it out
	/*
		for index := 0; index < len(days); index++ {
			fmt.Printf("The index is -> %v the day is -> %v \n", index, days[index])
		}
	*/

	// other way of doing this is
	/*
		for i := range days {
			fmt.Printf("The index is -> %v the day is -> %v \n", i, days[i])
		}
	*/

	// we can do it in one more way
	for index, day := range days {
		fmt.Printf("They index -> %v the day -> %v \n", index, day)
	}

	/*
		insted of a dedicated while loop
		the while loop is built into the for loop
	*/

	var rogueValue int = 1
	for rogueValue < 10 {
		fmt.Printf("The rogue value is -> %d \n", rogueValue)

		if rogueValue == 5 {
			// to jump to a label we will write goto and the label name
			goto jumpToHere
		}

		rogueValue++
	}

jumpToHere:
	fmt.Println("Jumping out of the loop")

	/*
		Here with the this we also have jump to statements where we can jump seedha to a statement
	*/

}
