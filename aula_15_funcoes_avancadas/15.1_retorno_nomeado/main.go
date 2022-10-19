package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	return n1 + n2, n1 - n2
}

func calculosMatematicos2(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma1, sub1 := calculosMatematicos(10, 5)
	fmt.Println(soma1, sub1)

	fmt.Println("------")

	soma2  , sub2 := calculosMatematicos2(10 , 5)
	fmt.Println(soma2 , sub2)
}
