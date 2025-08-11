package main

import "fmt"

func main() {
	//pointers exist to confirm that whatever you are passing is real and not copy of that thing

	//basically there are two ways
	//pass by value and pass by reference

	// 	Pass by Value
	// You give the function a copy of your data.
	// The function works on that copy.
	// The original variable stays unchanged (unless it’s a reference type internally).

	// 	Pass by Reference
	// You give the function a direct way (like a pointer, reference, or memory address) to access your original variable.
	// The function works on the same data in memory.
	// Any change is visible outside the function.

	// In Go, everything is pass-by-value, but you can pass a pointer to simulate pass-by-reference behavior.


	//Conclusions
	//pointers exist to let us pass by reference to work with the original variable

	var name string= "aditya"
	fmt.Println("My name is: ", name);

	// var ptr *string ( this is a pointer, can be named anything rather than ptr as well and its type is string and we use asterick sign to define them as i did here)
	//another symbol is '&'

	var marks *int//<nil> if no value assigned
	fmt.Println("My pointer is: ", marks)

	var namePtr = &name//namePtr is referencing to name (& means reference) pointer reference to driect memory address
	fmt.Println("My pointer is: ", namePtr)//0yc000334a0a2 this is the memory address where name is stored
	fmt.Println("My pointer is: ", *namePtr)//gives the value stored in this pointer

	var myMarks int = 35
	var marksPtr = &myMarks
	fmt.Println("marksPtr: ", marksPtr)
	//some operations
	*marksPtr = *marksPtr * 2//it means *myMarks = 35 and hence on 35 x 2 = 70
	fmt.Println("New myMarks: ", myMarks)//70 original value changed
	fmt.Println("marksPtr: ", marksPtr)
	fmt.Println("New marksPtr: ", *marksPtr)//70
	//now you can see the actual value got changed
}
