package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint
}

func main() {
	var u usuario

	u.nome = "Daniel"
	u.idade = 18
	fmt.Println(u)

	u2 := usuario{"martinho", 21, endereco{"Rua josefa Leonardo dos Santos", 54}}
	fmt.Println(u2)

	u3 := usuario{nome: "bruno"}
	fmt.Println(u3)

	endereco_u4 := endereco{"Rua josefa Leonardo dos Santos", 54}
	u4 := usuario{"marinalva", 51, endereco_u4}
	fmt.Println(u4)
}
