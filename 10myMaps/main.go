package main

import "fmt"

func main() {
	languages := make(map[string]string) //[string] this is going to be the key and string is gonna be the value

	languages["JS"] = "Javascript" //ofcourse the key can be number and vice versa
	languages["TS"] = "Typescript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	fmt.Printf("The type of languages is: %T\n", languages) //map[string]string
	fmt.Println("List of all the languages: ", languages)   //map[JS:Javascript PY:Python RB:Ruby TS:Typescript]
	//to get the individual value
	fmt.Println("JS shorts for: ", languages["JS"])

	//to remove a key
	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	//to loop thru a map
	for key, value := range languages {
		fmt.Printf("For key %v, Value is %v\n", key, value) //%v is a placeholder for the value itself
	}
	for _, value := range languages {
		fmt.Printf("Value is %v\n", value) //%v is a placeholder for the value itself
	}//here we are ignoring the key


}
