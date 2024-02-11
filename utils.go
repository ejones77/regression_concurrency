package main

import (
	"log"
	"math/rand"

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
