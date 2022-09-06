/*
Write a program to find different aggregates (min, max, mean and median) value in
existing list of numbers. [5, 10, 15, 12, 500, 950, 0, 0.5, 6.75, 840, 1500, 7, 84, 15000]
*/

package main

import (
	"fmt"
	"sort"
)

func median(data []float64) float64 {
	dataCopy := make([]float64, len(data))
	copy(dataCopy, data)

	sort.Float64s(dataCopy)

	var median float64
	l := len(dataCopy)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
		// fmt.Printf("Median of the array is :")
	}

	return median
}

func mean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	var sum float64
	for _, d := range data {
		sum = sum + d
	}
	return sum / float64(len(data))

}

func minValue(data []float64) float64 {
	min := data[0]
	for _, v := range data {
		if v < min {
			min = v
		}
	}
	return min

}

func maxValue(data []float64) float64 {
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max

}

func main() {
	fmt.Println("MEDIAN OF THE ARRAY IS :", median([]float64{5, 10, 15, 12, 500, 950, 0.5, 6.75, 840, 1500, 7, 84, 15000}))
	fmt.Println("MEAN OF THE ARRAY IS :", mean([]float64{5, 10, 15, 12, 500, 950, 0.5, 6.75, 840, 1500, 7, 84, 15000}))
	fmt.Println("Minimum value is from array is : ", minValue([]float64{5, 10, 15, 12, 500, 950, 1, 6.75, 840, 1500, 7, 84, 15000}))
	fmt.Println("Minimum value is from array is : ", maxValue([]float64{5, 10, 15, 12, 500, 950, 1, 6.75, 840, 1500, 7, 84, 15000}))
	// fmt.Println(median([]float64{3, 3, 3, 3, 2, 22, 2, 2, 2, 2, 1, 1, 1, 1, 1, 111}))
}
