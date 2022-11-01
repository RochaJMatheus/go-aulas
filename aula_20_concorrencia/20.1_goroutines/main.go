package main

import (
	"fmt"
	"time"
)

func main() {

	go escrever("Ola mundo !!!") // Utilizar GO ROUTINE Executa a proxima as proximas linhas mesmo que a funcao não tenha terminado a execução
	
	escrever("Programando em GO ")
}

func escrever(texto string) {

	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
