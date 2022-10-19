package main

import (
	"fmt"
)

func main() {
	//Array tem que ser obrigatoriamente criado com um tamanho pré definido
	var array1 [5]string
	array1[0] = "Posição 1 "
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//Com slice não se tem a limitação de declaração de tamanho do mesmo
	slice := []int{10}
	fmt.Println(slice)

	//Add itens no slice
	slice = append(slice, 17)
	fmt.Println(slice)

	slice2 := array2[0:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	//Arrays internos
	//Função make cria um array de 15 posições retornando um slice que retorna as 10 primeiras posições
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Length
	fmt.Println(cap(slice3)) //Capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Length
	fmt.Println(cap(slice3)) //Capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) //Length
	fmt.Println(cap(slice4)) //Capacidade

}
