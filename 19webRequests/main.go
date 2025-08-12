package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"net/http"
)

const url = "https://beatbrawls-monad.onrender.com/api/songs"

func main() {
	fmt.Println("Hello go for web requests")
	response, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("The type of the Response is %T\n", response) //*http.Response this is a pointer means the original not the copy
	fmt.Println("Hello go for web requests: ", response)

	defer response.Body.Close() //needs to be done otherwise it will keep up the connection which is not so good, defer because we want it to end at last

	
	//better to proceed to these only if the statusCode is 200 OK
	// if response.StatusCode == 200 {
		databytes, err := io.ReadAll(response.Body) //ioutil is deprecated ig but io works (replaced ioutil)
		checkNilErr(err)
		fmt.Println("Response Body: ", string(databytes))
	// }

}

func checkNilErr(err error) {
	if err != nil {
		panic(err) //nice practice for error handling
	}
}
