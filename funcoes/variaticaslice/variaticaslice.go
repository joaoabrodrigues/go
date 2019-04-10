package main

import "fmt"

func printApproved(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"João", "José", "Maria", "Ana"}
	printApproved(aprovados...)
}
