package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	//step 1
	reader:= bufio.NewReader((os.Stdin))//Stdin means standard input and output and hover over reader and will get to know that it is a pointer
	fmt.Println("Enter the rating for our services: ")
	//step 2
	//comma ok || comma err syntax
	//there is no try and catch in go and if something goes wrong nothing is there to catch the error
	//but we have comma ok syntax for that
	input, _:=reader.ReadString('\n')//now input is the user input and err is the error if we dont want to handle/care about the err (or anything else we simply use underscore instead of that in go) {input, err} {input, _} and here \n means that is where we stop taking the input means on a enter( or line change)
	fmt.Println("Thanks for your time: ", input)
	fmt.Printf("Type of the rating: %T",input )//type is gonna be string
}

//there are two steps in takin input from the user in Go
//using packages bufio and os and then comma ok or comma err syntax