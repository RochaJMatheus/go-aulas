package main

import "fmt"

func main() {
	//Variavel string
	var variavelStr string = "Matheus"

	//Inferencia de tipo
	varivelStr2 := "José"

	fmt.Println(variavelStr)
	fmt.Println(varivelStr2)
	var (
		var4Str = "Rocha"
		var5Str = "<><><>"
	)
	fmt.Println(var4Str, var5Str)

	//Constantes
	const cosntante1 string = "Constante1"
	fmt.Println(cosntante1)

	//inversao de valore entre variaveis

	variavelStr, varivelStr2 = varivelStr2, variavelStr
	fmt.Println(variavelStr, varivelStr2)
}
