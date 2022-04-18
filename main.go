package main

import (
	"fmt"
)

func main() {
	fmt.Println("Calculating the sum...")
}

func Add(nums ...int) (result int) {
	for _, num := range nums {
		result+= num
	}

	return
}
