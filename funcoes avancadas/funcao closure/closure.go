package main

import "fmt"

func main() {
	texto := "Hello world"
	fmt.Println(texto)

	funcaoNova := closure()

	funcaoNova()

}

func closure() func() {
	texto := "Olá mundo"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}
