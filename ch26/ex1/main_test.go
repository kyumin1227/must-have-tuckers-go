package main

import (
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

func TestSquare1(t *testing.T) {
	assert := assert2.New(t)
	assert.Equal(81, square(9), "square(9) should be 81")
}

func TestSquare2(t *testing.T) {
	assert := assert2.New(t)
	assert.Equal(9, square(3), "square(3) should be 9")
}
