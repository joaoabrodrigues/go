package main

import "fmt"

// Não existe operador ternário em go
func obterResultado(nota float64) string {
	//return nota >= 7 ? "Aprovado" : "Reprovado"
	if nota >= 7 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(9.5))
}
