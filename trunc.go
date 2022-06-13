
package main

import "fmt"

func main() {

	fmt.Println("Enter a floating1 point number: ")
	
	var floating1 float32
	fmt.Scanln(&floating1)
	var x float32 = floating1
	var y int = int(x)
	fmt.Println(y)
	
}
