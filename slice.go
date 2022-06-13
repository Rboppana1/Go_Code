package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var test_num string
	testSlice := make([]int, 3)
	for true {
		fmt.Printf("\nPlease enter an integer('X' to Exit):")
		fmt.Scan(&test_num)
		if test_num == "X" {
			break
		}
		i, err := strconv.Atoi(test_num)
		if err == nil {
		}
		testSlice = append(testSlice, i)
		sort.Sort(sort.IntSlice(testSlice))
		fmt.Printf(fmt.Sprint(testSlice))
	}

}