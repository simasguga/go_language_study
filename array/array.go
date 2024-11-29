package main

import (
	"fmt"
)

func main() {

	// Declaração de array do tipo any, aceitando qualquer tipo de dado
	var tipoQualquer []any
	var tipoString []string
	var tipoNumerico []int

	// Atribuir valores ao array de qualquer tipo
	tipoQualquer = append(tipoQualquer, 1, "dois", 3, 4, "cinco", 6, 7, 8, 9, 10)

	// Atribuir valores ao array de string
	tipoString = append(tipoString, "Jorge", "Maria", "Felicio", "Ana", "Michele")

	// Atribuir valores ao array de números
	tipoNumerico = append(tipoNumerico, 2, 3, 5, 8, 12, 11, 9)

	// Adicionando novos índices ao array existente
	var arr2 = []any{11, "doze", 13, 14, 15}
	tipoQualquer = append(tipoQualquer, arr2)

	// Loop para acessar e verificar o tipo dos elementos
	for i := 0; i < len(tipoQualquer); i++ {
		switch v := tipoQualquer[i].(type) {
		case []any: // Verifica se o tipo é um slice de "any"
			fmt.Printf("Elemento no índice %d é um array: %v\n", i, v)
		default:
			// Ignorar outros tipos
		}
	}
}
