package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
	Wait groups is nothing but a group of threads which would be saying that
	We have a go-routine which is still working and we have to wait till it completes
	And until it completes the main function should not end

	There are mainly 3 functions in wait groups
		1. Add(int n) -> This will tell how many go threads are added to the wait group
		2. Done() 	  -> This will tell the wait group that one of the thread has done its work and shouldnt wait for them
		3. Wait()	  -> This will tell the main function to wait till the go-routines havent completed their work
*/

var waitGroup sync.WaitGroup // usually they are pointers

var signals = []string{"test"}

/*
	Now we are using mutex in go routines this is because
		In case there are multiple go routines which are trying to read/write a particular memory
		Then there might be a issue where correct data wouldnt be written/read from the memory
	So we add mutex to say that this memory is currently locked and you cannot read/write from this memory

	There are different types of mutex avaliable like sync.Mutex and sync.RWMutex

	sync.Mutex has the following properties :-
		1. Lock() -> Locks the mutex for read/write
		2. Unlock() -> Unlocks the mutex for read/write

	sync.RWMutex has the following properties :-
		1. Lock() -> Locks the mutex for writing
		2. RLock() -> Locks the mutex for reading
		3. RUnlock() -> Unlocks the mutex for reading
		4. TryLock() -> TryLock tries to lock mutex for writing and reports whether it succeeded (very rare)
		5. TryRLock() -> TryRLock tries to lock mutex for reading and reports whether it succeeded. (very rare)
		6. Unlock() -> Unlocks the mutex for writing
*/

var mutex sync.Mutex

func main() {

	websiteList := []string{"https://google.com", "https://go.dev", "https://meta.com", "https://instagram.com", "https://github.com", "https://reddit.com"}

	for _, endpoint := range websiteList {
		// fire new threads for each of the website
		go getStatusCode(endpoint)
		waitGroup.Add(1) // how many go-routines are added?
	}

	waitGroup.Wait() // this is responsible to not let the main function end before all the threads have done their tasks
	fmt.Println(signals)
}

/*
	Now whenever we fireup new threads we will push that thread into a WaitGroup
	So waitgroup would be responsible for management of go routines(threads)
	Once the GoRoutine is completed its our responsibility for calling out Done()
		so as to tell the wait group that yea my work is over

*/

func getStatusCode(endpoint string) {

	/*
		here we are saying once this function ends
		tell the wait group i am done and they can remove me
	*/
	defer waitGroup.Done()

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Err in endpoint")
	} else {
		answer := fmt.Sprintf("%d status code for %s \n", response.StatusCode, endpoint)

		fmt.Print(answer)

		mutex.Lock()                      // Lock the mutex till we have written our answer
		signals = append(signals, answer) // write our answer to the signal
		mutex.Unlock()                    // Unlock the mutex because we are done here

		// this will lead to perfect working every time we run without errors
	}
}
