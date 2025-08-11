package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())//basically it returns the number of logical CPUs reuable by the current process
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))//how many Go is actually using right now.
}


// returns the number of logical CPUs (hardware threads) that the Go runtime can potentially use for scheduling goroutines.

// Key points:

// Logical CPUs include physical cores plus hyper-threaded ones.

// This value is determined when the program starts, based on your machine’s CPU and OS settings.

// It doesn’t necessarily mean your program will use all of them — by default, GOMAXPROCS controls how many logical CPUs the Go scheduler uses simultaneously.

// You can change it with runtime.GOMAXPROCS(n) if you want to limit or increase CPU usage.