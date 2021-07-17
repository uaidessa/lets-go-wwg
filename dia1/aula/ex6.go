package main

import "fmt"

func main() {

	var numeros [6]string
	numeros[0] = "zero"
	numeros[1] = "um"
	numeros[2] = "dois"
	numeros[3] = "tres"
	numeros[4] = "quatro"
	numeros[5] = "cinco"

	fmt.Printf("o tipo é: %T\n", numeros)
	fmt.Println(numeros)

}
