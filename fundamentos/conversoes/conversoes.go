package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal) // 6

	// cuidado
	fmt.Println("Teste", string(97)) // ASCII

	// int para string
	fmt.Println("Teste", strconv.Itoa(97))

	// string para int
	num, _ := strconv.Atoi("123") // _ ignora o valor retornado
	fmt.Println(num - 122)

	num2, err := strconv.Atoi("123")
	if err == nil {
		fmt.Println(num2 - 1)
	}

	b, _ := strconv.ParseBool("1") // true
	if b {
		fmt.Println("Verdadeiro")
	}
}
