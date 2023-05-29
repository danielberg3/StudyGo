package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":  "Daniel",
		"nome2": "Lucas",
	}

	fmt.Println(usuario)
	delete(usuario, "nome2")
	fmt.Println(usuario)

	usuario2 := map[string]map[int]string{
		"nomes": {
			1: "Pedro",
			2: "João",
			3: "Paulo",
		},

		"idades": {
			1: "15",
			2: "16",
			3: "17",
		},
	}

	fmt.Println(usuario2["nomes"])
}
