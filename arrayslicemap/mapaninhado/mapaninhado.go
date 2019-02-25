package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 15456.78,
			"Giovana":  8456.78,
		},
		"J": {
			"João": 11325.45,
		},
		"P": {
			"Paula": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")
	for letra, funcs := range funcsPorLetra {
		for nome, salario := range funcs {
			fmt.Println(letra, nome, salario)
		}
	}

}
