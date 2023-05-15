package main

import (
	"errors"
	"fmt"
)

func main() {

	//Números

	var numero int32 = 1231
	fmt.Println(numero)

	var numero2 int64 = 165786
	fmt.Println(numero2)

	var numero3 uint = 31313
	fmt.Println(numero3)

	var numero4 byte = 255
	fmt.Println(numero4)

	var numero5 rune = 429496728
	fmt.Println(numero5)

	var numeroReal32 float32 = 3252532.43
	fmt.Println(numeroReal32)

	var numeroReal64 float64 = 43252.5235252
	fmt.Println(numeroReal64)

	var texto string = "palavra"
	fmt.Println(texto)

	char := 'B'
	fmt.Println(char)

	var boolean bool = true
	fmt.Println(boolean)

	var erro error = errors.New("Erro de validação")
	fmt.Println(erro)
}
