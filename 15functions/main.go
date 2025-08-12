package main

import "fmt"

func main(){
	fmt.Println("Welcome to the functions in go")
	greeting()//here we are calling the function after definition
	// func greeterTwo(){//this is not allowed, you do not declare a function inside another func(main)
	// 	fmt.Println("Hey")
	// }
	greeterTwo()

	//sum function
	var result int = sum(5,6)
	fmt.Println("The sum of the two numbers is: ", result)
}

// Anonymous functions can be used inside other functions (like inside the main), not at the top level in Go
// func(){
// 	fmt.Println("ANonymous functions exist in GO as well. IIFE like in JS")
// }()//this is IIFE like in JS

func greeterTwo(){
	fmt.Println("Hey")
}
func greeting(){
	fmt.Println("Namaste")//declaring the function
}

//notice i have written func sum(argu...)int{}//means the return type of the function is int only
func sum(a int, b int) int{//The problem is that in Go, the type comes after the variable name in function parameters, so you should write a int, b int instead of int a, int b.
	var sum int = a+b
	return sum
}