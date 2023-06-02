package main

import "fmt"

func main() {

	texto := func() string {
		return "Hello world"
	}

	fmt.Println(texto())

	texto2 := func(texto string) string {
		return fmt.Sprintf("recebendo -> %s", texto)
	}("Olá mundo")

	fmt.Println(texto2)

}
