package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	//given two integers 5 and 4
	expectedValue := 9
	firstInt := 5
	secondInt := 4

	//when we call Add and pass them as arguments
	actualValue := Add(firstInt, secondInt)

	//should return 9
	assert.Equal(t, expectedValue, actualValue)
}
