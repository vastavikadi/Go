package main

import "fmt"

func main(){
	var fruitList =[]string{}//while declaring like this always remember to initialised as well (it will work even if not initialised). so in slices it is allowed to use as many values we want and memory changes as per our config . But in array we needed to give the size.
	fmt.Printf("The type of the fruitList is %T\n", fruitList)
	fruitList= []string{"Apple", "Pineapple", "Orange"}
	fmt.Println("The fruitList is:", fruitList)
	//fruitList[0]="mango"//this was used to be done in arrays but we will not use it here
	fruitList = append(fruitList, "mango", "banana")//here we will be using the append method and in this we are calling the fruitList with whatevr it had earlier and then we add two more new values and store that in fruitList fimally
	fmt.Println("The fruitList is:", fruitList)
	// fruitList = fruitList[1:]//somehow similar to what happens in python
	fruitList = append(fruitList[1:])
	fmt.Println("The fruitList is:", fruitList)


	
}
