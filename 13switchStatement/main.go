package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber:= rand.Intn(6)+1//populated with rand and then 6 means generating 6 numbers from 0 to 5 and hence adding 1 to each to make it 1 to 6

	fmt.Println("THe number on the dice is : ", diceNumber)

	switch diceNumber{
	case 1://1 is the value from diceNumber and in case of string case "Monday"
		fmt.Println("Dice number is: 1 and you can move 1")
	case 2:
		fmt.Println("Dice number is: 2 and you can move 2")
	case 3:
		fmt.Println("Dice number is: 3 and you can move 3")
		fallthrough//used it in 3 and hence gonna include the action of next ( here 4 as well)
	case 4:
		fmt.Println("Dice number is: 4 and you can move 4")
	case 5:
		fmt.Println("Dice number is: 5 and you can move 5")
	case 6:
		fmt.Println("Dice number is: 6 and you can move 6. throw the dice once more")
	default:
		fmt.Println("Default")
	}
}//in the golang playground it always gives 3 because of their security measures 