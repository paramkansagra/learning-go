package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch Case")

	// lets try to generate some random numbers
	// and lets seed that value using the CurrentNano second

	/*
		so time.Now() -> this will give the current time
		time.Now().UnixNano() -> will give a large number which is the number of seconds till now
	*/

	rand.Seed(time.Now().UnixNano())
	var diceNumber int = rand.Intn(6) + 1
	fmt.Println("Value of the dice is -> ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1st and you can open")
	case 2:
		fmt.Println("You can move to 2nd spots")
	case 3:
		fmt.Println("You can move to 3rd spot")
		fallthrough
	case 4:
		fmt.Println("You can move to the 4th spot")
	case 5:
		fmt.Println("You can move to the 5th spot")
	case 6:
		fmt.Println("You can move to the 6th spot and roll the dice again")
	default:
		fmt.Println("What was that ? ")
	}

	// if you want to hit all the other test cases under it after a case is hit
	// you have to write "fallthrough" to run after that
}
