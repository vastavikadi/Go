package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //now it is gonna be called as coursename and dont give any space in between, this is called aliases
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //means i dont want this field to be reflected to whoever is consumnig this json
	Tags     []string `json:"tags,omitempty"` //it removes the field if empty and no space in between
}

func main() {
	//going to encode the json means gonna convert to a valid json (from normal string to json)
	lcoCourses := []course{
		{"Reactjs Bootcamp", 299, "Lco.in", "123abc", []string{"webdev", "js"}},
		{"GOlang Course", 0, "Lco.in", "7847fr", []string{"webdev", "go"}},
		{"Web3", 59, "Lco.in", "abc", nil}}

	finalJson, err := json.Marshal(lcoCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson) //%s format specifier is used to print the uninterpreted bytes of a string or slice as a string. It expects a string value and outputs it directly

	//the output comes out to be lil messy: [{"Name":"Reactjs Bootcamp","Price":299,"Platform":"Lco.in","Password":"123abc","Tags":["webdev","js"]},{"Name":"GOlang Course","Price":0,"Platform":"Lco.in","Password":"7847fr","Tags":["webdev","go"]},{"Name":"Web3","Price":59,"Platform":"Lco.in","Password":"abc","Tags":null}] and notice null (where it was nil) and the password is visible which is not good when it comes to encoding

	finalJson2, err := json.MarshalIndent(lcoCourses, "", "\t") // try (lcoCourses, "Hey", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson2) // this has a prefix(basically adds something whatever u pass in it at every start of the line while json printing, indent takes \t here means adds an indentation of tab for each line)
	//this is the output from MarshalIndent
	// 		[
	//         {
	//                 "Name": "Reactjs Bootcamp",
	//                 "Price": 299,
	//                 "Platform": "Lco.in",
	//                 "Password": "123abc",
	//                 "Tags": [
	//                         "webdev",
	//                         "js"
	//                 ]
	//         },
	//         {
	//                 "Name": "GOlang Course",
	//                 "Price": 0,
	//                 "Platform": "Lco.in",
	//                 "Password": "7847fr",
	//                 "Tags": [
	//                         "webdev",
	//                         "go"
	//                 ]
	//         },
	//         {
	//                 "Name": "Web3",
	//                 "Price": 59,
	//                 "Platform": "Lco.in",
	//                 "Password": "abc",
	//                 "Tags": null
	//         }
	// ]//still the password and null persists

	//output from the aliases:
	// 	[
	//         {
	//                 "coursename": "Reactjs Bootcamp",
	//                 "Price": 299,
	//                 "website": "Lco.in",
	//                 "tags": [
	//                         "webdev",
	//                         "js"
	//                 ]
	//         },
	//         {
	//                 "coursename": "GOlang Course",
	//                 "Price": 0,
	//                 "website": "Lco.in",
	//                 "tags": [
	//                         "webdev",
	//                         "go"
	//                 ]
	//         },
	//         {
	//                 "coursename": "Web3",
	//                 "Price": 59,
	//                 "website": "Lco.in"
	//         }
	// ]

	//decoded json
	fmt.Println("DECODED JSON: -")
	DecodeJson() //this is the output: [{Reactjs Bootcamp 299 Lco.in  [webdev js]} {GOlang Course 0 Lco.in  [webdev go]} {Web3 59 Lco.in  []}]
	fmt.Println("STRUCT JSON: -")
	makeStructFromJSON()
}

func DecodeJson() {
	jsonData := []byte(`
	[
        {
                "coursename": "Reactjs Bootcamp",
                "Price": 299,
                "website": "Lco.in",
                "tags": [
                        "webdev",
                        "js"
                ]
        },
        {
                "coursename": "GOlang Course",
                "Price": 0,
                "website": "Lco.in",
                "tags": [
                        "webdev",
                        "go"
                ]
        },
        {
                "coursename": "Web3",
                "Price": 59,
                "website": "Lco.in"
        }
]`)

	var lcoCourses []course

	checkValid := json.Valid(jsonData) //to check the validity(means the structure of the json is correct)
	fmt.Println(checkValid)            // returns true but as soon as ek comma bhi aaya then its wrong

	if checkValid {
		fmt.Println("JSON WAS VALID")
		json.Unmarshal(jsonData, &lcoCourses) //Unmarshal parses the JSON-encoded data and passing the &lcoCourses because it stores in the original

		fmt.Printf("%#v", lcoCourses) //[]main.course{main.course{Name:"Reactjs Bootcamp", Price:299, Platform:"Lco.in", Password:"", Tags:[]string{"webdev", "js"}}, main.course{Name:"GOlang Course", Price:0, Platform:"Lco.in", Password:"", Tags:[]string{"webdev", "go"}}, main.course{Name:"Web3", Price:59, Platform:"Lco.in", Password:"", Tags:[]string(nil)}}
		// fmt.Println(lcoCourses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}
}

func makeStructFromJSON() {
	// some cases where you just want to add data to key value
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	// map[string]interface{} only works if the root JSON value is an object starts with {} as jsonDataFromWeb not like jsonData
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Println("STRUCT FROM JSON: -")
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}
	//used when
	// 	Don’t know the JSON structure ahead of time (or don’t want to create a Go struct for it).
	// 	Still want to access JSON fields dynamically in Go.
}


//NOTE for : Unmarshal
// Here’s what’s happening under the hood:
// map[string]interface{} means:
// Keys are always strings (just like JSON object keys).
// Values are stored as interface{} so they can hold any JSON type — strings, numbers, booleans, arrays, nested maps, etc.

// When json.Unmarshal reads JSON into this type:
// JSON objects ({ ... }) → become map[string]interface{}
// JSON arrays ([ ... ]) → become []interface{}
// JSON numbers → become float64 by default
// JSON strings → become string
// JSON booleans → become bool
// JSON null → becomes nil