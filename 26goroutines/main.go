package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //usually these are pointers
var mut sync.Mutex    // pointer

func main() {
	// go greeter("Hello")//these are two different functions which were going to print hello and world 5 times one after another but what if we wnated to print hello once and then world once and so on. The way is to go with the goroutines and to use go routines we use keyword go and thats what i did. But this way it is going to print the world 5 times and no hello because we never told that go routine when to come back and print hello. now for that we used time.Sleep but that is not the ideal way
	// rather we can use channels to communicate between goroutines and wait groups to synchronize them
	//maybe you want to use 3 apis for data so we can use wait groups to wait for all of them to finish and collect the data and return 
	// greeter("world")
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)//without the go routines it will hitting lco then go then google and so on but with the concurrency it will hit all of them at the same time 
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}
}