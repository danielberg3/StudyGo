package main

import "fmt"

func resultadoAluno(n1, n2 float64) bool {
	fmt.Println("Calculando média...")
	defer fmt.Println("Aprovado: ")
	var media float64 = (n1 + n2) / 2
	fmt.Println(media)

	if media > 6 {
		return true
	} else {
		return false
	}
}

func main() {
	resultado := resultadoAluno(6, 7)
	fmt.Println(resultado)
}
