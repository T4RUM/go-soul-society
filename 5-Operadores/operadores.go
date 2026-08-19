package main

import "fmt"

func main() {
	// =========================================================================
	// OPERADORES ARITMÉTICOS
	// =========================================================================

	// Operadores aritméticos são utilizados para realizar cálculos.
	//
	// +  → soma
	// -  → subtração
	// *  → multiplicação
	// /  → divisão
	// %  → resto da divisão inteira

	soma := 2 + 2
	subtracao := 2 - 5
	multiplicacao := 2 * 4
	divisao := 8 / 2
	restoDaDivisao := 10 % 3

	fmt.Println("=== Operadores aritméticos ===")
	fmt.Println("2 + 2 =", soma)
	fmt.Println("2 - 5 =", subtracao)
	fmt.Println("2 * 4 =", multiplicacao)
	fmt.Println("8 / 2 =", divisao)
	fmt.Println("10 % 3 =", restoDaDivisao)

	// Quando os dois valores envolvidos na divisão são inteiros,
	// o resultado também será inteiro.
	//
	// 9 / 2 matematicamente seria 4.5,
	// mas como estamos dividindo dois inteiros, o resultado será 4.
	divisaoInteira := 9 / 2

	fmt.Println("9 / 2 =", divisaoInteira)

	// =========================================================================
	// OPERAÇÕES ENTRE TIPOS DIFERENTES
	// =========================================================================

	// Mesmo que int16 e int32 representem números inteiros,
	// eles continuam sendo tipos diferentes para o Go.

	var numero16 int16 = 10
	var numero32 int32 = 20

	// Isto NÃO compila:
	//
	// somaInvalida := numero16 + numero32
	//
	// O Go não realiza automaticamente a conversão entre os tipos.

	// Para realizar a operação, podemos converter explicitamente
	// um dos valores para o tipo do outro.
	somaConvertida := int32(numero16) + numero32

	fmt.Println("\n=== Tipos diferentes ===")
	fmt.Println("int32(numero16) + numero32 =", somaConvertida)

	// =========================================================================
	// DECLARAÇÃO E ATRIBUIÇÃO
	// =========================================================================

	// Com `var`, podemos declarar uma variável informando
	// explicitamente seu tipo.
	var nome string = "Wesley"

	// `:=` é chamado de declaração curta.
	//
	// Ele cria a variável e permite que o Go descubra
	// automaticamente seu tipo através do valor.
	sobrenome := "Murat"

	fmt.Println("\n=== Declaração de variáveis ===")
	fmt.Println("Nome:", nome)
	fmt.Println("Sobrenome:", sobrenome)

	// O operador `=` possui outro objetivo.
	//
	// Ele atribui um novo valor a uma variável que já existe.
	linguagem := "Go"

	fmt.Println("Antes:", linguagem)

	linguagem = "Golang"

	fmt.Println("Depois:", linguagem)

	// Resumindo:
	//
	// := → declara uma variável
	// =  → atribui um valor

	// =========================================================================
	// OPERADORES RELACIONAIS
	// =========================================================================

	// Operadores relacionais são utilizados para comparar valores.
	//
	// O resultado de uma comparação sempre será um bool:
	//
	// true
	// false
	//
	// >   → maior que
	// >=  → maior ou igual
	// <   → menor que
	// <=  → menor ou igual
	// ==  → igual
	// !=  → diferente

	numeroA := 1
	numeroB := 2

	fmt.Println("\n=== Operadores relacionais ===")

	fmt.Println("1 > 2  =", numeroA > numeroB)
	fmt.Println("1 >= 2 =", numeroA >= numeroB)
	fmt.Println("1 < 2  =", numeroA < numeroB)
	fmt.Println("1 <= 2 =", numeroA <= numeroB)
	fmt.Println("1 == 2 =", numeroA == numeroB)
	fmt.Println("1 != 2 =", numeroA != numeroB)

	// Atenção:
	//
	// =  → atribuição
	// == → comparação de igualdade

	// =========================================================================
	// OPERADORES LÓGICOS
	// =========================================================================

	// Operadores lógicos trabalham principalmente com valores booleanos.
	//
	// && → E lógico (AND)
	// || → OU lógico (OR)
	// !  → NÃO lógico (NOT)

	temIdadeMinima := true
	temDocumento := false

	fmt.Println("\n=== Operadores lógicos ===")

	// AND
	//
	// Só será true quando os dois valores forem true.
	fmt.Println(
		"true && false =",
		temIdadeMinima && temDocumento,
	)

	// OR
	//
	// Será true se pelo menos um dos valores for true.
	fmt.Println(
		"true || false =",
		temIdadeMinima || temDocumento,
	)

	// NOT
	//
	// Inverte o valor booleano.
	fmt.Println(
		"!true =",
		!temIdadeMinima,
	)

	// Go não possui um operador específico para XOR lógico
	// entre valores booleanos.
	//
	// Porém, para dois bools, podemos utilizar != para representar
	// a ideia de "um é diferente do outro".
	fmt.Println(
		"true != false =",
		temIdadeMinima != temDocumento,
	)

	// O operador ^ também existe em Go,
	// mas com números inteiros ele está relacionado a operações bit a bit.
	// Esse conceito pode ser estudado separadamente.

	// =========================================================================
	// OPERADORES UNÁRIOS
	// =========================================================================

	// Um operador unário trabalha com apenas um valor.

	numeroPositivo := 10

	// O sinal de menos transforma o valor utilizado
	// naquela expressão em negativo.
	numeroNegativo := -numeroPositivo

	ativo := true

	fmt.Println("\n=== Operadores unários ===")

	fmt.Println("Valor original:", numeroPositivo)
	fmt.Println("-numeroPositivo:", numeroNegativo)

	// ! também é um operador unário,
	// pois trabalha apenas com um valor.
	fmt.Println("!ativo:", !ativo)

	// =========================================================================
	// ATRIBUIÇÕES COMPOSTAS
	// =========================================================================

	numero := 10

	fmt.Println("\n=== Atribuições compostas ===")
	fmt.Println("Valor inicial:", numero)

	// += soma um valor e armazena o resultado
	// novamente na mesma variável.
	//
	// numero += 15
	//
	// pode ser entendido como:
	//
	// numero = numero + 15
	numero += 15

	fmt.Println("Depois de numero += 15:", numero)

	// -= funciona da mesma maneira,
	// porém realizando uma subtração.
	numero -= 10

	fmt.Println("Depois de numero -= 10:", numero)

	// Também existem:
	//
	// *=
	// /=
	// %=

	// =========================================================================
	// INCREMENTO E DECREMENTO
	// =========================================================================

	// ++ aumenta o valor da variável em 1.
	numero++

	fmt.Println("Depois de numero++:", numero)

	// -- diminui o valor da variável em 1.
	numero--

	fmt.Println("Depois de numero--:", numero)

	// Existe um detalhe importante em Go:
	//
	// ++ e -- são instruções.
	//
	// Eles NÃO podem ser utilizados como parte de outra expressão.

	// Isto funciona:
	//
	// numero++

	// Isto NÃO funciona:
	//
	// outroNumero := numero++

	// Isto também NÃO funciona:
	//
	// fmt.Println(numero++)

	// Além disso, Go não possui a forma:
	//
	// ++numero
	//
	// Utilizamos:
	//
	// numero++

	// =========================================================================
	// OPERADOR TERNÁRIO
	// =========================================================================

	// Go NÃO possui operador ternário.
	//
	// Em algumas linguagens encontramos algo semelhante a:
	//
	// texto := numero > 5 ? "Maior que 5" : "Menor ou igual a 5"
	//
	// Essa sintaxe não existe em Go.

	// Quando precisamos escolher entre dois caminhos,
	// normalmente utilizamos if e else.

	valor := 15
	var texto string

	if valor > 5 {
		texto = "É maior que 5"
	} else {
		texto = "É menor ou igual a 5"
	}

	fmt.Println("\n=== Condicional no lugar do ternário ===")
	fmt.Println(texto)
}
