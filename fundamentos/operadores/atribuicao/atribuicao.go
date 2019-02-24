package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 4 // inferência de tipo
	i += 4 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	fmt.Println(i)

	x, y := 1, 2 // criando e inicializando
	fmt.Println(x, y)

	x, y = y, x // atribuindo
	fmt.Println(x, y)
}
