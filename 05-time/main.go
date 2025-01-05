package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time in golang")

	// we are trying to get the current time
	timeNow := time.Now()
	fmt.Println("The present this is -> ", timeNow)

	// we are trying to format the time
	// now why we are using 01-02-2006 -> because the documentation says us so
	fmt.Println("The date is -> ", timeNow.Format("01-02-2006"))

	// why moday -> not because today is monday but DOCUMENTATION
	fmt.Println("The date is -> ", timeNow.Format("01-02-2006 Monday"))

	// why 15:04:05 because DOCUMENTATION
	fmt.Println("The date is -> ", timeNow.Format("01-02-2006 Monday 15:04:05"))

	// lets try creating a date
	createdDate := time.Date(2020, time.December, 20, 10, 2, 1, 0, time.Local)
	fmt.Println("The created date is -> ", createdDate.Format("01-02-2006 Monday"))

	// we can also get the nano second and unix nano second for random values
	nanoSecond := timeNow.Nanosecond()
	fmt.Println("The current nano seconds -> ", nanoSecond)

	unixNanoSecond := timeNow.UnixNano()
	fmt.Println("Unix nano second -> ", unixNanoSecond)
}

// to build file we can write -> go build
//this will build for the operating system defined in GOOS env variable
// now to make it for other operating system we can write
// GOOS="windows" go build -> this will build for windows
// GOOS="linux" go build -> this will build for linux
// GOOS="darwin" go build -> this will build for macos
