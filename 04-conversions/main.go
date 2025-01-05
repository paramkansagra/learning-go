package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	// making a new reader for input
	reader := bufio.NewReader(os.Stdin)

	// reading the string and the err from reader
	inputRating, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", inputRating)

	// we are now trying to convert it to float
	// we are having this rating and 64 bits float to be converted upon
	// numRating, err := strconv.ParseFloat(inputRating, 64) -> will return a error as 5\n -> not trimming the end

	numRating, err := strconv.ParseFloat(strings.TrimSpace(inputRating), 64) // now we are trimming the space from the inputReading and get the rating

	// if there is some error
	if err != nil {
		fmt.Println("The error is", err)
		// panic() ->  we can cause a panic situation and close the program
	} else {
		fmt.Println("Added 1 to your rating ", numRating+1)
	}
}
