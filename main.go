package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-gota/gota/dataframe"
)

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

	start := time.Now()
	FitSequential(df, allModels, trainDf, testDf)
	elapsed := time.Since(start)
	fmt.Printf("FitSequential took %s\n", elapsed)

	start = time.Now()
	FitConcurrent(df, allModels, trainDf, testDf)
	elapsed = time.Since(start)
	fmt.Printf("FitConcurrent took %s\n", elapsed)
}
