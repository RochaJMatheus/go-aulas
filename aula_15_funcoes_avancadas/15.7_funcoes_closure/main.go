package main

import "fmt"

func closure() func() {
	text := "Dentro da funcao CLOSURE"
	funcao := func() {
		fmt.Println(text)
	}
	return funcao
}

func main() {
	texto:= "Dentro da funcao MAIN"
	fmt.Println(texto)

	funcaNova := closure()
	funcaNova()
}