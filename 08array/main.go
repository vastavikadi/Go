package main
import "fmt"
func main(){
//fmt.Println("This is a class lecture about arrays but arrays are not much used in golang like other prorgramming lanuages. Now lets study more about this topic.")
var fruitList [4]string//here we are just defining and not initialising
fruitList[0] = "Apple"
fruitList[1] = "Tomato"
fruitList[3] = "Apple"
fmt.Println("Fruit list is: ", fruitList)//this is output [Apple Tomato  Apple]notice there is an extra space between tomato and apple
fmt.Println("Length of the array is: ", len(fruitList))//does not matter if there is only 3 values inside the array once the length is declared tht is waht it is gonna return not gonna count the actualy no of values inside the array

var vegList = [3]string{"potato", "tomato", "apple"}//here even we make it 5 and give 3 values and then print the length then the length is gonna be 5
fmt.Println("The vegList is: ", vegList)
fmt.Println("The length of the vegList is: ", len(vegList))

//number array
var numList = [3]int{1,2, 3}
fmt.Println("the numList is: ", numList)
}
