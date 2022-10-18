package main

import (
	"errors"
	"fmt")

func main() {
	//Numeros inteiros
	var numero int = 100000
	fmt.Println(numero)

	var numero2 uint = 10000
	fmt.Println(numero2)

	//Alias
	//Int32 = Rune
	var numero3 rune = 123456
	fmt.Println(numero3)
	
	//BYTE = UINT8
	var numero4 byte = 122
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println("Numero real" , numeroReal1)

	var numeroReal2 float64 = 12300000000.00
	fmt.Println("Numero real2" , numeroReal2)

	//Inferencia de um numero real
	numeroReal3 := 123.000
	fmt.Println(numeroReal3)
	//Strings

	var str string = "Matheus Jose Rocha"
	fmt.Println("String" , str)

	//Inferencia de tip String
	str2 :="INferencia string"
	fmt.Println(str2)

	//"TIPO" Char
	char := 'B'
	fmt.Println("TIPO CHAR " , char)


	var texto string
	fmt.Println(texto)

	var int int16 
	fmt.Println(int)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)
	
	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}