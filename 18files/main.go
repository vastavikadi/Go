package main

import (
	//Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is the file content we will be writing in a file using GO."

	file, err := os.Create("./myFirstFile.txt") //creating the file using the os module
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content) //this returns a length (can read from the docs and can also have errors)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println("The length of the created file is: ", length)

	readFile("./myFirstFile.txt")
}
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename) //always read in bytes not in string (even from the internet or json)
	// if err != nil {
	// 	panic(err)//not good if using again and again should have a func for it
	// }
	checkNilErr(err)
	fmt.Println("The text data inside the file is: ", databyte) //in bytes the data comes like this: [84 104 105 115 32 105 115 32 116 104 101 32 102 105 108 101 32 99 111 110 116 101 110 116 32 119 101 32 119 105 108 108 32 119 114 105 116 105 110 103 32 105 110 32 97 32 102 105 108 101 32 117 115 105 110 103 32 71 79 46] need to convert to string

	textByte := string(databyte) //so first we needed to convert the raw bytes to a string and then print it and instead of directly overwriting the databyte variable i declared a new string variable for the task
	fmt.Println("The text data inside the file is: ", textByte)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err) //nice practice for error handling
	}
}
