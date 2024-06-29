package main

import "fmt"

func main() {
	//Variaveis
	//var + nome da variavel + tipo
	var nome string
	nome = "Joao"
	fmt.Println(nome)

	nome = "Alex"
	fmt.Println(nome)

	var idade int
	idade = 18
	fmt.Println(idade)
	
	//Inferencia de tipos do GO
	var a = "Alex Silva"
	fmt.Printf("Type: %T - Value: %v\n", a, a)

	//Criando varias variaveis em um só comando
	var b, c = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//Tipo de declaracao mais usada, por conta dessa inferencia de tipos
	Str:= "apple"
	fmt.Println(Str)

	num := 1.5
	fmt.Println(num)

	//Constantes
	const cpf = 22222222222
	fmt.Println(cpf)
}