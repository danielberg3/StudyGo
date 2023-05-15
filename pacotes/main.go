package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Hello world")
	auxiliar.Escrever()
	err := checkmail.ValidateFormat("aleatorio@gmail.com")
	fmt.Println(err)
}
