package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(1, 5)
	assert.Equal(t, 6, result, "Fungsi penjumlahan error")
}
