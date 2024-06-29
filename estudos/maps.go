package main

import "fmt"

func main() {
	idade := map[string]int{}
	idade["Alex"] = 29
	idade["Alice"] = 27
	fmt.Println(idade)
	fmt.Println(idade["Alex"])
	fmt.Println(idade["Alice"])

	anoNasc := map[string]int{
		"Alex": 1994,
		"Alice": 1996
	}
}