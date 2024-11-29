package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		imp_este := 4
		if i == imp_este {
			fmt.Println("indice definido: ", imp_este)
		} else {
			fmt.Println("indice não definido: ", i)
		}
	}
}
