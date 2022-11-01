package main

import "fmt"

func main() {
	// Diferença entre canal com buffer e sem canal com Buffer só bloqueia quando sua capacidade máxima é atingida
	canal := make(chan string, 2)
	canal <- "Olá mundo"
	canal <- "Programando em GO"
	//Estourou buffer vai dar deadlock make(chan string, 2) definido dentro dele que temos 2 canais  , se o canal atingir o máximo ele vai dar block deadlock
	//canal <- "OPA DEADLOCK"

	mensagem := <-canal
	fmt.Println(mensagem)
}
