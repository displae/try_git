package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	result := multiply(1, 5)
	assert.Equal(t, 5, result, "Fungsi pengalian error")
}
