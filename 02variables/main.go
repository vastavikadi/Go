package main

import "fmt"

// jwtToken := 79327401 walrus operator cannot be used outside of a method like here (main)

//these can be used
// var variable2 int =47932
// var variable3 =47932

//declaration of constants
const loginToken string = "84yhu3sh4"  //this is the way of defining the constants
const LoginToken string = "842ninsine" //now that the first letter of the variable is capital it means it is public and can be accessed from any other file just like Public of OOPs

func main() {
	var name string = "Aditya"
	fmt.Println(name)
	fmt.Printf("The type of the variable is: %T \n", name) //%T is a placeholder for types just like in C %d is a format specifier for integers
	// fmt.Println("The type of the variable is: %T", name) so %T does not work with Println only works with Printf

	var smallValue uint8 = 255 //uint stands for unsigned int and uint goes from 0 to 255 as 2^8 = 256 hence it contains 256 digits (from 0 to 255)
	fmt.Println(smallValue)
	fmt.Printf("The type of the variable is: %T \n", smallValue)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("The type of the variable is: %T \n", isLoggedIn)

	var largeValue uint32 = 764698363 //adding even a single digit means going out of the limits for what set in unit32
	fmt.Println(largeValue)
	fmt.Printf("The type of the variable is: %T \n", largeValue)

	var decimalValue float32 = 432.4532
	fmt.Println(decimalValue)
	fmt.Printf("The type of the variable is: %T \n", decimalValue)

	//default values and aliases
	var anotherVariable int //here not initlialised but by default got 0 as its value
	fmt.Println(anotherVariable)
	fmt.Printf("The type of the variable is: %T \n", anotherVariable)

	var anotherVariable2 string //here not initlialised but by default gets an empty string (prints nothing)
	fmt.Println(anotherVariable2)
	fmt.Printf("The type of the variable is: %T \n", anotherVariable2)

	//by default zero is assigned as a value in Go without any initial value and hence false for bool and empty string for a string

	//implicit type/way of declaring a variable
	var website = "Hey Its me" //no error because the lexer already treated it as a string on the basis of the initial value we gave
	fmt.Printf("The type of the variable is: %T \n", website)
	// website = 56 // cannot do this because already being treated as a string
	// fmt.Println('Hey')//do not run if got inside the single quotes and not double

	//no var style
	numberOfUser := 300000 //being defined without the use of var keyword and this :+ called (declare and initialize) short variable declaration operator and called walrus operator in python
	fmt.Println(numberOfUser)
	// numberOfUser=44.43 cannot change the type

	fmt.Println(loginToken)
	fmt.Println(LoginToken) //even the variables are case sensitive and This means that identifiers like age, Age, and AGE are considered three distinct and different variables.
	//  a variable starting with an uppercase letter is exported and accessible outside its package, while a variable starting with a lowercase letter is unexported and only accessible within its package.
	//  This case sensitivity applies to all identifiers, including variable names, function names, and type names.
}
