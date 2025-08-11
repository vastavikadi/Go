package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)
	//only for is there in go for loops no other like while or do while

	//type 1 for loop
	// for d := 0; d < len(days); d++ {//no concept of ++d in go
	// 	fmt.Println(days[d])
	// }

	//type 2 for loop
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	//type 3 for loop
	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n",index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("index is  and value is %v\n", day)
	// }//ignoring the index here

	rougueValue := 1

	for rougueValue < 10 {//similar to while loop in other languages

		if rougueValue == 2 {
			goto lco//break also is here in Go similar to other language and goto breaks out to execute something out of the loop
		}

		if rougueValue == 5 {
			// continue//using before makes everything written under this unreahable like return
			rougueValue++
			// continue//if we didnt write rougueValue++ then understand continue continues the loop and break breaks out of it but on continue again we get the value to be 5 as we are not updating the value so we need to do that to continue further in the loop
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

lco://lco can be any other name
	fmt.Println("Hey 2")
}