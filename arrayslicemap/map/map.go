package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[1] = "Maria"
	aprovados[2] = "Pedro"
	aprovados[3] = "Ana"
	fmt.Println(aprovados)

	for id, nome := range aprovados {
		fmt.Printf("%s (ID: %d)\n", nome, id)
	}

	fmt.Println(aprovados[3])
	delete(aprovados, 3)
	fmt.Println(aprovados)

}
