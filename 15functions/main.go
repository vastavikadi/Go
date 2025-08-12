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
	result := sum(5,6)
	// var result int = sum(5,6)
	fmt.Println("The sum of the two numbers is: ", result)

	//can add as many number as we want
	proSum, _:= proAdder(4,5,3,3,4,5,3,4,5,3,2,2,5,5,6,7,6,44)//ignoring the string that is being returned thru comma err syntax
	fmt.Println("The sum of all these numbers is: ",proSum)
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
//function signatures :- what are you gonna pass in the function and what that func will return
func sum(a int, b int) int{//The problem is that in Go, the type comes after the variable name in function parameters, so you should write a int, b int instead of int a, int b.
	var sum int = a+b
	return sum
}

func proAdder(values ...int)(int, string){//since a lot of values hence it is a slice now and hence for loop works on it 
	total:=0
	for _,val:= range values{
		total=total+val
	}
	// for _,val:= range values{
	// 	total=total+val
	// }
	return total,"Hey I am Aditya"
}