package main

import (
	"fmt"
	"math/rand"
)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range(xi) {
		fmt.Printf("Index is %v, value is %v\n", i, v)
	}

	m := map[string]int {
		"James": 42,
		"Moneypenny": 32,
	}

	for k, v := range(m) {
		fmt.Printf("Key is %v, value is %v\n", k, v)
	}

	age := m["James"] 
	fmt.Println(age)

	age = m["Q"]
	fmt.Println(age)

	if v, ok := m["Q"]; !ok {
		fmt.Printf("there is no Q, the default value is %v\n", v)
	}

	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("x is 3")
		}
	}
}