package main

import "fmt"

func greeter() {
	fmt.Println("Namaste saab")
}

/*
	So the function body would look like
		func ___{functionName} (____{variableName} ___{dataType} , ...) ____{return type / function signature if any} {
			_____ {main function body}
		}
*/

func adder(valOne int, valTwo int) int {
	return (valOne + valTwo)
}

/*
	Lets make a function where we dont know the amount of arguments we are getting
	we will add all of them and return the answer value
*/

func proAdder(values ...int) int {
	// if you hover over values then it is a slice of the integers

	var ans int = 0
	for _, val := range values {
		ans += val
	}

	return ans

	/*
		We can also do something like _,OK syntax by returning multiple things like this ->
		return ans , "Hello from the pro adder function"

		and in header insted of writing the return value as int
		we have to write as (int , string) // with the brackets
	*/
}

func main() {
	greeter()
	fmt.Println("Welcome to functions in golang")

	result := adder(10, 15)

	fmt.Println("The result is ->", result)

	result = proAdder(10, 20, 30, 40, 50)
	fmt.Println("The new result is -> ", result)
}
