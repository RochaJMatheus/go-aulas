package main

import (
	"fmt"
	"time"
)

func main() {
	//Só vai poder enviar e receber strings
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	fmt.Println("Depois da funçao escrever ser executada ")

	//Aguardando receber dados //Obriga canal receber valor

	//Entra em dead lock pois espera um dado que nunca vai chegar para o canal
	/*
		for {
			mensagem := <-canal
			fmt.Println(mensagem)
		}
	*/
	//Para tratar deadlock basta verificar se o canal está aberto ou fechado OBS deve-se fechar os mesmo
	/*
		for {
			mensagem, aberto := <-canal
			if !aberto {
				fmt.Println("Canal fechado fim do programa")
				break
			}
			fmt.Println(mensagem)
		}
	*/

	//Outra sintaxe para utilizar canais
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//Enviar dado pelo canal ( chan <- valor está enviando dados ) (<- chan está esperando para receber um valor)
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
