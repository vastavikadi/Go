//functions and methods are same but mwthods are functions defined for a class
//we do not have classes in go but we do have structs
//functions inside structs are called methods

package main

import "fmt"

func main() {
	Aditya := User{"Aditya", "vastavikadi@gmail.com", 14, "Shukla"}
	Aditya.GetEmail()
	Aditya.NewMail()                                        //adi@gmail.com
	fmt.Println("this is the user's email: ", Aditya.Email) //vastavikadi@gmail.com means it does not changes the original value means a copy was passed hence pointers needed now

}

type User struct {
	Name     string
	Email    string
	Age      int
	Lastname string
}

func (u User) GetEmail() { //we are passing the struct User as u and GetEMail is capital and hence can be public and imported in other files
	fmt.Println("this is the user's email: ", u.Email)
}

// func (u User) NewMail() {//this is just passing the copy not the original
// 	emailPtr := &u.Email
// 	*emailPtr = "adi@gmail.com"
// 	fmt.Println("this is the user's email: ", u.Email)
// }
func (u *User) NewMail() {//here we are passing the original User
	emailPtr := &u.Email
	*emailPtr = "adi@gmail.com"
	fmt.Println("this is the user's email: ", u.Email)
}
