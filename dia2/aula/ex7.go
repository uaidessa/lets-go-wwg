package main

import (
	"fmt"
)

type Pessoa struct {
	nome         string
	idade        int
	corPreferida string
}

func main() {
	pessoa1 := Pessoa{
		nome:         "Andressa",
		idade:        29,
		corPreferida: "Verde",
	}
	pessoa2 := Pessoa{
		nome:         "Jamie",
		idade:        30,
		corPreferida: "preto",
	}
	pessoa3 := Pessoa{
		nome:         "Steve",
		idade:        35,
		corPreferida: "azul",
	}
	fmt.Println(pessoa1.nome, pessoa2.nome, pessoa3.nome)
	fmt.Printf("A pessoa se chama %s e tem %d anos e sua cor preferida é %s\n", pessoa1.nome, pessoa1.idade, pessoa1.corPreferida)
	fmt.Printf("A pessoa se chama %s e tem %d anos e sua cor preferida é %s\n", pessoa2.nome, pessoa2.idade, pessoa2.corPreferida)
	fmt.Printf("A pessoa se chama %s e tem %d anos e sua cor preferida é %s", pessoa3.nome, pessoa3.idade, pessoa3.corPreferida)
}
