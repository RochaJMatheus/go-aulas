package main

import "fmt"

func main() {

	//Operadores aritmeticos + - / * %

	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 10
	multiplicacao := 10 * 2
	restoDiv := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDiv)

	/* Tipos diferentes não podem intergir ex abaixo >
	var numero1 int16 = 10
	var numero2 int32 = 25

	soma := numero1 + numero2
	fmt.Println(soma)
	*/

	var numero1 int16 = 10
	var numero2 int16 = 25

	soma1 := numero1 + numero2
	fmt.Println(soma1)

	//Atribuicao
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	fmt.Println("--------")
	//OPeradores Lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//Operadores Unarios INCREMETADORES

	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 //numero = numero - 20
	fmt.Println(numero)

	numero *= 3  //numero = numero * 3
	numero /= 10 // numero = numero / 10
	numero %= 3
	fmt.Println(numero)

	//Operador Ternario não existe em GOLANG

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5 "
	}

	fmt.Println(texto)
}
