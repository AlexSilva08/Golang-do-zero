package main

import "fmt"

func main() {
	somaDosValores := soma(42,13)
	fmt.Println(somaDosValores)

	fmt.Println("soma de 42 + 13 = ")
	fmt.Println(soma(42, 13))

	fmt.Println(somaFloat(25.5, 25.3))
	fmt.Println(subtracao(5,2))

	name1, _ , name3 := printName("Alex")
	fmt.Println(name1)
	//fmt.Println(name2)
	fmt.Println(name3)

	name, sobrenome := printCompleteName("Alex", "Silva")
	fmt.Println(name)
	fmt.Println(sobrenome)

	namePub := PrintaNomePub("Alex Publico")
	fmt.Println(namePub)
}

//Exemplo soma 1
// func soma (x int, y int) int {
// 	return x + y
// }

//Exemplo soma 2
func soma (x int, y int) int {
	somaDosValores := x + y
	return somaDosValores
}

func somaFloat (a float64, b float64) float64 {
	return a + b
}

//Exemplo subtracao
func subtracao (x int, y int) int {
	return x - y
}

//Exemplo de funcao retornando strings
func printName (name string) (string, string, string) {
	return name, name, name
}

//Repare que a funcao comeca com letra minuscula, ou seja a funcao eh privada
//Portanto soh pode ser usada dentro do proprio pacote, que nesse caso eh o main
func printCompleteName (nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

//funcao publica
func PrintaNomePub (nome string) string {
	return nome
}