package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Indice atual:", i)
	}

	fmt.Println("Novo loop iniciado!")
	for i := 6; i <= 12; i++ {
		fmt.Println("Indice atual:", i)
	}
	for i := 3; i <= 5; i++ {
		if i == 4 {
			fmt.Println("Deve printar apenas o valor definido no If")
			fmt.Println("Indice atual:", i)
			fmt.Println("Utilizando Loop para paralizar loop")
			break
		}
	}
}
