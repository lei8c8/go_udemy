package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45}

	for i, v := range slice {
		fmt.Printf("index = %v, value = %v\n", i, v)
	}
}