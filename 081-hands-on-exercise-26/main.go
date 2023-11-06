package main

import "fmt"

func main() {
	var x int = 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
	xi := []int{42, 43, 44, 45, 46, 47}
	fmt.Println(xi)
}