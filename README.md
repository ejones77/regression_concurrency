# regression_concurrency
Using Go to compare sequential and concurrent execution of linear regression models.

## Overview

Submitted assignment for Northwestern MSDS 431

This repository performs simple and multiple linear regression on all feature combinations from a provided file of Boston housing data.

## Setup

To run these tests locally:
- Clone the repository
- From within the repo's directory, run `go build .`
- Then run `./regression_concurrency.exe`

## About

`main.go` calculates the time it takes to run two different methods of fitting the linear models, one with sequential processing, and one with concurrent processing. The data is read into a dataframe from the gota library.

The sequential and concurrent versions each run 100 times, and the time each version takes is printed to the terminal.

The functionality is split up across various files:

- `utils.go` includes three helper functions 
    - `FindAllFits` takes in a slice of strings and finds all possible combinations of those strings using gonum/stat/combin
    - `meanSquaredError` uses actual and predicted values as float slices to calculate MSE
    - `trainTestSplit` takes in a gota dataframe and splits the dataframe into two under a reproducible random seed. 80% of the data is used to train while 20% is used to test.

- `multiple_regression.go` uses gonum's matrix library to convert the dataframe into a matrix and segment the target vector from the feature matrix so that the normal equation can calculate the coefficients of a linear model.

- `fit_model.go` has three functions: `FitModel` which is applied by both `FitSequential` and `FitSequential`
    - `FitModel` first considers whether there is one feature, so that simple linear regression should be performed, or if there's multiple features, and multiple regression is needed. 
    - `FitSequential` loops through all model combinations one at a time in a simple for loop, then prints the results to a `outputs/output-sequential.txt`
    - `FitConcurrent` uses goroutines to call `FitModel` and store results while calling a `Sync.WaitGroup` to ensure all the goroutines are done before moving the results to `processResults`

- `process_results.go` handles the application of the MSE and AIC calculations. The MSE is also used to calculate the AIC with a formula [described here](https://robjhyndman.com/hyndsight/lm_aic.html#:~:text=Since%20we%20don't%20know)


## Results

My results after compiling and running locally:
`Sequential`: 35.33 seconds
`Concurrent`: 18.36 seconds

That's nearly twice as fast! Concurrent programming can open a lot of doors with respect to performance. I found that goroutines are easier to work with compared to Python's threading and multiprocessing libraries. To expand on the experiment, this framework could be applied to existing linear regression models as a quick check to see if the best feature combinations have been selected when making predictions. 