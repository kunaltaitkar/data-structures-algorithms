package main

import "fmt"

func main() {
	input := [5]int{1, 2, 3, 4, 5}

	start := 0
	end := len(input) - 1

	for start < end {
		temp := input[start]
		input[start] = input[end]
		input[end] = temp
		start = start + 1
		end = end - 1
	}

	fmt.Println(input)
}
