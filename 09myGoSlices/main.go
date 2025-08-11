package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{} //while declaring like this always remember to initialised as well (it will work even if not initialised). so in slices it is allowed to use as many values we want and memory changes as per our config . But in array we needed to give the size.
	fmt.Printf("The type of the fruitList is %T\n", fruitList)
	fruitList = []string{"Apple", "Pineapple", "Orange"}
	fmt.Println("The fruitList is:", fruitList)
	//fruitList[0]="mango"//this was used to be done in arrays but we will not use it here
	fruitList = append(fruitList, "mango", "banana") //here we will be using the append method and in this we are calling the fruitList with whatevr it had earlier and then we add two more new values and store that in fruitList fimally
	fmt.Println("The fruitList is:", fruitList)
	// fruitList = fruitList[1:]//somehow similar to what happens in python
	fruitList = append(fruitList[1:])
	fmt.Println("The fruitList is:", fruitList)

	//using make keyword to define slices
	highScore := make([]int, 4) //defined the size to be 4 first
	highScore[0] = 234
	highScore[1] = 433
	highScore[2] = 554
	highScore[3] = 455
	// highScore[4]=455 will error an error because already going out of the defined size

	highScore = append(highScore, 45, 32, 56) //append reallocates the memory depending on the need
	fmt.Println("the highscores are: ", highScore)

	sort.Ints(highScore)//directly sorts the slice and append tht to the slice 
	fmt.Println(highScore)//printing the slice now and got it in ascending order because sort did it and stored in the highScore
	fmt.Println(sort.IntsAreSorted((highScore)))//returns true because it is already sorted


	//how to remove a value from slices based on index
	var techStacks = []string{"React", "Express", "Fastify", "Nest", "Next"}
	fmt.Println("Our techStacks: ",techStacks)
	var index int = 2//lets say we are going to remove fastify from the list
	techStacks = append(techStacks[:index], techStacks[index+1:]...)//using spread operator like in js to spread all the elements
	fmt.Println("This is the new techStacks after removing index 2: ",techStacks)
}
