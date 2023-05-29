package main

import "fmt"

func main() {
	if num := -2; num > 4 {
		fmt.Println("Número é maior que 4")
	} else if num < 0 {
		fmt.Println("Número é menor que 0")
	} else {
		fmt.Println("Número está entre 4 e 0")
	}
}
