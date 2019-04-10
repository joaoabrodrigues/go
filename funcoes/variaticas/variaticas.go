package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	media := media(8.9, 9.0, 9.3, 9.7, 10.0)
	fmt.Printf("Média: %.2f", media)
}
