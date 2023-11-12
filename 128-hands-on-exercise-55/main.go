package main

import "fmt"

type engine struct {
	electric bool
	number int
}

type vehicle struct {
	engine
	make string
	model string
	doors int
	color string
}

func main() {

	v1 := vehicle {
		engine: engine{
			electric: true,
			number: 4,
		},
		make: "toyota",
		model: "rav4",
		doors: 4,
		color: "read",
	}

	v2 := vehicle {
		engine: engine{
			electric: false,
			number: 5,
		},
		make: "honda",
		model: "crv",
		doors: 4,
		color: "white",
	}
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v1.engine.electric)
	fmt.Println(v2.engine)
	fmt.Println(v2.number)

}