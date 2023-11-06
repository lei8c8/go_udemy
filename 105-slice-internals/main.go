package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []float64{3, 1, 4, 2}
	fmt.Println(medianOne(n))

}

func medianOne(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2
	if len(x) % 2 == 1 {
		return x[i / 2]
	}
	return (x[i - 1] + x[i]) / 2

}