package main

import (
	"fmt"
	"sort"
)

func average(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

func mean_and_median(numbers []int) (interface{}, interface{}) {
	length := len(numbers)
	if length == 0 {
		return nil, nil
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	mean := float64(sum) / float64(len(numbers))

	median := float64(numbers[0])

	sort.Ints(numbers)
	if length%2 == 0 {
		mid := length / 2
		median = float64(numbers[mid-1]+numbers[mid]) / 2.0
	} else {
		mid := (length + 1) / 2
		median = float64(numbers[mid-1])
	}

	return mean, median
}

func main() {
	a := []int{1, 4, 5, 8, 0, 7}
	avg := average(a)
	fmt.Println("average =", avg)

	// returns multiple values
	mean, median := mean_and_median(a)
	fmt.Println("mean =", mean, "median =", median)
}
