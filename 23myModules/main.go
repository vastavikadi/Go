package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {//A ResponseWriter interface is used by an HTTP handler to construct an HTTP response. A ResponseWriter may not be used after [Handler.ServeHTTP] has returned.
	w.Write([]byte("<h1>Hell Yeah</h1>"))//writing a response
}

//go run, go build and go mod are the tools. go run runs and go build builds the app
//go mod keeps the track of all the modules installed to build the project/app along with their versions (more like dependency managament system)
//gorilla mux used for routing

//go get is a tool used to pull in dependency from the web
//go mod has indirectly infront of dependencies which they are not being used yet

//ABOUT go.sum
// It stores cryptographic hashes (checksums) of the exact versions of modules you’ve used
// These hashes let Go verify that the module you download in the future is bit-for-bit identical to the one you originally got.
// If the file contents change (e.g., someone tampers with the module), Go will refuse to build.

// go.mod → Records what modules and versions your project requires.
// go.sum → Records hashes of those modules so they can be verified.
// Think of go.mod as a shopping list, and go.sum as the receipts proving you got exactly what you ordered.

// go mod tidy
// Add: Finds any imports in your code that aren’t listed in go.mod and adds them with the correct versions.
// Remove: Finds any dependencies in go.mod that aren’t actually used in your code and removes them.
// Update go.sum: Ensures the go.sum file has the correct checksums for all required modules (and nothing extra).

//go mod verify this verifies all the modules you have used
//go list -m all packages being used
//go list all :- everything installed

//go list -m -versions github.com/gorilla/mux
// github.com/gorilla/mux v1.2.0 v1.3.0 v1.4.0 v1.5.0 v1.6.0 v1.6.1 v1.6.2 v1.7.0 v1.7.1 v1.7.2 v1.7.3 v1.7.4 v1.8.0 v1.8.1 listed all the versions

//go mod why github.com/gorilla/mux :_ tells where being used
//go mod graph :- tells packages dependent on other packages

//go mod edit -go :- can edit the versions of go being used


//once you have fetched the packages from the web it is stored in cache locally but not everything is there and if you need something extra or the whole code use go mod vendor:- fetches everthing from the web and then stores in a vendor folder and uses from there