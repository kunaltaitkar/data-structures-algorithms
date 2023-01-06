package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5}
	current := input[0]
	for i := 0; i < len(input); i++ {
		if i == len(input)-1 {
			input[0] = current
			break
		}
		next := input[i+1]
		input[i+1] = current
		current = next
	}
	fmt.Println(input)
}
