package main

import "fmt"

func main() {

	var variavel1 string = "variável 1"
	fmt.Println(variavel1)

	variavel2 := "variavel 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "variavel5", "variavel6"

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5)
	fmt.Println(variavel6)

}
