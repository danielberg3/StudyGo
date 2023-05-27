package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type estudante struct {
	pessoa
	curso     string
	matricula string
}

func main() {

	pessoa := pessoa{"Daniel", 18}
	estudante := estudante{pessoa, "Sistemas de Informação", "3132345"}

	fmt.Println(estudante)
}
