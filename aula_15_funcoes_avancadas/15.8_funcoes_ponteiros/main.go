package main

import "fmt"

//Parametro por copia
func inverteSinal(n1 int) int {
	return n1 * -1
}

//Parametro por referencia
func inverteSinalComPonteiro(n1 *int){
	*n1 = *n1 * -1
}

func main() {
	numero := 20
	numeroInvertido := inverteSinal(numero)
	fmt.Println(numeroInvertido)
	nvNumero := 40
	fmt.Println("Novo numero" , nvNumero)
	inverteSinalComPonteiro(&nvNumero)
	fmt.Println("Novo numero Após a função de inversão com ponteiro" , nvNumero)
}
