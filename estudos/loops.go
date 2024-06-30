package main

import (
	"fmt"
	//"time"
)

//Loops
//Lacos de repeticoes
//Repetir tarefas

func main() {

	sum := 0
	//i == indice do for
	for i := 0; i <= 10; i++ {
		fmt.Println("Sum:", sum)
		fmt.Println("Indice:", i)
		sum += i
		// E como sum = sum + i
		// Para subtrair e  sum -= i -> sum = sum - i
		numero := 100
		fmt.Println("Numero:", numero)
		//Essa variavel numero so existe dentro do escopo do FOR
	}
	fmt.Println(sum)
	//fmt.Println("Numero:", numero)
	//da erro pq esta fora do FOR

	// for {
	// 	fmt.Println("Loop infinito ate vc dar o comando ctrl + c para sair")
	// 	time.Sleep(2 * time.Second)
	// 	//ctrl + c para sair do loop infinito
	// }

	// for range
	// frutas := []string{"laranja", "banana", "uva", "kiwi"}
	// for _, fruta := range frutas {
	// 	//O range pega elemento por elemento da lista
	// 	fmt.Println(fruta)
	// }

	frutas := []string{"laranja", "banana", "uva", "kiwi"}
	for i, fruta := range frutas {
		//O range pega elemento por elemento da lista
		fmt.Println("Fruta:", fruta, "Indice:", i)
	}
}