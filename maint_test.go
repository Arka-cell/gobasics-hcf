package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestLoopHCF(t *testing.T) {
	x := 30
	y := 12
	var expected int = 6
	result := compute_highest_factor(x, y)
	if result != expected {
		response := "Assertion error! Expected" + strconv.Itoa(expected) + "instead of " + strconv.Itoa(result)
		t.Fatalf(response)
	}
}

func TestEuclideanHCF(t *testing.T) {
	x := 30
	y := 12
	var expected int = 6
	result := hcf_euclidean(x, y)
	if result != expected {
		response := "Assertion error! Expected" + strconv.Itoa(expected) + "instead of " + strconv.Itoa(result)
		t.Fatalf(response)
	}
}

func TestInput(t *testing.T) {
	var y int
	var x int
	os.Args = []string{"30", "12"} // os.Args accepts only list of strings.
	y_n, y_err := fmt.Scan(&y)
	x_n, x_err := fmt.Scan(&x)
	if !check_integer(y_n, y_err, x_n, x_err) {
		response := "Assertion error! Expected Non-valid integers that were confirmed as integers"
		t.Fatalf(response)
	}
}
