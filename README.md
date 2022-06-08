# Golang Basics:
The following takes users input of two numbers and finds the greatest common factors between the numbers.
The programs obliges the user to enter valid integers only. And it has two algorithms:
1. Simple iteration algorithm:
    Assuming that the inputs are denoted as `x` and `y` It will loop between the range of 1 and the smallest number from the two input while generating an index denominated as `i` and saves it if both `x % i` and `y % i` are both equal to 0.
2. Simple iteration algorithm:
    Assuming that `y > 0`, we will divide the greater by smaller and take the remainder. Then, divide the smaller by this remainder and repeat the process until the remainder is 0.

## How to run:
1. Open your command line and run `go run main.go` to run your program.
2. Open your command line and run `go test` to run your unit tests.

## What has been tested:
1. A function that will check if an input from `fmt.Scan` is of type `int` based on the result of `fmt.Scan`
2. 