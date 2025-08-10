package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to our trails"
	fmt.Println(welcome)
	fmt.Println("Please rate our services and help us understand and improve our flaws:")
	reader:= bufio.NewReader((os.Stdin))
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the ratings. ", input)
	fmt.Printf("Type of the ratings is: %T \n", input)

	//but by default it is string but i want to convert to some other type
	//cmd to run programs : go run main.go

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) gives error because along with the input a \n also goes in and this leads to the error and for that we need the strings package to trim and then pass
	if err!= nil{//nil means null or absense
		fmt.Println(err)
	} else{
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}