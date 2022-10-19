package main

import (
	"fmt"
)

func main() {
	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	}else{
		fmt.Println("Menor ou igual a 15")
	}
	
	//If init
	if outroNumero := numero; outroNumero > 0{
		fmt.Println("NUMERO É MAIOR QUE ZERO .....")
	}else{
		fmt.Println("NUMERO É MENOR QUE ZERO .....")
	}
}