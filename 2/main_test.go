package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLineParsing(t *testing.T) {
	assert := assert.New(t)
	line := "48 51 52 53 52"
	result := []int{48, 51, 52, 53, 52}
	parsedLine := parseLine(line)
	assert.Equal(result, parsedLine)
}

func TestIsLineValid(t *testing.T) {
	assert := assert.New(t)
	invalidData := []int{48, 51, 52, 53, 52}
	alsoInvalidData := []int{10, 6, 5, 4}
	anotherInvalidData := []int{1, 3, 2, 4, 5}
	validData := []int{7, 6, 5, 4}
	alsoValidData := []int{7, 9, 10, 12}

	assert.Equal(validateData(invalidData), false)
	assert.Equal(validateData(alsoInvalidData), false)
	assert.Equal(validateData(anotherInvalidData), false)
	assert.Equal(validateData(validData), true)
	assert.Equal(validateData(alsoValidData), true)
	assert.Equal(validateData([]int{94, 92, 89, 86, 83, 80}), true)
}

func TestDeleteFromSlice(t *testing.T) {
	assert := assert.New(t)
	slice := []int{1, 5, 20, 3, 8}
	assert.Equal(deleteFromSlice(slice, 0), []int{5, 20, 3, 8})
	assert.Equal(deleteFromSlice(slice, 1), []int{1, 20, 3, 8})
	assert.Equal(deleteFromSlice(slice, 2), []int{1, 5, 3, 8})
	assert.Equal(deleteFromSlice(slice, 3), []int{1, 5, 20, 8})
	assert.Equal(deleteFromSlice(slice, 4), []int{1, 5, 20, 3})
}
