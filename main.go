package main

import (
	"fmt"
)

func check_integer(y_n int, y_err error, x_n int, x_err error) bool {
	if y_n != 1 || y_err != nil || x_n != 1 || x_err != nil {
		fmt.Print("\nEnter Valid Integers please!\n")
		return true
	}
	return false
}

func find_smallest(x int, y int) int {
	if x > y {
		return y
	}
	return x

}

func swap_small_great(x int, y int) [2]int {
	if x > y {
		arr := [2]int{x, y}
		return arr
	}
	arr := [2]int{y, x}
	return arr
}

func compute_highest_factor(x int, y int) int {
	var hcf int
	smallest := find_smallest(x, y)
	for i := 1; i <= smallest+1; i++ {
		if x%i == 0 && y%i == 0 {
			hcf = i
		}
	}
	return hcf
}

func hcf_euclidean(x int, y int) int {
	arr := swap_small_great(x, y)
	x, y = arr[0], arr[1]
	for y > 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	var x int
	var y int
	var result int
	var check_entries bool = true

	for check_entries {
		fmt.Print("Enter first number: ")
		x_n, x_err := fmt.Scan(&x)
		fmt.Print("Enter second number: ")
		y_n, y_err := fmt.Scan(&y)
		check_entries = check_integer(y_n, y_err, x_n, x_err)
	}

	result = compute_highest_factor(x, y)
	fmt.Print("\nUsing loops to find greatest factor: ", result)
	result = hcf_euclidean(x, y)
	fmt.Print("\nUsing Euclidian algorithm to find the greatest factor: ", result)
}
