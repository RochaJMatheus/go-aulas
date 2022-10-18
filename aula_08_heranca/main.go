package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	//Instanciação

	p1 := pessoa{nome: "Matheus", sobrenome: "Rocha", idade: 20, altura: 170}
	fmt.Println(p1)

	e1 := estudante{pessoa: p1 ,curso: "ADS" ,faculdade: "Unicesumar"}
	fmt.Println(e1.nome)
}