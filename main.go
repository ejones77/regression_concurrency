package main

import (
	"fmt"

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

func main() {
	//file, _ := os.Open("given_files/boston.csv")
	cols := []string{"A", "B", "C", "D"}
	x := FindAllFits(cols)
	fmt.Println(x)
	fmt.Println(len(x))
}
