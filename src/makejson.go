package main

import (
	"encoding/json"
	"fmt"
)

/**
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create
a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should
use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/
func main() {
	var  name string
	var address string
	personMap := make(map[string]string)

	fmt.Printf("Enter your name\n")
	fmt.Scan(&name)
	personMap["name"] = name

	fmt.Printf("Enter your address\n")
	fmt.Scan(&address)
	personMap["address"] = address
	barr, _ := json.Marshal(personMap)
	fmt.Println(string(barr))
}
