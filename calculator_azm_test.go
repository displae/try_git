package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstract(t *testing.T) {
	result := Substract(5, 1)
	assert.Equal(t, 4, result, "Fungsi pengurangan error")
}

func Substract(num1 int, num2 int) int {
	return num1 - num2
}
