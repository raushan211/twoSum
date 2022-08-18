package main

import "fmt"

func main() {
	input := []int{2, 7, 11, 15}

	// target = 9

	for i := 0; i < len(input); i++ {
		for j := 1; j < len(input); j++ {
			if input[i]+input[j] == 9 && i != j {
				fmt.Println(i, j)

			}
		}
	}
}
