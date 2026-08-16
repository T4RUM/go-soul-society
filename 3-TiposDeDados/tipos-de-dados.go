package main

import (
	"errors"
	"fmt"
)

func main() {
	// Números inteiros com sinal: int8, int16, int32 e int64.
	// O valor abaixo não caberia em um int8, cujo limite máximo é 127.
	var numero int64 = 10000000000
	fmt.Println(numero)

	// int possui tamanho dependente da arquitetura (32 ou 64 bits).
	// Usamos aqui um valor que cabe nas duas possibilidades.
	var numeroDois int = 1000000000
	fmt.Println(numeroDois)

	// Com :=, um literal inteiro recebe int como tipo padrão.
	numeroTres := 10000000
	fmt.Println(numeroTres)

	// Inteiros sem sinal não representam valores negativos.
	var numeroQuatro uint = 1000
	fmt.Println(numeroQuatro)

	// rune é um alias para int32 e normalmente representa um code point Unicode.
	var numeroCinco rune = 123456
	fmt.Println(numeroCinco)

	// byte é um alias para uint8.
	var numeroSeis byte = 8
	fmt.Println(numeroSeis)

	// Números de ponto flutuante: float32 e float64.
	var numeroSete float32 = 0.1
	fmt.Println(numeroSete)

	var numeroOito float64 = 0.22222222
	fmt.Println(numeroOito)

	// Um literal decimal inferido com := recebe float64 como tipo padrão.
	numeroNove := 12313.54
	fmt.Println(numeroNove)

	// Strings representam texto.
	var nome string = "Wesley"
	fmt.Println(nome)

	curso := "Golang"
	fmt.Println(curso)

	// Aspas simples representam uma rune, não uma string.
	caractere := 'B'
	fmt.Println(caractere)
	fmt.Printf("%c\n", caractere)

	// Valores booleanos podem ser true ou false.
	var booleanoUm bool = false
	fmt.Println(booleanoUm)

	// error é uma interface embutida. Seu valor zero é nil.
	var erro error
	fmt.Println(erro)

	// errors.New cria um novo valor que implementa a interface error.
	var erroInterno error = errors.New("erro interno")
	fmt.Println(erroInterno)
}
