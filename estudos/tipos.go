package main

import "fmt"

func main() {
	//Tipos:

	//bool (true/false)
	fmt.Printf("Type: %T - Value: %v\n", true, true)

	//string - sequencia de bytes
	fmt.Printf("Type: %T - Value: %v\n", "alx", "alx")
	fmt.Printf("Type: %T - Value: %v\n", "1", "1")

	//int
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)

	//float (flaot64/float32) - decimal
	fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.233)
}

