package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"João":    11325.45,
		"Maria":   15456.78,
		"Rafaela": 1200.0,
	}

	funcsESalarios["José"] = 1350.0
	delete(funcsESalarios, "não existe e nada acontece, feijoada")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

	for _, salario := range funcsESalarios {
		fmt.Println(salario)
	}

	for nome, _ := range funcsESalarios { // equivalente a -> for nome := range
		fmt.Println(nome)
	}
}
