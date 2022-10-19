package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando I ", i)
		time.Sleep(time.Second)
	}
	fmt.Println("---------------")
	for j := 0; j < 10; j++ {
		i++
		fmt.Println("Incrementando J ", j)
		time.Sleep(time.Second)
	}
	fmt.Println("---------------")
	//For each
	nomes := [3]string{"Matheus", "José", "Rocha"}
	for _, nome := range nomes {
		fmt.Println("Nome", nome)
	}
	fmt.Println("---------------")

	//Iterar string
	for indice, letra := range "PALAVRA" {
		fmt.Println("INDICE: ", indice, "Letra:", letra)
		fmt.Println("---------------")

		fmt.Println("INDICE: ", indice, "Letra:", string(letra))
	}

	//Iterar maps
	usuario := map[string]string{
		"Nome":      "Matheus",
		"Sobrenome": "Rocha",
	}
	for chave, valor := range usuario {
		fmt.Println("Chave: ", chave, "Valor: ", valor)
	}
}
