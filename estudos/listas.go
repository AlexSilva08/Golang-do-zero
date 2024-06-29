package main

import "fmt"

func main() {
	
	//Array - Tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array)

	numPrimos := [6]int {2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	//Intervalo, apartir do zero ate a 2, no caso antes da posicao 3
	fmt.Println(numPrimos[0:3])
	//Depois da posicao 2
	fmt.Println(numPrimos[2:])

	//Slices
	//var slice []string
	slice := make([]string, 5)
	slice[0] = "Hello"
	slice[1]  = "World"

	fmt.Println(slice[0], slice[1])
	fmt.Println(slice)

	slice[3] = "Alex"

	fmt.Println(slice)

	numPares := []int {2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

	//Para "apendar" o valor tem que ser do mesmo tipo da slice
	numPares = append(numPares, 14, 16, 18)
	fmt.Println(numPares)

	//Listas

	//1 Arrays (tem tamanho fixo, tipo [7 valores])  e Slices (nao tem tamanho fixo): Homogeneos
	//len() - Retorna o tamanho das Arrays e Slices
	//append() - Adiciona valor nos slices

	//todos os elementos devem ter os mesmos tipos

	//[1, 2, 3, 4] - []int
	//["Alex", "Alice", "Golang"] - []string

	//2 Maps: Heterogeneos
	//pode mesclar tipos
	//estrutura chave - valor
	//[key] = value

	//A chave tem um tipo, ja o valor pode ter outro

	//map[string]int
	//{"Alex": 29, "Alice": 27}

	//map[string]string
	//{"Alex": "Silva", "Alice": "Pinheiro"}
}