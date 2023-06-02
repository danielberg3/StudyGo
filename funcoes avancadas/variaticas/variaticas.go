package main

import "fmt"

func main() {
	resultado := soma(1, 2, 6, 432, 42, 64, 5)
	fmt.Println(resultado)
}

func soma(numeros ...int) int {
	soma := 0

	for _, numero := range numeros {
		soma += numero
	}

	return soma
}
