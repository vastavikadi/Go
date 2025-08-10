package main

import (
	"fmt"
	"time"
)

func main() {
	//to get the current time
	presentTime := time.Now()//can have even more methods further time.Now().Nanosecond()
	fmt.Println(presentTime) //2025-08-10 20:10:57.8700003 +0530 IST m=+0.000547601 this is output

	fmt.Println(presentTime.Format("01-02-2006"))        //here 01-02-2006 will always be the same otherwise it will not run
	fmt.Println(presentTime.Format("01-02-2006 Monday")) //always Monday for the day to be printed correctly, if u dont write Monday and write Sunday no error but gives the wrong output
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//to create the time manually yourself
	createdDate := time.Date(2005, time.August, 26, 4, 0, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}
