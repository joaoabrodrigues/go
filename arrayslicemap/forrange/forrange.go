package main

import "fmt"

func main() {
	numeros := [...]int{3, 4, 5, 6, 7, 8} // compilador conta o tamanho do array

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
