package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Matheus",
		"sobrenome": "Rocha",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Matheus",
			"ultimo":   "Rocha",
		},
		"curso": {
			"faculdade": "usp",
			"nome":      "ADS",
		},
	}
	fmt.Println(usuario2)
	//Deletar uma chave
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	//Adicionar uma nova chave
	usuario2["signo"] = map[string]string{
		"nome": "bla bla",
	}
	fmt.Println(usuario2)
}
