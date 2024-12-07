package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMulValue(t *testing.T) {
	assert := assert.New(t)
	result, do := getMulValue("mul(2,4)", true)
	assert.Equal(result, 8)
	result, _ = getMulValue("xmul(5,3)%&mul(3,7)!@^d", true)
	assert.Equal(result, 36)
	result, _ = getMulValue("+mul(32,64]then(mul(11,8)mul(8,5))", true)
	assert.Equal(result, 128)
	result, do = getMulValue("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5)don't()mul(2,3))", true)
	assert.Equal(result, 48)
	assert.Equal(do, false)
}
