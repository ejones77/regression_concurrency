package main

import (
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/mat"
)

func MultipleLinearRegression(df dataframe.DataFrame, target string, features []string) ([]float64, error) {
	// We need to transform the features to a matrix
	X := mat.NewDense(df.Nrow(), len(features)+1, nil)
	for i, feature := range features {
		floats := df.Col(feature).Float()
		for j, val := range floats {
			X.Set(j, i, val)
		}
	}

	// Add a column of ones for the intercept
	for i := 0; i < df.Nrow(); i++ {
		X.Set(i, len(features), 1)
	}

	y := mat.NewVecDense(df.Nrow(), df.Col(target).Float())

	// using the normal equation (X'X^-1)X'y
	var XTX, XTy, coef mat.Dense
	XTX.Mul(X.T(), X)
	XTy.Mul(X.T(), y)
	err := coef.Solve(&XTX, &XTy)
	if err != nil {
		return nil, err
	}

	return coef.RawMatrix().Data, nil
}
