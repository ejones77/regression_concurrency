package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllFits(t *testing.T) {
	cols := []string{"A", "B", "C", "D"}
	x := FindAllFits(cols)
	assert.Equal(t, 15, len(x), "Expected 15 possible.")
}
