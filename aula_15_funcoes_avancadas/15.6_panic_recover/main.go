package main

import "fmt"

func recuperarExcucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExcucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A media é exatamente 6 \n")
}

func main() {
	//Vai dar panic
	fmt.Println(alunoAprovado(6, 6))
	fmt.Print("Pós execução")
}
