package main

import (
	"fmt"
)

func main() {

	var pesoDoTomate, quantidadeDeGarrafasDeAzeite, unidadesDeSabonete float64
	pesoDoTomate = 2.600
	quantidadeDeGarrafasDeAzeite = 2
	unidadesDeSabonete = 7

	var precoDoTomate, precoDoAzeite, precoDoSabonete float64
	precoDoTomate = 10.3
	precoDoAzeite = 19
	precoDoSabonete = 5.80

	valorDaCompra := (pesoDoTomate * precoDoTomate) + (quantidadeDeGarrafasDeAzeite * precoDoAzeite) + (unidadesDeSabonete * precoDoSabonete)
	fmt.Printf("o valor da compra deu: %v, só uma garrafa de azeite já custou %v", valorDaCompra, precoDoAzeite)

}
