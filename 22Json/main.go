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
}
