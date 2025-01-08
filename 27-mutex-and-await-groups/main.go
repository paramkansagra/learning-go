package main

import (
	"fmt"
	"sync"
)

/*
	Race Conditions are conditions where multiple threads or go routines are trying to
	either read or write to the same memory location.
		1. This can cause either in incosistent data being written to the memory thus not fulfilling our end goal.
		2. Or it can happen ki we are reading the older data rather than reading the newer data written.
*/

func main() {
	/*
		So we want to explore the different race conditions in go
		Now to check weather a program has a race condition or not we will be using the command
			`go run --race main.go` -> this will tell us about all the race conditions that are happening in the program
		So Race conditions happen when there are multiple threads or go routines trying to write data to the same variables
		So to eliminate the race conditions we are going to use mutex
	*/

	fmt.Println("Race Condition")

	// Creating a wait group to handel the virtual threads
	waitGroup := &sync.WaitGroup{}

	// Creating a mutex for writing the data consistently
	mutex := &sync.Mutex{}

	// Creating the mutex to help us from the race conditions

	// this is the slice we are trying to write data to
	var score = []int{0}

	/*
		Now because we are going to add 3 different go routines we can write seedha waitGroup.add(3)
		it is the same as waitGroup.add(1) three times
	*/
	waitGroup.Add(3)

	/*
		Iffies or instant running functions are of the structure

		func (____{argument 1 } , ______{argument 2} , ...... ){
			______ (body of the function)
		} (___{data as argument passed to the function}
		 )
	*/

	go func(waitgroup *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("Routine 1")

		/*
			So before writing the data to the slice we will lock the slice
			So that no one else can write anything to it
			And then we will write our data
			After our write operation is completed we have to unlock it as well -> this is our responsibility
		*/

		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()

		waitgroup.Done()
	}(waitGroup, mutex)
	go func(waitgroup *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("Routine 2")

		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()

		waitgroup.Done()
	}(waitGroup, mutex)
	go func(waitgroup *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("Routine 3")

		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()

		waitgroup.Done()
	}(waitGroup, mutex)

	waitGroup.Wait()

	fmt.Println(score)
}

/*
	Thus after using the mutex the race condition would be gone and
	Thus most of our issues are also gone
*/
