package main

import "fmt"

func main() {
	defer fmt.Println("hello world") //defer is now gonna remove the line from here and will add just before the curly braces means gonna be executed at last
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three") //multiple defer follows LIFO execution last in and first out hence three executes first/ before one and two
	fmt.Println("hello world no defer")

	//the loop using defer
	fmt.Println("The loop: ")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)//here 4 will be the last one to be pushed inside the loop hence executes first, defer stores the print in the stack and then when it ends then the stack starts executing (LIFO)
	}
}
