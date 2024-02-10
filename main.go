package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/combin"
)

// Calculate possible linear fits
func FindAllFits(cols []string) [][]string {
	// "Combinations" returns an index here
	// We match that with the position in "cols"
	var allModels [][]string
	n := len(cols)
	for k := 1; k <= n; k++ {
		options := combin.Combinations(n, k)
		for _, option := range options {
			models := make([]string, len(option))
			for i, idx := range option {
				models[i] = cols[idx]
			}
			allModels = append(allModels, models)
		}
	}
	return allModels
}

func meanSquaredError(actual, predicted []float64) float64 {
	n := len(actual)
	// needs to return a float so the total starts that way too
	total := 0.0
	for i := 0; i < n; i++ {
		diff := predicted[i] - actual[i]
		total += diff * diff
	}
	return total / float64(n)
}

type DF = dataframe.DataFrame

func trainTestSplit(df DF, trainSize float64) (DF, DF) {
	if trainSize < 0 || trainSize > 1 {
		log.Fatalf("trainSize has to be between 0 and 1")
	}
	r := rand.New(rand.NewSource(42))
	perm := r.Perm(df.Nrow())
	df = df.Subset(perm)

	trainNum := int(trainSize * float64(df.Nrow()))

	trainDf := df.Subset(perm[:trainNum])
	testDf := df.Subset(perm[trainNum:])

	return trainDf, testDf
}

func FitModel(df dataframe.DataFrame, cols []string) ([]float64, error) {
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

func main() {

	file, err := os.Open("given_files/boston.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	df := dataframe.ReadCSV(file)
	trainDf, testDf := trainTestSplit(df, 0.8)

	cols := df.Names()
	var newCols []string
	for _, col := range cols {
		if col != "mv" && col != "neighborhood" {
			newCols = append(newCols, col)
		}
	}
	allModels := FindAllFits(newCols)
	fmt.Printf("\n num models: %v", len(allModels))

	type RegressionResult struct {
		Coefficients []float64
	}

	regressions := make([]RegressionResult, len(allModels))
	for i, model := range allModels {
		coefficients, err := FitModel(trainDf, model)
		if err != nil {
			log.Fatal(err)
		}
		regressions[i] = RegressionResult{Coefficients: coefficients}
	}

	output, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	for i, model := range allModels {
		result := regressions[i]

		predicted := make([]float64, testDf.Nrow())
		for j, col := range model {
			feature := testDf.Col(col).Float()
			for i, val := range feature {
				predicted[i] += result.Coefficients[j] * val
			}
		}

		actual := testDf.Col("mv").Float()
		mse := meanSquaredError(actual, predicted)
		fmt.Fprintf(output, "Features: %v\n", model)
		fmt.Fprintf(output, "MSE: %v\n", mse)
	}

}
