package main

import (
	"encoding/json"
	"fmt"
)


func main() {

	var personName string
	var adrs string

	fmt.Println("Enter a name")

	fmt.Scan(&personName)

	fmt.Println("Enter the address ")

	fmt.Scan(&adrs) 

	var userTest = map[string]string{"fname": personName, "address": adrs} 

	userTestJSON, err := json.Marshal(userTest) 

	if err != nil {
		fmt.Println("error whie marshalling userTest map")
	}

	fmt.Println("Jason: ", string(userTestJSON)) 

}
