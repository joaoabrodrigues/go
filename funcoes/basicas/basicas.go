package main

import "fmt"

// sem parametros e sem retorno
func f1() {
	fmt.Println("Funcao 1")
}

// com parametros e sem retorno
func f2(p1 string, p2 string) {
	fmt.Printf("Funcao 2: %s %s\n", p1, p2)
}

// sem parametros e com retorno
func f3() string {
	return "Funcao 3"
}

// com parametros e com retorno
func f4(p1, p2 string) string {
	return fmt.Sprintf("Funcao 4: %s %s", p1, p2)
}

// com dois retornos
func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Println("Funcao 5:", r51, r52)
}
