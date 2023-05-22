package main

import "fmt"

func main() {
	var numero int = 100
	var ponteiro *int = &numero

	fmt.Println(numero, *ponteiro) //Desrreferenciar
	numero = 200
	fmt.Println(&numero, *ponteiro)

}
