package main

import "fmt"

func main() {
	posicao := 2

	//switch valor { <cases equals> }
	switch posicao {
	case 1:
		fmt.Println("Primeiro lugar")
	case 2:
		fmt.Println("Segundo lugar")
	case 3:
		fmt.Println("Terceiro lugar")
	}

	nome := "Colt"

	switch nome {
	case "Alex":
		fmt.Printf("%v e uma pessoa", nome)
	case "Bem te vi":
		fmt.Printf("%v e um passaro", nome)
	case "Colt":
		fmt.Printf("%v e um cachorro", nome)
	}
}