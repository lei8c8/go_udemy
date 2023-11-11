package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m[`bond_james`] = []string{`shaken`, `martinis`, `fast car`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}

	for k, v := range m {
		fmt.Println(k, v)
	}

	n := map[string][]string {
		`bond_james`: {`shaken`, `martinis`, `fast car`},
		`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`: {`cats`, `ice cream`, `sunsets`},
	}

	for k, v := range n {
		fmt.Printf("%#v, %#v\n", k, v)
	}

	n[`fleming_ian`] = []string{`steals`, `cigars`, `espionage`}
	for k, v := range n {
		fmt.Printf("%#v, %#v\n", k, v)
	}
}