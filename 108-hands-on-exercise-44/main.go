package main

import "fmt"

func main() {
	slice := [][]int {
		{42, 43, 44, 45, 46},
		{47, 48, 49, 50, 51},
		{52, 53, 54, 55, 56},
		{57, 58, 59, 60, 61},
	}

	for _, v := range slice {
		fmt.Println(v)
	}
}