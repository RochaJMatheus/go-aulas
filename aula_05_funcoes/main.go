package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//Funções podem ter mais de um retorno !!!


func calculosMatematicos(n1 ,n2 int8) (int8,int8){
	soma:= n1 + n2
	subtracao := n1 - n2
	return soma , subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(log string) string {
		fmt.Println(log)
		return log
	}
	result := f("Olha a log")
	fmt.Println(result)

	resultadoSoma1 , resultadoSubtracao := calculosMatematicos(10 , 15)
	fmt.Println(resultadoSoma1 , resultadoSubtracao)

	//Usando apenas um valor dos returns
	resultadoSoma2 , _ := calculosMatematicos(10 , 15)
	fmt.Println(resultadoSoma2)

}
