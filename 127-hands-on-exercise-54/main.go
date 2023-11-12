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
		lastName: "Chn",
		favoriteIceCreamFlavors: []string{"Ice", "Coke"},
	}

	m := map[string]person {
		p1.lastName: p1,
		p2.lastName: p2, 
	}

	fmt.Println(m[p1.lastName])
	fmt.Println(m[p2.lastName])

	for i, v := range m["Chen"].favoriteIceCreamFlavors {
		fmt.Printf("%#v - %#v\n", i, v)
	}

	for i, v := range m["Chn"].favoriteIceCreamFlavors {
		fmt.Printf("%#v - %#v\n", i, v)
	}

	m2 := make(map[string]person)
	m2[p1.lastName] = p1
	m2[p2.lastName] = p2

	for _, v := range m2 {
		fmt.Println(v)
		for _, v2 := range v.favoriteIceCreamFlavors {
			fmt.Println(v2)
		}
	}

}