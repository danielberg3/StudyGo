package main

import "fmt"

func main() {

	var array [5]int

	array[0] = 3
	fmt.Println(array)

	array2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(array2)

	array3 := [...]string{"posição 1", "posição 2", "posição 3", "posição 4"}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice = append(slice, 9)
	fmt.Println(slice)

	slice2 := array2[0:3]
	fmt.Println(slice2)

	slice3 := make([]float32, 10, 15)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
