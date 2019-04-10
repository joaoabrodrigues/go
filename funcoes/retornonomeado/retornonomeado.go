package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main() {
	r1, r2 := trocar(2, 3)
	fmt.Println(r1, r2)

	segundo, primeiro := trocar(7, 8)
	fmt.Println(segundo, primeiro)
}
