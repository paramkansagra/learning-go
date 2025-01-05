package main

import (
	"fmt"
	"runtime"
)

func main() {
	// we will try and get some information about the cpu
	var numCpuCores int = runtime.NumCPU()
	fmt.Println("Number of cpu cores -> ", numCpuCores)

	// lets try and get where go is located
	var goRootTree string = runtime.GOROOT()
	fmt.Println("Go is located at -> ", goRootTree)
}
