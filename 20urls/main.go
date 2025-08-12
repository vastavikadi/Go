package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://loc.dev:3000/learn?coursename=reactjs&paymentMode=paytm"

func main() {
	fmt.Println("THe url is: ", myUrl)

	//parsing the url means getting all the necessary details out of it
	result, _ := url.Parse(myUrl)
	fmt.Println(result)
	fmt.Println(result.Scheme) //https
	fmt.Println(result.Host)   //a lot of other methods - loc.dev:3000
	fmt.Println(result.ForceQuery)//false
	fmt.Println(result.Fragment)
	fmt.Println(result.OmitHost)//false
	fmt.Println(result.Opaque)
	fmt.Println(result.Path)// /learn
	fmt.Println(result.RawPath)
	fmt.Println(result.RawFragment)
	fmt.Println(result.RawQuery)// coursename=reactjs&paymentMode=paytm
	fmt.Println(result.User)
	fmt.Println(result.Port())// 3000

	qparams:= result.Query()// url.Values
	fmt.Printf("the type of the qparams is: %T\n", qparams)
	fmt.Println(result.Query())// map[coursename:[reactjs] paymentMode:[paytm]] using map for it and coursename is the key and the value is reactjs

	//constructing a url from the details (raw)
	partsOfUrl := &url.URL{
		Scheme:"https", 
		Host: "lco.dev",
		Path: "/learn",
		RawPath:"user=Aditya",
	}

	anotherNewUrl := partsOfUrl.String()
	fmt.Println("The url i created: ", anotherNewUrl)
}
