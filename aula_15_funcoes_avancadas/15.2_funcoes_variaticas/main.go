package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

func escrever(texto string, nums ...int) {
	for _, num := range nums {
		fmt.Println(texto, num)
	}
}

func main() {

	var total int
	total = soma(1, 2, 3, 4, 5, 5)
	fmt.Println("Total ", total)

	fmt.Println("-----")

	escrever("Ola", 1, 2, 3, 5, 6, 8, 7, 4, 5, 5)
}
