package main

import "fmt"

func main() {
	si := []string{"a", "b", "c"}
	fmt.Println(si)

	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-----------")

	xi = append(xi, 10, 11, 12, 13)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("-----------")

	jd := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}

	fmt.Println(jd)
	fmt.Println(jm)

	xxs := [][]string{jd, jm}
	fmt.Println(xxs)
	fmt.Println("-----------")
	fmt.Println("-----------")
	fmt.Println("-----------")

	a := []int{0, 1, 2, 3, 4, 5}
	b := make([]int, 6)
	copy(b, a)
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-----------")
	a[0] = 7
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-----------")
}
