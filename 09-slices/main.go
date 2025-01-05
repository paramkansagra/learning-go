package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var carList = []string{"BMW", "MERC", "AUDI"}               // when we are not defined the length or the values in it then it becomes a slice
	fmt.Printf("The data type of carList is -> %T \n", carList) // []string -> slices

	// can we add more values ?? YESSS
	carList = append(carList, "Ferrari") // append(nameOfSlice , element1 , element2 ...) or append(nameOfSlice , anotherSliceWeWantToAppend)
	fmt.Println("The values in the car list is -> ", carList)

	// now if we want some part of the array we can also use this append function
	newCarList := append(carList[1:3])
	fmt.Println("The new car list is -> ", newCarList)

	newNewCarList := append(carList[:3])
	fmt.Println("The new car list is -> ", newNewCarList)

	// lets make another slice using make
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 223
	highScores[2] = 222
	highScores[3] = 221

	// lets get the data structure of it
	fmt.Printf("Data type of highScores is -> %T \n", highScores)
	fmt.Println("Values stored in highScores -> ", highScores)

	// lets append some values
	highScores = append(highScores, 4)                         // append will re allocate the whole memory and add the new value
	fmt.Println("Values stored in highScores -> ", highScores) // thus allocate is very expensive to do

	// lets try to sort the slice
	sort.Ints(highScores) // now lets print them
	fmt.Println("Sorted values of highScores -> ", highScores)
	fmt.Println("Is the array sorted -> ", sort.IntsAreSorted(highScores))

	// remove a value from slices based on index
	var courses = []string{"reactjs", "js", "swift", "python", "ruby"}
	fmt.Println("Courses -> ", courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // ... will just expand the list into individual elements

	fmt.Printf("After deleting the %d -> %v \n", index, courses) // %v => for printing the slices
}
