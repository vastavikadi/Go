package main

import "fmt"

func main() {
	//structs are alternative of classes in golang
	//go has no concept of inheritance and no super or parent

	Aditya:= User{"Aditya", "vastavikadi@gmail.com", true, 18}
	fmt.Println(Aditya)//{Aditya vastavikadi@gmail.com true 18}
	fmt.Printf("The type of the variable is: %T", Aditya)//main.User
	fmt.Printf("Aditya's details are: %+v\n", Aditya)//{Name:Aditya email:vastavikadi@gmail.com Status:true Age:18}
	fmt.Printf("Aditya's name is: %+v\n", Aditya.Name)
}

type User struct{
	Name string
	email string
	Status bool
	Age int

}//User and everything is having first letter capital and hence can be used in other files means they are public