package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

type RegressionResult struct {
	Coefficients []float64
	Model        []string
}

func processResults(results []RegressionResult, allModels [][]string, testDf DF, filename string) {
	output, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	for i, result := range results {
		model := allModels[i]
		predicted := make([]float64, testDf.Nrow())
		// computes predicted values using the associated model
		for j, col := range model {
			feature := testDf.Col(col).Float()
			for i, val := range feature {
				predicted[i] += result.Coefficients[j] * val
			}
		}

		actual := testDf.Col("mv").Float()
		mse := meanSquaredError(actual, predicted)

		n := float64(len(actual))
		k := float64(len(result.Coefficients))
		// using mse to estimate variane
		// https://robjhyndman.com/hyndsight/lm_aic.html#:~:text=Since%20we%20don't%20know,nlog(2%CF%80
		AIC := 2*k + n*math.Log(mse)
		fmt.Fprintf(output, "AIC: %v\n", AIC)
		fmt.Fprintf(output, "Features: %v\n", model)
		fmt.Fprintf(output, "MSE: %v\n", mse)
	}
}
