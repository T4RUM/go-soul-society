package main

import "fmt"

func main() {
	// Variáveis declaradas com `var` e tipo explícito.
	var (
		nome      string = "Wesley"
		idade     int    = 28
		estudando bool   = true
	)

	fmt.Println(nome, idade, estudando)

	// Declaração curta.
	// O Go infere que `curso` é uma string a partir do valor atribuído.
	curso := "Golang"

	fmt.Println(curso)

	// Variáveis relacionadas podem ser agrupadas em um bloco `var`.
	var (
		documento1 string = "Documento"
		documento2 string = "Document"
	)

	fmt.Println(documento1, documento2)

	// Constantes não podem ter seu valor alterado após a declaração.
	const pi float64 = 3.1415926

	fmt.Println(pi)
}
