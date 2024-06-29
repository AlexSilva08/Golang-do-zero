package main

import "fmt"

//Structs
//Criando sua propria estrutura de dados

//type <nome da estrutura> struct { <campos> }
type Pessoa struct {
	Nome string
	Sobrenome string
	Idade int
	AnoNasc int
	CadAtivo bool
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
fmt.Println(Pessoa{"Alex", "Silva", 29, 1994, true})
fmt.Println(Pessoa{Nome: "Alice", Sobrenome: "Pinheiro", Idade: 27, AnoNasc: 1996, CadAtivo: true})
fmt.Println(Pessoa{Nome: "Joao"})//o restante sera setado como zero value

//Escolhendo um campo especifico
people1 := Pessoa{Nome: "Alex", Sobrenome: "Silva", Idade: 29, AnoNasc: 1994, CadAtivo: true}
fmt.Println(people1.Nome)

people1.Idade = 30

people2 := Pessoa{Nome: "Alice", Sobrenome: "Pinheiro", Idade: 27, AnoNasc: 1996, CadAtivo: true}

pessoas := []Pessoa{}
pessoas = append(pessoas, people1, people2)
//fmt.Println(pessoas)

//Struct + map
// alunos := map[string][]Pessoa{}
// alunos["programacao"] = pessoas
// fmt.Println(alunos)

//Setando direto
var alunos = map[string][]Pessoa{
	"Programacao": {{Nome: "Alex", Sobrenome: "Silva", Idade: 29, AnoNasc: 1994, CadAtivo: true}, {Nome: "Alice", Sobrenome: "Pinheiro", Idade: 27, AnoNasc: 1996, CadAtivo: true}},
	"Engenharia":{{Nome: "Ale", Sobrenome: "Pere", Idade: 29, AnoNasc: 1994, CadAtivo: true}, {Nome: "Maria", Sobrenome: "Alice", Idade: 27, AnoNasc: 1996, CadAtivo: true}},
}
fmt.Println(alunos)

//struct herdando campos de outra struct
prof := Profissao{people1, "dev"}
fmt.Println(prof)
//Buscando um campo de Pessoa
fmt.Println(prof.Pessoa.AnoNasc)
}