package main

import "fmt"

func _switch() string {
	numero := 4

	switch numero {
	case 1:
		return "Número igual a 1"
	case 2:
		return "Numero igual a 2"
	case 3:
		return "Número igual a 3"
	default:
		return "Número maior que 3 e menor que 1"

	}
}

func main() {
	resultado := _switch()
	fmt.Println(resultado)
}
