package main

import (
	"fmt"
	"strings"
)

func main() {

fmt.Println("Enter a word: ")
	
	var word string
	fmt.Scanln(&word)
	var myString string = word

	//myString := "iapapplen"

	// Option 1: (Recommended)
	if strings.HasPrefix(myString, "i") &&
strings.HasSuffix(myString, "n") &&
strings.Contains(myString, "a"){
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}

	
}
