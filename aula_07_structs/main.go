package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	end endereco
}

type endereco struct{
	lograduro string 
	numero uint8
}

func main() {
	//Instanciar
	var u usuario
	u.nome = "Matheus"
	u.idade = 21
	fmt.Println(u)


	//Possamos dizer que  é um construtor com parameteros
	
	end := endereco{"Rua2 " , 42}
	u2 := usuario{"matheus", 21 ,end }
	fmt.Println(u2)

	//Passar apenas um parametro para a "CLASSE"
	u3 := usuario{nome:"José"}
	fmt.Println(u3)
}
