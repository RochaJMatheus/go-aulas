package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Numero inválido"
	}

}

func diaDaSemana2(numero int) string {
	var diaSemana string
	switch {
	case numero == 1:
		diaSemana = "Domingo"
		//fallthrough Joga para proxima condição mesmo atendendo a condição acima
	case numero == 2:
		diaSemana = "Segunda-Feira"
	case numero == 3:
		diaSemana = "Terça-Feira"
	case numero == 4:
		diaSemana = "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		diaSemana = "Sexta-Feira"
	case numero == 7:
		diaSemana = "Sábado"
	default:
		diaSemana = "Numero inválido"
	}
	return diaSemana
}
func main() {

	dia := diaDaSemana(3)
	fmt.Println(dia)

	fmt.Println("------------")

	dia2 := diaDaSemana2(2)
	fmt.Println(dia2)
}
