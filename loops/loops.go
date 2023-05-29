package main

import (
	"fmt"
)

func main() {

	i := 0

	for i < 10 {
		i++
		fmt.Println("Executando...", i)
		//time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	//for {
	//	fmt.Println("executando ...")
	//	time.Sleep(time.Second)
	//}

	usuario := map[string]string{
		"nome":      "Daniel",
		"sobrenome": "Berg",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	list := [5]string{"Animal", "Fruta", "Objeto", "País", "Nome"}

	for _, nome := range list {
		fmt.Println(nome)
	}

	for indice, letra := range "DANIEL" {
		fmt.Println(indice, string(letra))
	}
}
