package main

import (
	"add/calculator"
	"add/helpers"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	numbers := helpers.ConvertToSliceOfInts(args)

	result := calculator.Add(numbers...)

	fmt.Println("Calculating the sum...")
	fmt.Println(result)
}
