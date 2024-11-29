// inicialiar com package main.

package main

// importar o pacote fmt responsável por realizar prints no terminal.
import "fmt"

// definir a função main que será executada quando o comando "go run nome_do_arquivo.go for executado no terminal"

// Variaveis fora do escopo da função devem ser declaradas da forma "Tipo de variavel - Nome da Variavel - Tipo do valor"
// exemplo

var numero int = 12

const texto string = "texto da variavel"

func main() {
	// dentro do escopo da função é permitido declarar variaveis de forma curta
	texto2 := "texto 2"
	numero2 := 21
	fmt.Println(texto2, ",", numero2)
}

// abrir terminal e executar "go run nome_do_arquivo.go" arquivo atual = go run variaveis.go
