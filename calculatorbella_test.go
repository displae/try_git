package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModulo(t *testing.T) {
	result := Modulo(4, 2)
	assert.Equal(t, 0, result, "Fungsi Modulo Salah")
}
