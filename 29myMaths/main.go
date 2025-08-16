package main

import (
	"fmt"
	"math/big"
	// "math"
	// "math/rand"
	"crypto/rand"
)

func main() {
	// fmt.Println("Hello, World!")
	// //how to generate random numbers
	// //basically rand.Seed is deprecated and now we dont need the seed to get the random numbers each time
	// random := rand.Intn(100)//here it is returning a random number between 0 and 100
	// fmt.Println(random)
	// r := rand.New(rand.NewSource(42))// it is creating a new random number generator with a specific seed
	// fmt.Println(r.Intn(100))//here it is returning a random number between 0 and 100
	// fmt.Println(r.Float64())//here it is returning a random float64 number between 0 and 1
	// fmt.Println(r.Perm(5))//here it is returning a random permutation of the numbers 0 to 4

	//how to generate random numbers from crypto/rand
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(100))//it is returning a random number between 0 and 100
	fmt.Println(myRandomNum)
}