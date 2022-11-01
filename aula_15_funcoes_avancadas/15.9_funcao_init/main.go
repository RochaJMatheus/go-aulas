package main

import "fmt"

//Função que será executada antes da função main
//Muito util para inicializar algo antes de executar 
//POde se ahver uma função init por arquivo

 var  n int
func init(){
	fmt.Println("Vai executar antes da funcao main")
	n = 10
}


func main(){
	fmt.Println("Vai excutar a main " ,n)
}