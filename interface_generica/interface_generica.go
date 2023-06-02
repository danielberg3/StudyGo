package main

import "fmt"

func main() {
	escreverValor(true)
	escreverValor(123)
	escreverValor("Daniel")
	escreverValor(4242.421)
}

func escreverValor(i interface{}) {
	fmt.Println(i)
}
