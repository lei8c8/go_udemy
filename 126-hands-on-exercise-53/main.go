package main

import "fmt"

type person struct {
	firstName string
	lastName string
	favoriteIceCreamFlavors []string
}

func main() {

	p1 := person {
		firstName: "Tony",
		lastName: "Chen",
		favoriteIceCreamFlavors: []string{"Ice", "Cream"},
	}

	p2 := person {
		firstName: "Tim",
		lastName: "Chen",
		favoriteIceCreamFlavors: []string{"Ice", "Coke"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for _, f := range p1.favoriteIceCreamFlavors {
		fmt.Println(f)
	}

	for _, f := range p2.favoriteIceCreamFlavors {
		fmt.Println(f)
	}

}