package main

import "fmt"

func main() {
	fmt.Println("Maps in Go-Lang")
	/*
		Another way to make maps is by using the inbuilt var
		var myMap[string]string

		here the first [DataType] is the data type of the key
		And the other one is the data type of the value which is against the key
	*/

	/*
		one way of creating the map is by using make

		there are 2 ways of making any slice or map or any struct or anything
		1. make -> will allocate the memory based on the data structure and also init it
		2. new -> it will allocate 0 memory and also will not init the data structure we have to do it ourselves

		Both will be returning a pointer to the data structure
	*/
	lanuages := make(map[string]string)

	lanuages["JS"] = "JavaScript"
	lanuages["PY"] = "Python"
	lanuages["CPP"] = "C++"

	fmt.Println("The languages map is -> ", lanuages)
	fmt.Println("JS is short for ->", lanuages["JS"])

	// deleting value from the map is like ->
	delete(lanuages, "JS")
	fmt.Println("New map is like -> ", lanuages)

	// lets loop thru the map and check
	for key, value := range lanuages {
		fmt.Printf("Key -> %v Value -> %v \n", key, value)
	}
}
