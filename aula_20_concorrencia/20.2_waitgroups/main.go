package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	//Sempre especificar a quantidade de rotinas que ele tem que realizar
	waitGroup.Add(2)
	go func() {
		escrever("Olá mundo")
		waitGroup.Done()
		//Add -1
	}()
	go func() {
		escrever("Programando em go")
		waitGroup.Done()

		//Add -1
	}()

	//Função main espera a contagem as go routines chegar a zero
	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
