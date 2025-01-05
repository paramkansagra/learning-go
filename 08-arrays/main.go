package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to arrays")

	// lets make a list of cars
	var carList [4]string
	carList[0] = "BMW"
	carList[1] = "Audi"
	carList[2] = "Merc"
	carList[3] = "Ferrari"

	fmt.Println("The car list is -> ", carList)
	fmt.Println("The length of car list is -> ", len(carList))

	// lets see and check if we can get a input from the user and make a list of size
	/*
		reader := bufio.NewReader(os.Stdin)
		nStr, _ := reader.ReadString('\n')

		n, _ := strconv.ParseInt(strings.TrimSpace(nStr), 10, 64)
	*/

	// var newList [n]string // -> not possible beacause it says we have to make it of a fixed size

	var bikeList = [3]string{"BMW", "Kawasaki", "Suzuki"}
	fmt.Println("The bike list is -> ", bikeList)
}
