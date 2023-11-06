package main

import "fmt"

func print_dot() {
	fmt.Println("------------")
}

func main() {
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	print_dot()

	xi = append(xi, 45, 46, 47, 99, 777)
	fmt.Println(xi)

	xi = append(xi, []int{1, 2, 3}...)
	fmt.Println(xi)
	
	print_dot()
}