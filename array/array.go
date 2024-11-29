package main

func main() {
	// declaração de array do tipo any, aceitando qualquer tipo de dado
	var arr []any
	arr = append(arr, 1, "dois", 3, 4, "cinco", 6, 7, 8, 9, 10)

	// acessando o valor de um indice a outro.
	// fmt.Println(arr[3:6])

	// loop para exibir todos os indices.
	// i := 0
	// for i < len(arr) {
	// 	fmt.Println(arr[i])
	// 	i++
	// }

	// adicionando novos indices ao array existente
	var arr2 = []any{11, "doze", 13, 14, 15}
	arr = append(arr, arr2)
	i := 0
	// for i < len(arr) {
	// 	fmt.Println(arr[i])
	// 	i++
	// }

	// acessando array dentro de outro array
	for i < len(arr) {

	}

}
