package main

import (
	"fmt"
)

func main() {
	nascimento := 2006
	anoAtual := 2021
	idade := anoAtual - nascimento

	if idade >= 16 {
		fmt.Printf("Você tem %d. Você está apto a votar!", idade)
		return // se for falso pula pro final - substitui o uso do else
	}
	fmt.Printf("Você tem %d. Você não pode votar!", idade)

}
