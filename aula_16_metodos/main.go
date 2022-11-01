package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	isMaiorIdade bool
}

/*
 Explicação método
*/
func (u usuario) salvar(){
	fmt.Printf("Salvando dos dados do usuário  %s no banco de dados\n" , u.nome)
}

/*
	Método pode retornar funções
*/
func (u usuario) maiorDeIdade() bool{
	return u.idade  >= 18
}


/*Método para alterar os dados do "usuário" necessário uso do ponteiro*/
func (u *usuario) fazerAniversario() {
	u.idade ++
}


func main() {
	usuario1 := usuario{nome: "Matheus", idade: 21}
	fmt.Println("Dados usuário 1 " , usuario1)
	usuario1.salvar()

	usuario2 := usuario{nome: "José", idade:21}
	fmt.Println("Dados do usuário 2 ",usuario2)
	usuario2.salvar()

	//Verificando usuário maior de idade com retorno
	usuario1.isMaiorIdade = usuario1.maiorDeIdade()
	fmt.Println("Usuário maior de idade: ",usuario1.isMaiorIdade)
	usuario2.fazerAniversario()
	fmt.Println("Idade do usuário 2 " ,usuario2.idade )
}