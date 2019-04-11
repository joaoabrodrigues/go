package main

import "fmt"

type sport interface {
	ligarTurbo()
}

type luxury interface {
	fazerBaliza()
}

type luxurySport interface {
	sport
	luxury
	// pode adicionar mais métodos
}

type bmw struct{}

func (b bmw) ligarTurbo() {
	fmt.Println("Turbo ligado!")
}

func (b bmw) fazerBaliza() {
	fmt.Println("Estacionado!")
}

func main() {
	var b luxurySport = bmw{}
	b.ligarTurbo()
	b.fazerBaliza()
}
