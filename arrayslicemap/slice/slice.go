package main

import (
	"fmt"
	"reflect"
)

func main() {
	array := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}

	fmt.Println(array, slice)
	fmt.Println(reflect.TypeOf(array), reflect.TypeOf(slice))

	array2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! Slice define um pedaço de um array
	slice2 := array2[1:3] // não cria um novo array, cria ponteiro para o primeiro elemento referenciado
	fmt.Println(array2, slice2)

	slice3 := array2[:2] // novo slice, mas aponta para o mesmo array
	fmt.Println(array2, slice3)

	// slice é: o tamanho e um ponteiro para um elemento de um array
	slice4 := slice2[:1]
	fmt.Println(slice2, slice4)
}
