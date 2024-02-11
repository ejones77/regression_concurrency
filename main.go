package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-gota/gota/dataframe"
)

type DF = dataframe.DataFrame

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
	for i := 0; i < 100; i++ {
		FitSequential(df, allModels, trainDf, testDf)
	}
	elapsed := time.Since(start)
	fmt.Printf("FitSequential took %s\n", elapsed)

	start = time.Now()
	for i := 0; i < 100; i++ {
		FitConcurrent(df, allModels, trainDf, testDf)
	}
	elapsed = time.Since(start)
	fmt.Printf("FitConcurrent took %s\n", elapsed)
}
