package main

import "fmt"

func main() {
	// Função simples com parâmetros e um único retorno.
	resultado := somar(15, 10)
	fmt.Println("Soma:", resultado)

	// Função anônima armazenada em uma variável.
	minhaFuncao()

	// Uma função pode retornar vários valores.
	soma, subtracao, multiplicacao, divisao := calculadora(10, 5)

	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
	fmt.Println("Multiplicação:", multiplicacao)
	fmt.Println("Divisão:", divisao)

	// Quando não precisamos de um dos valores retornados,
	// podemos ignorá-lo usando o identificador em branco (_).
	nome, _ := nomeCompleto("Wesley", "Murat")

	fmt.Println("Nome:", nome)
}

// somar recebe dois números inteiros e retorna a soma entre eles.
func somar(num1, num2 int) int {
	return num1 + num2
}

// Uma função também pode ser armazenada dentro de uma variável.
// Mesmo estando em uma variável, o tipo de minhaFuncao continua sendo func().
var minhaFuncao = func() {
	fmt.Println("Essa função está armazenada em uma variável.")
}

// calculadora recebe dois números inteiros e retorna quatro valores.
// Os retornos correspondem a soma, subtração, multiplicação e divisão.
func calculadora(n1, n2 int) (int, int, int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2

	return soma, subtracao, multiplicacao, divisao
}

// nomeCompleto demonstra uma função que retorna mais de um valor.
func nomeCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}
