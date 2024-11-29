package main

import "fmt"

func main() {
	for horas := 1; horas <= 5; horas++ {
		// uso do \n para quebrar linha
		fmt.Println("Executa o primeiro for.\n")
		fmt.Println("Horas:", horas)
		for min := 1; min <= 30; min++ {
			fmt.Println("Executa o segundo for.\n")
			fmt.Println("Minutos:\n", min)
		}
	}
}
