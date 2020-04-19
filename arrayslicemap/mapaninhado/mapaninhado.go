package main

import "fmt"

func main()  {
	funcsPorLetra := map[string]map[string]float64 {
		"G": {
			"George": 1000.0,
			"Gabriela": 2000.0,
		},
		"J": {
			"João": 11120.00,
		},
		"P": {
			"Pedro Junior": 3000,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
