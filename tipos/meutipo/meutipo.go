package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "Conceito A"
	case n.entre(7.0, 8.99):
		return "Conceito B"
	case n.entre(5.0, 7.99):
		return "Conceito C"
	case n.entre(3.0, 5.99):
		return "Conceito D"
	default:
		return "Conceito E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
