package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()
	for {
		/*
			Usando o select Ele vai verificar se um canal já recebeu um dado
			Nesse caso o canal 1 não vai esperar o canal 2 completar o sleep de 2 segundos ele simplesmente vai executar assim que receber um valor
		*/
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}

}
