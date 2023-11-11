package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m[`bond_james`] = []string{`shaken`, `martinis`, `fast car`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}

	delete(m, `no_dr`)
	
	for k, v := range m {
		fmt.Println(k, v)
	}

	to_be_deleted := "dummy"
	v, ok := m[to_be_deleted]
	if !ok {
		fmt.Println("not found", to_be_deleted)
	} else {
		fmt.Printf("key=%#v, value=%#v deleted\n", to_be_deleted, v)
	}

	to_be_deleted = "bond_james"
	v, ok = m[to_be_deleted]
	if !ok {
		fmt.Println("not found", to_be_deleted)
	} else {
		fmt.Printf("key=%#v, value=%#v deleted\n", to_be_deleted, v)
	}

	s := make([]string, 3, 3)
	fmt.Println(s)
	fmt.Println(len(s))
}