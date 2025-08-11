package main

import "fmt"

func main() {
	fmt.Println("welcome to if and else conditional statements/control flow statements")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {//the identation should be exactly this otherwise it will throw an error
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	//can create the variables on the go as well need not to be defined in advance

	if 9%2==0{
		fmt.Println("the number is even")
	}else{
		fmt.Println("the number is odd")
	}

	if num:=3; num<10{//defining and using in the same line 
		fmt.Println("the number is less than 10")
	}else{
		fmt.Println("the number is not less than 10")
	}
}
