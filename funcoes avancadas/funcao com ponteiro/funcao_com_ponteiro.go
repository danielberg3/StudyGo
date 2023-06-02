package main

import "fmt"

func main() {
	numero := 10
	fmt.Println(numeroInvertido(numero))
	fmt.Println(numero)

	novoNumero := 20

	fmt.Println(novoNumero)
	funcaoComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}

func numeroInvertido(numero int) int {
	return numero * -1
}

func funcaoComPonteiro(numero *int) {
	*numero = *numero * -1
}
