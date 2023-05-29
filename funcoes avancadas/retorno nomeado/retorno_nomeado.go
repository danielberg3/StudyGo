package main

import "fmt"

func somaSubtracao(n1, n2 int) (soma, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	fmt.Println(somaSubtracao(15, 40))
}
