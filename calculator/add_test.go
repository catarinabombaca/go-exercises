package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	var (
		assert = assert.New(t)
		testNums = []int{5,4,3,2,-10}
	)

	t.Run("given 2 integers 4 and 5, when we add them, then return 9 as a result", func(t *testing.T) {
		expectedValue := 9
		num1 := testNums[1]
		num2 := testNums[0]

		actualValue := Add(num1, num2)

		assert.Equal(expectedValue, actualValue)
	})

	t.Run("given a variable number of integers, when we add them, then return the total sum as a result", func(t *testing.T) {
		expectedResult := 4
		actualResult := Add(testNums...)

		assert.Equal(expectedResult, actualResult)
	})
}