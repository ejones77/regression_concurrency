package main

import (
	"fmt"
	"testing"

	"github.com/go-gota/gota/dataframe"
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

func TestMultipleLinearRegression(t *testing.T) {
	df := dataframe.LoadRecords(
		[][]string{
			{"A", "B", "C"},
			{"1", "2", "4"},
			{"2", "3", "7"},
			{"7", "3", "3"},
		},
	)

	target := "C"
	features := []string{"A", "B"}

	coef, err := MultipleLinearRegression(df, target, features)
	if err != nil {
		t.Fatalf("Expected no error in MultipleLinearRegression, but got: %v", err)
	}
	fmt.Println(coef)

	expected := []float64{-0.8, 3.8, -2.8}
	assert.InEpsilonSlice(t, expected, coef, 1e6)
}

func TestFitModel(t *testing.T) {
	df := dataframe.LoadRecords(
		[][]string{
			{"mv", "A", "B"},
			{"4", "1", "2"},
			{"7", "2", "3"},
			{"3", "7", "3"},
		},
	)

	// Test simple regression
	cols := []string{"A"}
	coef, err := FitModel(df, cols)
	if err != nil {
		t.Fatalf("Expected no error in FitModel, but got: %v", err)
	}
	expected := []float64{0, 1}
	assert.InEpsilonSlice(t, expected, coef, 1e6)

	// Test multiple regression
	cols = []string{"A", "B"}
	coef, err = FitModel(df, cols)
	if err != nil {
		t.Fatalf("Expected no error in FitModel, but got: %v", err)
	}
	expected = []float64{-0.8, 3.8, -2.8}
	assert.InEpsilonSlice(t, expected, coef, 1e6)
}
