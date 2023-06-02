package main

import "fmt"

func main() {
	u := usuario{"Daniel", 18}

	u.salvarUser()

	maiorIdade := u.maiordeIdade()
	fmt.Println(maiorIdade)

	u.aniversario()
	fmt.Println(u.idade)
}

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvarUser() {
	fmt.Printf("Salvando usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiordeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) aniversario() {
	u.idade++
}
