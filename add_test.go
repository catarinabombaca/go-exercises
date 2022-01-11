package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("given 2 integers 4 and 5, when we add them, then return 9 as a result", func(t *testing.T) {
		expectedValue := 9
		firstInt := 5
		secondInt := 4

		actualValue := Add(firstInt, secondInt)

		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("given a variable number of integers, when we add them, then return the total sum as a result", func(t *testing.T) {
		var (
			number1 = 5
			number2 = 4
			number3 = 3
			number4 = 2
			number5 = -10
			expectedResult = 4
		)

		actualResult := Add(number1, number2, number3, number4, number5)

		assert.Equal(t, expectedResult, actualResult)

	})
}
