package helpers

import "strconv"

func ConvertToSliceOfInts(args []string) []int {
	var numbers []int

	for _, arg := range args {
		numberFromArg, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, numberFromArg)
	}

	return numbers
}