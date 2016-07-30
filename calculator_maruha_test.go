package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	result := Divide(6, 3)
	assert.Equal(t, 2, result, "Fungsi pembagian salah.")
}
