package main

import "fmt"

func main() {
	s := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "I'm 008"},
	}

	for _, v := range s {
		for _, word := range v {
			fmt.Println(word)
		}
	}
}