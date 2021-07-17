package main

import (
	"fmt"
)

func main() {
	nascimento := 1965
	anoAtual := 2021
	idade := anoAtual - nascimento

	if idade < 18 {
		fmt.Printf("Você tem %d. Você é menor de idade!", idade)
	} else if idade >= 18 && idade < 60 {
		fmt.Printf("Você tem %d. Você é maior de idade!", idade)
	} else {
		fmt.Printf("Você tem %d. Você é menoridoso!", idade)
	}
}
