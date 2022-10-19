package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1 ")
}

func funcao2 (){
	fmt.Println("Executando a funcao 2")
}

func alunoAprovado(n1 , n2 float32) bool{
	//Será executado imediatamamente antes da sintaxe so return 
	defer fmt.Println("Media calulada Resultado será retornado ")
	fmt.Println("Entrando na funcao para verificar se o aluno está aprovado")
	media := n1 + n2 / 2
	if media >= 6{
		
		return true
	}else {
		return false
	}
}

func main() {
	//Defer adia a excução da funcao até o ultimo momento possivel
	fmt.Println(alunoAprovado(7,8))
	
}
