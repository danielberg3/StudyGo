package main

import "fmt"

func main() {
	fmt.Println(calcularmedia(6, 6))
	fmt.Println("pós execução")
}

func calcularmedia(n1, n2 float64) bool {
	media := (n1 + n2) / 2

	defer tratarPanic()

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("número igual a 6")
}

func tratarPanic() {
	if r := recover(); r != nil {
		fmt.Println("Erro tratado")
	}
}
