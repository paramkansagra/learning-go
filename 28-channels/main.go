package main

import (
	"fmt"
	"sync"
)

/*
	Channels are a way in which multiple go routines can talk to each other
		They wont be still aware about how much time they will take other go routine to complete
		So we are just waiting for some signals / information from other go routines
		So we do not need to complete and come back and do it concurrently
	We can do it on the go using channels in go routines
*/

func main() {
	fmt.Println("About channels in GO")

	/*
		Here we are defining a channel that is going to hold integer
			in channels we can hold anything like integer or strings or anything
		We can also have it to store our custom data and all as well

		We can also tell it to have a specific size of channel like 1,2,3 -> This is called a buffered channel
		In buffered channel if there are more values we are trying to push in then it wont throw up an error but it wont also consume it as well

		After pushing the value we can close the channel. -> syntax -> close(myChannelName)

		If there is any listner to a channel then there is no issue but if we try to push something to a channel after it is closed
			it would go into a panic mode.

		Also if we are listning to a closed channel it would return 0 or empty string or something empty which we might/might not be sending
		If we are sending a empty thing in the channel and the channel is already closed
			-> To fir hame kaise pata ki this is coming from a closed channel or an open channel pushing empty stuff

		So in this case we will use another syntax
			value , isChannelOpen := <- myChannel

		So this will tell us weather the channel is open or not.
		This will help us work with channels in a better way
	*/
	myChannel := make(chan int) // by default it becomes a buffered channel with size 1

	/*
		// This will generate the error called DeadLock

		myChannel <- 5 // we are pushing the value 5 to myChannel

		// lets print the values from channels
		fmt.Println(<-myChannel)
	*/

	// lets create a wait group
	waitGroup := &sync.WaitGroup{}

	/*
		In channels if we have to send the value there should be some listner who is listning to the value
		Without this you cannot have a channel

		Now if we are trying to pass too many values in the channels it will throw up an error

		Jitni values push kar rahe ho utne listners hone chaie
	*/

	// lets create a go routine as well to use these wait groups and channels
	waitGroup.Add(2)

	/*
		Channels are also bi directional in nature we can consume or send a value as well
		but in case we want to make a channel send only and recive only that is also possible

		This would be done in function defination
		To make it recive only the syntax would change to ===> myChannel <-chan int
		To make it send only the syntax would change to ====> myChannel chan<- int

		If a channel is recive only then we are not allowed to close the channels
		So if we try to close the channels it would throw up an error
	*/

	go func(myChannel chan int, waitGroup *sync.WaitGroup) {
		// fmt.Println("Value recived from myChannel -> ", <-myChannel) // one of the listner listning to the value
		// fmt.Println("Value recived from myChannel -> ", <-myChannel) // another listner listning to the value

		recivedChannelValue, isChannelOpen := <-myChannel

		if isChannelOpen {
			fmt.Println("Value recived from channel -> ", recivedChannelValue)
		} else {
			fmt.Println("Channel is closed and no value recived")
		}

		waitGroup.Done()
	}(myChannel, waitGroup)
	go func(myChannel chan int, waitGroup *sync.WaitGroup) {

		// now lets close the channel and see what we can do
		close(myChannel)

		/*
			// If we send value to closed channels it would generate an error
			// try and see

			myChannel <- 5
		*/

		/*
			fmt.Println("first value sending to myChannel -> ", 5)
			myChannel <- 5

			fmt.Println("second value sending to myChannel -> ", 6)
			myChannel <- 6
		*/

		// done with the go routine so update the waitGroup
		waitGroup.Done()
	}(myChannel, waitGroup)

	waitGroup.Wait()
}
