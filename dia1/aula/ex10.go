package main

import "fmt"

func main() {

	ano := make(map[int]string)
	ano[1] = "janeiro"
	ano[2] = "fevereiro"
	ano[3] = "março"
	ano[4] = "abril"
	ano[5] = "maio"
	ano[6] = "junho"
	ano[7] = "julho"
	ano[8] = "agosto"
	ano[9] = "setembro"
	ano[10] = "outubro"
	ano[11] = "novembro"
	ano[12] = "dezembro"

	fmt.Println(ano[1], ano[12])
	fmt.Println(len(ano))
}
