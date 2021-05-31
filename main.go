package main

import "fmt"

func main() {
	input := [][]int{{25, 30}, {2, 19}, {14, 23}, {4, 8}}
	fmt.Println(fmt.Sprintf("Input:  %v", input))
	output, e := Merge(input)
	if e != nil {
		panic(e)
	}
	fmt.Println(fmt.Sprintf("Output: %v", output))
}
