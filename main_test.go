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

func TestMeanSquaredError(t *testing.T) {
	actual := []float64{1.0, 2.0, 3.0}
	predicted := []float64{1.0, 2.5, 3.5}

	mse := meanSquaredError(actual, predicted)
	expected := 0.16666667

	epsilon := 1e-7
	assert.InEpsilon(t, expected, mse, epsilon)
}
