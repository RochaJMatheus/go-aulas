package main

import "fmt"

func main() {

	retorno := func(text string) string {
		return fmt.Sprintf("Recebido -> %s", text)
	}("Oi eu sou uma função anonima")
	fmt.Println(retorno)
}
