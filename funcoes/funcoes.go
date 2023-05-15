package main

import "fmt"

func main() {

	resultado := soma(43, 54)
	fmt.Println(resultado)

	var txt = func(txt string) {
		fmt.Println(txt)
	}

	txt("Hello World")

	primeiro, _ := doubleRetorno(7, 8)
	fmt.Println(primeiro)
	
}

func soma(n1 int8, n2 int8) int8 {
	soma := n1 + n2
	return soma
}

func doubleRetorno(n1 int8, n2 int8) (int8, int8) {
	return n1, n2
}
