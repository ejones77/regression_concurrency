package main

import (
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/mat"
)

// MultipleLinearRegression performs multiple linear regression on the given dataframe.
// It returns the coefficients of the linear model.
func MultipleLinearRegression(df dataframe.DataFrame, target string, features []string) ([]float64, error) {
	// Create a matrix of features
	X := mat.NewDense(df.Nrow(), len(features)+1, nil)
	for i, feature := range features {
		floats := df.Col(feature).Float()
		for j, val := range floats {
			X.Set(j, i, val)
		}
	}

	// Add a column of ones for the intercept term
	for i := 0; i < df.Nrow(); i++ {
		X.Set(i, len(features), 1)
	}

	// Create a vector of target values
	y := mat.NewVecDense(df.Nrow(), df.Col(target).Float())

	// Calculate the coefficients using the normal equation: (X'X)^-1 X'y
	var XTX, XTy, coef mat.Dense
	XTX.Mul(X.T(), X)
	XTy.Mul(X.T(), y)
	err := coef.Solve(&XTX, &XTy)
	if err != nil {
		return nil, err
	}

	return coef.RawMatrix().Data, nil
}
