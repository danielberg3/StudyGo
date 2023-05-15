package main

import "fmt"

func main() {

	//Operadores Aritméticos

	fmt.Println(5 + 3)
	fmt.Println(5 - 3)
	fmt.Println(5 * 3)
	fmt.Println(5 / 3)
	fmt.Println(5 % 3)

	fmt.Println("-----------------------------------")

	//Operadores de atribuição

	soma := 4 + 8
	fmt.Println(soma)

	var soma2 int = 5 + 7
	fmt.Println(soma2)

	fmt.Println("-----------------------------------")

	//Operadores relacioanais

	fmt.Println(3 > 5)
	fmt.Println(3 >= 5)
	fmt.Println(3 < 5)
	fmt.Println(3 <= 5)
	fmt.Println(3 == 5)
	fmt.Println(3 != 5)

	fmt.Println("-----------------------------------")

	//Operadores lógicos

	verdadeiro := true
	falso := false

	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	fmt.Println("-----------------------------------")

	//Operadores unários

	numero := 10

	numero++
	numero--

	numero += 10
	numero -= 10

	numero *= 3
	numero /= 3

	numero %= 10

	fmt.Println(numero)

}
