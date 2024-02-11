package main

import (
	"log"
	"sync"

	"gonum.org/v1/gonum/stat"
)

func FitModel(df DF, cols []string) ([]float64, error) {
	if len(cols) == 1 {
		// Perform simple linear regression
		target := df.Col("mv").Float()
		feature := df.Col(cols[0]).Float()
		alpha, beta := stat.LinearRegression(feature, target, nil, false)
		return []float64{alpha, beta}, nil
	} else {
		// Perform multiple linear regression
		return MultipleLinearRegression(df, "mv", cols)
	}
}

func FitSequential(df DF, allModels [][]string, trainDf DF, testDf DF) {
	regressions := make([]RegressionResult, len(allModels))
	for i, model := range allModels {
		coefficients, err := FitModel(trainDf, model)
		if err != nil {
			log.Fatal(err)
		}
		regressions[i] = RegressionResult{Coefficients: coefficients, Model: model}
	}

	processResults(regressions, allModels, testDf, "outputs/output-sequential.txt")
}

func FitConcurrent(df DF, allModels [][]string, trainDf DF, testDf DF) {
	results := make([]RegressionResult, len(allModels))
	// helps ensure the goroutines are done before trying to compute results
	var wg sync.WaitGroup

	for i, model := range allModels {
		wg.Add(1)
		go func(i int, model []string) {
			defer wg.Done()
			coefficients, err := FitModel(trainDf, model)
			if err != nil {
				log.Fatal(err)
			}
			results[i] = RegressionResult{Coefficients: coefficients, Model: model}
		}(i, model)
	}

	wg.Wait()
	processResults(results, allModels, testDf, "outputs/output-concurrent.txt")
}
