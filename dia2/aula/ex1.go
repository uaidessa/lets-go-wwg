package main

import (
	"fmt"
)

type Apartamento struct {
	numero              int
	nomeDaProprietaria  string
	possuiVagaDeGaragem bool
}

func main() {
	ap1 := Apartamento{
		numero:              1301,
		nomeDaProprietaria:  "Talia",
		possuiVagaDeGaragem: true,
	}
	fmt.Printf("%+v\n", ap1)
	fmt.Printf("O apartamento numero %d tem como proprietária a %s. ele tem vaga na garagem? %t", ap1.numero, ap1.nomeDaProprietaria, ap1.possuiVagaDeGaragem)
}
