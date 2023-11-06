package main

import "fmt"

func main() {
	array := [5]int{}

	for i := 0; i < len(array); i++ {
		array[i] = i
	}

	for i, v := range array {
		fmt.Printf("%v - %T - %v\n", v, v, i)
	}
}