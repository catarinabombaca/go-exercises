package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	expectedValue := 9
	firstInt := 5
	secondInt := 4

	actualValue := Add(firstInt, secondInt)

	assert.Equal(t, expectedValue, actualValue)
}
