package main

import (
	"add/calculator"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	var numbers []int

	for _, arg := range args {
		numberFromArg, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, numberFromArg)
	}

	result := calculator.Add(numbers...)

	fmt.Println("Calculating the sum...")
	fmt.Println("result: ", result)
}
