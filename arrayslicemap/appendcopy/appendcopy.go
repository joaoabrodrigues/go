package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	var slice []int

	// array = append(array, 4, 5, 6)
	slice = append(slice, 4, 5, 6) // se no append, for adicionado mais elementos que a capacidade do slice, seu tamanho é dobrado automaticamente
	fmt.Println(array, slice)

	slice2 := make([]int, 2)
	copy(slice2, slice)
	fmt.Println(slice2)
}
