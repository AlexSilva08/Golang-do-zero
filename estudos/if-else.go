package main

import "fmt"

func main() {
	valor := 2

	// if (condicao) { <acao> }
	if valor == 1 { //true
		fmt.Println("Valor igual a 1")
	} else if valor == 2 { //if true
		fmt.Println("Valor igual a 2")
	} else { //false
		fmt.Println("Valor nao e igual a 1 e tambem nao e igual a 2")
	}

	//if com operacoes
	fmt.Println(1 + 1)
	fmt.Println(2 - 1)
	fmt.Println(2 * 2)
	fmt.Println(10 / 2)
	fmt.Println(10 % 3)

	if valor % 2 == 0 {
		fmt.Printf("%d e par", valor)
	} else {
		fmt.Printf("%d e impar", valor)
	}
}