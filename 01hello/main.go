package main

import "fmt"

func main() { //just like C programming language main is the entry point for the program
	fmt.Println("Hello, World!")
}

// The formal syntax uses semicolons ";" as terminators in a number of productions. Go programs may omit most of these semicolons using the following two rules:

// When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a line's final token if that token is
// an identifier
// an integer, floating-point, imaginary, rune, or string literal
// one of the keywords break, continue, fallthrough, or return
// one of the operators and punctuation ++, --, ), ], or }
// To allow complex statements to occupy a single line, a semicolon may be omitted before a closing ")" or "}".

//data types
// In Go almost everything is a type just like in js everything is an object
// Go is a statically typed language, which means that variable types are explicit and checked at compile time. The basic data types in Go include:

// 1. Boolean: Represents true or false values.
// 2. Numeric: Includes integers (int, int8, int16, int32, int64), unsigned integers (uint, uint8, uint16, uint32, uint64), and floating-point numbers (float32, float64).
// 3. String: Represents a sequence of characters.
// 4. Array: A fixed-size sequence of elements of the same type.
// 5. Slice: A dynamically-sized, flexible view into the elements of an array.
// 6. Map: A collection of key-value pairs, where each key is unique.
// 7. Struct: A composite data type that groups together variables (fields) of different types.
// 8. Pointer: A variable that stores the memory address of another variable.
// 9. Complex: A type that represents a complex number, with real and imaginary parts.


// Interface: A type that specifies a contract by defining a set of method signatures (but not their implementations).
// Channel: A communication mechanism that allows goroutines to synchronize and exchange data.
// Function: A block of code that performs a specific task and can be reused.
