package main

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "lei",
		friends: map[string]int{
			"abc": 4,
		},
		favDrinks: []string{"Coke", "Beer"},
	}

	fmt.Println(s)
}
