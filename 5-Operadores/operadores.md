# 🧮 05 — Operadores em Go

Em Go, **operadores** são símbolos utilizados para realizar operações com valores.

Eles aparecem em praticamente todo programa: para somar números, comparar valores, combinar condições, alterar variáveis e construir expressões.

Neste experimento, o objetivo é entender:

* o que são operadores;
* como funcionam os operadores aritméticos;
* a diferença entre divisão inteira e divisão com números de ponto flutuante;
* como funciona o operador de resto `%`;
* por que tipos numéricos diferentes não podem ser operados diretamente;
* como realizar conversões explícitas entre tipos compatíveis;
* a diferença entre declaração curta `:=` e atribuição `=`;
* como funcionam os operadores relacionais;
* como funcionam os operadores lógicos;
* como representar um comportamento semelhante ao XOR lógico;
* o que são operadores unários;
* como funcionam as atribuições compostas;
* como utilizar `++` e `--`;
* por que `++` e `--` não são expressões em Go;
* por que Go não possui operador ternário;
* como utilizar `if` e `else` quando precisamos escolher entre dois resultados.

---

## 🧪 Estrutura do experimento

Nesta aula, podemos organizar os arquivos da seguinte forma:

```text
go-soul-society/
│
├── 5-Operadores/
│   │
│   ├── 05-operadores.md
│   └── operadores.go
│
├── go.mod
└── README.md
```

O arquivo:

```text
5-Operadores/operadores.go
```

contém os exemplos utilizados para estudar operadores.

---

## 🧮 O que é um operador?

Um operador é um símbolo que informa ao programa que alguma operação deve ser realizada.

Por exemplo:

```go
2 + 2
```

O símbolo:

```text
+
```

é um operador.

Ele recebe os valores:

```text
2
2
```

e realiza uma operação entre eles.

Podemos imaginar:

```text
2 + 2
  ↑
operador
  ↓
realiza uma operação
  ↓
4
```

Os valores utilizados em uma operação são chamados de **operandos**.

No exemplo:

```go
2 + 2
```

temos:

```text
2  → operando
+  → operador
2  → operando
```

---

## ➕ Operadores aritméticos

Operadores aritméticos são utilizados para realizar cálculos numéricos.

Nesta aula utilizamos:

| Operador | Operação |
| --- | --- |
| `+` | soma |
| `-` | subtração |
| `*` | multiplicação |
| `/` | divisão |
| `%` | resto da divisão |

No código:

```go
soma := 2 + 2
subtracao := 2 - 5
multiplicacao := 2 * 4
divisao := 8 / 2
restoDaDivisao := 10 % 3
```

Cada expressão produz um valor que pode ser armazenado em uma variável.

---

## ➕ Soma

O operador:

```go
+
```

pode ser utilizado para somar valores numéricos.

Exemplo:

```go
soma := 2 + 2
```

Podemos visualizar:

```text
2 + 2
  ↓
4
  ↓
soma
```

Ao executar:

```go
fmt.Println("2 + 2 =", soma)
```

obtemos:

```text
2 + 2 = 4
```

---

## ➖ Subtração

O operador:

```go
-
```

realiza uma subtração.

Exemplo:

```go
subtracao := 2 - 5
```

Resultado:

```text
-3
```

Portanto:

```text
2 - 5
  ↓
-3
```

---

## ✖️ Multiplicação

O operador:

```go
*
```

realiza multiplicação.

Exemplo:

```go
multiplicacao := 2 * 4
```

Resultado:

```text
8
```

---

## ➗ Divisão

O operador:

```go
/
```

realiza uma divisão.

Exemplo:

```go
divisao := 8 / 2
```

Resultado:

```text
4
```

Existe, porém, um detalhe importante quando os dois operandos são inteiros.

---

## ⚠️ Divisão entre inteiros

Considere:

```go
divisaoInteira := 9 / 2
```

Matematicamente, poderíamos esperar:

```text
4.5
```

Mas os dois operandos são inteiros.

Por isso, a operação também produz um resultado inteiro:

```text
9 / 2
  ↓
4
```

A parte fracionária não é mantida nesse resultado.

Uma forma simples de lembrar:

```text
int / int
   ↓
resultado inteiro
```

Isso mostra novamente como os **tipos dos operandos influenciam o resultado de uma operação**.

---

## 🧩 Resto da divisão com `%`

O operador:

```go
%
```

retorna o resto de uma divisão inteira.

No experimento:

```go
restoDaDivisao := 10 % 3
```

Podemos decompor:

```text
10 ÷ 3

3 cabe 3 vezes em 10

3 × 3 = 9

10 - 9 = 1
```

Portanto:

```text
10 % 3
   ↓
1
```

---

## 🔬 Um uso comum de `%`

O resto da divisão é muito útil para descobrir se um número é divisível por outro.

Por exemplo:

```go
10 % 2
```

produz:

```text
0
```

Isso significa que `10` é divisível por `2` sem deixar resto.

Podemos imaginar:

```text
numero % 2 == 0
       ↓
resto igual a zero
       ↓
divisão exata por 2
```

Esse padrão será bastante útil quando estudarmos condições e repetições.

---

## 🧬 Operações dependem dos tipos

Na aula de tipos de dados vimos que:

```go
int16
```

e:

```go
int32
```

são tipos diferentes.

Mesmo que ambos representem números inteiros, Go não realiza automaticamente uma soma entre eles.

Considere:

```go
var numero16 int16 = 10
var numero32 int32 = 20
```

Isto não é válido:

```go
somaInvalida := numero16 + numero32
```

O problema não está nos valores.

Temos:

```text
10
20
```

O problema está nos tipos:

```text
numero16 → int16
numero32 → int32
```

Podemos visualizar:

```text
int16 + int32
      ↓
tipos diferentes
      ↓
operação não permitida diretamente
```

---

## 🔄 Conversão explícita

Para realizar a operação, podemos converter explicitamente um dos valores.

No experimento:

```go
somaConvertida := int32(numero16) + numero32
```

A expressão:

```go
int32(numero16)
```

produz um valor do tipo `int32` a partir do valor armazenado em `numero16`.

Agora a operação fica conceitualmente assim:

```text
int32(numero16) + numero32
       ↓              ↓
     int32          int32
          \          /
           \        /
             soma
              ↓
             30
```

---

## ⚠️ A conversão não altera a variável original

Quando escrevemos:

```go
int32(numero16)
```

isso não transforma permanentemente `numero16` em uma variável `int32`.

Ela continua sendo:

```text
numero16 → int16
```

A conversão produz um valor convertido para ser utilizado naquela expressão.

---

## 🏷️ Declaração e atribuição

Nesta aula aparecem três formas importantes:

```go
var
```

```go
:=
```

```go
=
```

Elas não significam a mesma coisa.

---

## 📦 Declaração com `var`

Podemos declarar uma variável utilizando:

```go
var nome string = "Wesley"
```

Nesse caso:

```text
var
 ↓
declara a variável

nome
 ↓
identificador

string
 ↓
tipo

"Wesley"
 ↓
valor
```

---

## ⚡ Declaração curta com `:=`

Dentro de funções, podemos utilizar:

```go
sobrenome := "Murat"
```

O símbolo:

```text
:=
```

é utilizado para uma **declaração curta**.

O compilador também infere o tipo a partir do valor:

```text
"Murat"
   ↓
string
   ↓
sobrenome → string
```

---

## 🔄 Atribuição com `=`

Depois que uma variável já existe, podemos atribuir um novo valor utilizando:

```go
=
```

Exemplo:

```go
linguagem := "Go"
```

Depois:

```go
linguagem = "Golang"
```

Podemos visualizar:

```text
linguagem := "Go"
      ↓
declaração

linguagem = "Golang"
      ↓
atribuição
```

---

## 🔬 `:=`, `=` e `==`

É importante não confundir esses símbolos.

| Símbolo | Significado |
| --- | --- |
| `:=` | declaração curta |
| `=` | atribuição |
| `==` | comparação de igualdade |

Podemos lembrar assim:

```text
:=  → cria uma variável
=   → atribui um valor
==  → compara dois valores
```

---

## ⚖️ Operadores relacionais

Operadores relacionais comparam valores.

O resultado dessas comparações é do tipo:

```go
bool
```

ou seja:

```text
true
```

ou:

```text
false
```

Nesta aula utilizamos:

| Operador | Significado |
| --- | --- |
| `>` | maior que |
| `>=` | maior ou igual |
| `<` | menor que |
| `<=` | menor ou igual |
| `==` | igual |
| `!=` | diferente |

---

## 🔎 Maior que `>`

Exemplo:

```go
1 > 2
```

A expressão pergunta:

```text
1 é maior que 2?
```

Resultado:

```text
false
```

---

## 🔎 Maior ou igual `>=`

Exemplo:

```go
1 >= 2
```

Resultado:

```text
false
```

O operador aceita duas possibilidades:

```text
maior
  OU
igual
```

---

## 🔎 Menor que `<`

Exemplo:

```go
1 < 2
```

Resultado:

```text
true
```

---

## 🔎 Menor ou igual `<=`

Exemplo:

```go
1 <= 2
```

Resultado:

```text
true
```

---

## 🟰 Igualdade com `==`

Para verificar se dois valores são iguais, utilizamos:

```go
==
```

Exemplo:

```go
1 == 2
```

Resultado:

```text
false
```

Novamente, não devemos confundir:

```text
=   → atribuição
==  → comparação
```

---

## 🚫 Diferente com `!=`

O operador:

```go
!=
```

verifica se dois valores são diferentes.

Exemplo:

```go
1 != 2
```

Resultado:

```text
true
```

---

## 🧠 Comparações produzem `bool`

Considere:

```go
numeroA := 1
numeroB := 2
```

A expressão:

```go
numeroA < numeroB
```

produz:

```text
true
```

Portanto, também poderíamos fazer:

```go
resultado := numeroA < numeroB
```

Nesse caso:

```text
resultado
    ↓
  true
    ↓
   bool
```

---

## 🧠 Operadores lógicos

Operadores lógicos são utilizados principalmente com valores booleanos.

Nesta aula utilizamos:

| Operador | Significado |
| --- | --- |
| `&&` | E lógico |
| `||` | OU lógico |
| `!` | NÃO lógico |

---

## 🔗 E lógico com `&&`

O operador:

```go
&&
```

representa o **E lógico**.

Para o resultado ser `true`, os dois lados precisam ser verdadeiros.

| A | B | `A && B` |
| --- | --- | --- |
| `true` | `true` | `true` |
| `true` | `false` | `false` |
| `false` | `true` | `false` |
| `false` | `false` | `false` |

No experimento:

```go
temIdadeMinima := true
temDocumento := false
```

Depois:

```go
temIdadeMinima && temDocumento
```

produz:

```text
false
```

---

## 🔀 OU lógico com `||`

O operador:

```go
||
```

representa o **OU lógico**.

Para o resultado ser `true`, pelo menos um dos lados precisa ser verdadeiro.

| A | B | `A || B` |
| --- | --- | --- |
| `true` | `true` | `true` |
| `true` | `false` | `true` |
| `false` | `true` | `true` |
| `false` | `false` | `false` |

No experimento:

```go
true || false
```

produz:

```text
true
```

---

## ❗ NÃO lógico com `!`

O operador:

```go
!
```

inverte um valor booleano.

Exemplo:

```go
!true
```

produz:

```text
false
```

Enquanto:

```go
!false
```

produz:

```text
true
```

Podemos visualizar:

```text
true
 ↓
 !
 ↓
false
```

---

## 🧩 XOR lógico

Go não possui um operador específico chamado `xor` para booleanos.

Quando queremos representar a ideia de que **os dois booleanos possuem valores diferentes**, podemos utilizar:

```go
!=
```

No experimento:

```go
true != false
```

produz:

```text
true
```

Para dois booleanos, essa comparação reproduz o comportamento esperado de um XOR lógico simples:

| A | B | `A != B` |
| --- | --- | --- |
| `true` | `true` | `false` |
| `true` | `false` | `true` |
| `false` | `true` | `true` |
| `false` | `false` | `false` |

---

## ⚠️ O operador `^` é outro conceito

Go possui o operador:

```go
^
```

Com dois números inteiros, ele representa um **XOR bit a bit**.

Isso é diferente de combinar diretamente valores `bool`.

Por enquanto, podemos separar os conceitos:

```text
booleanos diferentes
        ↓
       !=

XOR entre bits de inteiros
        ↓
        ^
```

Operações bit a bit podem ser estudadas em outro experimento.

---

## ⚡ Curto-circuito lógico

Os operadores:

```go
&&
```

e:

```go
||
```

possuem um comportamento chamado **curto-circuito**.

Com `&&`, se o lado esquerdo já for `false`, não é necessário avaliar o lado direito para descobrir o resultado final.

```text
false && algumaCondicao
  ↓
resultado já é false
```

Com `||`, se o lado esquerdo já for `true`, o resultado já será `true`.

```text
true || algumaCondicao
 ↓
resultado já é true
```

Esse comportamento será muito útil em condições mais complexas.

---

## ➖ Operadores unários

Um operador unário atua sobre apenas **um operando**.

No experimento:

```go
numeroPositivo := 10
numeroNegativo := -numeroPositivo
```

O operador:

```text
-
```

atua somente sobre `numeroPositivo`.

Podemos visualizar:

```text
numeroPositivo
      ↓
     10
      ↓
     -10
      ↓
numeroNegativo
```

---

## ❗ `!` também é unário

O operador lógico:

```go
!
```

também é um operador unário, porque atua sobre um único valor.

Exemplo:

```go
ativo := true
fmt.Println(!ativo)
```

Resultado:

```text
false
```

---

## 🔄 Atribuições compostas

Go possui operadores que combinam uma operação com uma atribuição.

Por exemplo:

```go
numero += 15
```

A ideia é equivalente a:

```go
numero = numero + 15
```

No experimento, começamos com:

```go
numero := 10
```

Depois:

```go
numero += 15
```

Resultado:

```text
25
```

---

## ➖ `-=`

Depois:

```go
numero -= 10
```

A ideia é:

```go
numero = numero - 10
```

Se `numero` valia `25`:

```text
25 - 10
   ↓
15
```

---

## 🧪 Outras atribuições compostas

Também podemos encontrar:

```go
numero *= 2
numero /= 2
numero %= 2
```

Resumo:

| Forma composta | Ideia equivalente |
| --- | --- |
| `x += y` | `x = x + y` |
| `x -= y` | `x = x - y` |
| `x *= y` | `x = x * y` |
| `x /= y` | `x = x / y` |
| `x %= y` | `x = x % y` |

---

## ⬆️ Incremento com `++`

Go possui:

```go
++
```

para incrementar uma variável em `1`.

Exemplo:

```go
numero++
```

Se:

```text
numero = 15
```

depois do incremento teremos:

```text
numero = 16
```

Podemos pensar:

```text
numero++
   ↓
aumenta numero em 1
```

---

## ⬇️ Decremento com `--`

O símbolo:

```go
--
```

reduz o valor em `1`.

Exemplo:

```go
numero--
```

Se `numero` vale `16`, depois teremos:

```text
15
```

---

## ⚠️ `++` e `--` não são expressões

Existe uma característica importante de Go.

Isto é válido:

```go
numero++
```

Mas isto não:

```go
outroNumero := numero++
```

Também não podemos fazer:

```go
fmt.Println(numero++)
```

Em Go, `++` e `--` são utilizados como instruções e não como expressões que produzem um valor.

Uma forma simples de lembrar:

```text
numero++
   ↓
altera numero

não retorna um valor
para outra expressão
```

---

## ⚠️ Não existe `++numero`

Em algumas linguagens podemos encontrar formas prefixadas e pós-fixadas.

Em Go, utilizamos:

```go
numero++
```

ou:

```go
numero--
```

Não utilizamos:

```go
++numero
```

nem:

```go
--numero
```

---

## 🔀 Go não possui operador ternário

Algumas linguagens possuem um operador ternário parecido com:

```text
condicao ? valor1 : valor2
```

Por exemplo:

```text
numero > 5 ? "Maior que 5" : "Menor ou igual a 5"
```

Essa sintaxe não existe em Go.

Portanto, isto não é válido:

```go
texto := numero > 5 ? "Maior que 5" : "Menor ou igual a 5"
```

---

## 🌿 Utilizando `if` e `else`

Quando precisamos escolher entre dois caminhos, podemos utilizar estruturas condicionais.

No experimento:

```go
valor := 15
var texto string

if valor > 5 {
    texto = "É maior que 5"
} else {
    texto = "É menor ou igual a 5"
}
```

A condição:

```go
valor > 5
```

produz um `bool`.

Como `valor` vale `15`:

```text
15 > 5
  ↓
true
```

Por isso é executado:

```go
texto = "É maior que 5"
```

---

## 🧠 Operadores dentro de condições

Observe como os conceitos começam a se conectar:

```go
if valor > 5 {
```

Temos:

```text
valor > 5
    ↓
operador relacional
    ↓
true ou false
    ↓
if decide qual bloco executar
```

Mais adiante, poderemos combinar várias condições usando operadores lógicos:

```go
if idade >= 18 && temDocumento {
    // ...
}
```

---

## 🧪 Precedência de operadores

Quando uma expressão possui vários operadores, existe uma ordem de avaliação.

Por exemplo:

```go
2 + 3 * 4
```

A multiplicação possui precedência sobre a soma.

Portanto:

```text
3 * 4
  ↓
12

2 + 12
   ↓
14
```

O resultado é:

```text
14
```

---

## 🧩 Alterando a ordem com parênteses

Podemos utilizar parênteses para deixar a ordem explícita.

Exemplo:

```go
(2 + 3) * 4
```

Agora:

```text
2 + 3
  ↓
5

5 * 4
  ↓
20
```

Os parênteses também podem melhorar a legibilidade de expressões mais complexas.

---

## 🔬 Resumo dos operadores estudados

| Categoria | Operadores principais | Resultado ou efeito |
| --- | --- | --- |
| Aritméticos | `+ - * / %` | produz valores numéricos |
| Relacionais | `> >= < <= == !=` | produz `bool` |
| Lógicos | `&& || !` | combina ou inverte booleanos |
| Atribuição | `=` | atribui valor a uma variável |
| Declaração curta | `:=` | declara variável dentro de função |
| Atribuição composta | `+= -= *= /= %=` | opera e atribui |
| Incremento | `++` | aumenta em `1` |
| Decremento | `--` | diminui em `1` |
| Unários | `+ - !` | atua sobre um único operando |

---

## 🧪 Experimento

Arquivo:

```text
5-Operadores/operadores.go
```

```go
package main

import "fmt"

func main() {
	// Operadores são símbolos que permitem realizar operações com valores.
	// Nesta aula veremos operadores aritméticos, relacionais, lógicos,
	// unários, de atribuição, incremento e decremento.

	// -------------------------------------------------------------------------
	// Operadores aritméticos
	// -------------------------------------------------------------------------

	// +  soma
	// -  subtração
	// *  multiplicação
	// /  divisão
	// %  resto da divisão inteira
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

	// Quando os dois operandos são inteiros, a divisão também produz
	// um resultado inteiro. A parte fracionária não é mantida.
	divisaoInteira := 9 / 2
	fmt.Println("9 / 2 =", divisaoInteira)

	// -------------------------------------------------------------------------
	// Operações entre tipos numéricos diferentes
	// -------------------------------------------------------------------------

	// Mesmo que int16 e int32 representem números inteiros, eles são
	// tipos diferentes em Go. Por isso, não podemos somá-los diretamente.
	var numero16 int16 = 10
	var numero32 int32 = 20

	// Isto não compila:
	//
	// somaInvalida := numero16 + numero32
	//
	// Para realizar a operação, precisamos converter explicitamente um
	// dos valores para um tipo compatível com o outro.
	somaConvertida := int32(numero16) + numero32

	fmt.Println("\n=== Tipos numéricos diferentes ===")
	fmt.Println("int32(numero16) + numero32 =", somaConvertida)

	// -------------------------------------------------------------------------
	// Declaração e atribuição
	// -------------------------------------------------------------------------

	// `var` declara uma variável e pode informar explicitamente seu tipo.
	var nome string = "Wesley"

	// `:=` é a declaração curta. Ela declara a variável e permite que
	// o compilador infira seu tipo a partir do valor.
	sobrenome := "Murat"

	fmt.Println("\n=== Declaração de variáveis ===")
	fmt.Println("Nome:", nome)
	fmt.Println("Sobrenome:", sobrenome)

	// O operador = não declara uma nova variável.
	// Ele atribui um novo valor a uma variável que já existe.
	linguagem := "Go"
	fmt.Println("Antes da atribuição:", linguagem)

	linguagem = "Golang"
	fmt.Println("Depois da atribuição:", linguagem)

	// -------------------------------------------------------------------------
	// Operadores relacionais
	// -------------------------------------------------------------------------

	// Operadores relacionais comparam dois valores e produzem
	// um resultado do tipo bool: true ou false.
	numeroA := 1
	numeroB := 2

	fmt.Println("\n=== Operadores relacionais ===")
	fmt.Println("1 > 2  =", numeroA > numeroB)
	fmt.Println("1 >= 2 =", numeroA >= numeroB)
	fmt.Println("1 < 2  =", numeroA < numeroB)
	fmt.Println("1 <= 2 =", numeroA <= numeroB)
	fmt.Println("1 == 2 =", numeroA == numeroB)
	fmt.Println("1 != 2 =", numeroA != numeroB)

	// -------------------------------------------------------------------------
	// Operadores lógicos
	// -------------------------------------------------------------------------

	// && representa E lógico.
	// || representa OU lógico.
	// !  representa NÃO lógico.
	temIdadeMinima := true
	temDocumento := false

	fmt.Println("\n=== Operadores lógicos ===")
	fmt.Println("true && false =", temIdadeMinima && temDocumento)
	fmt.Println("true || false =", temIdadeMinima || temDocumento)
	fmt.Println("!true         =", !temIdadeMinima)

	// Go não possui um operador dedicado para XOR lógico entre booleanos.
	// Quando queremos verificar se dois booleanos são diferentes,
	// `!=` pode representar esse comportamento.
	fmt.Println("true != false =", temIdadeMinima != temDocumento)

	// O operador ^ existe em Go, mas com números inteiros representa
	// XOR bit a bit, que é um conceito diferente.

	// -------------------------------------------------------------------------
	// Operadores unários
	// -------------------------------------------------------------------------

	// Um operador unário atua sobre apenas um operando.
	numeroPositivo := 10
	numeroNegativo := -numeroPositivo
	ativo := true

	fmt.Println("\n=== Operadores unários ===")
	fmt.Println("Valor original:", numeroPositivo)
	fmt.Println("-numeroPositivo:", numeroNegativo)
	fmt.Println("!ativo:", !ativo)

	// -------------------------------------------------------------------------
	// Atribuições compostas, incremento e decremento
	// -------------------------------------------------------------------------

	numero := 10

	fmt.Println("\n=== Atribuições compostas ===")
	fmt.Println("Valor inicial:", numero)

	// += soma e atribui o resultado à própria variável.
	numero += 15
	fmt.Println("Depois de numero += 15:", numero)

	// -= subtrai e atribui o resultado à própria variável.
	numero -= 10
	fmt.Println("Depois de numero -= 10:", numero)

	// ++ incrementa o valor em 1.
	numero++
	fmt.Println("Depois de numero++:", numero)

	// -- decrementa o valor em 1.
	numero--
	fmt.Println("Depois de numero--:", numero)

	// Em Go, ++ e -- são instruções e não expressões.
	// Por isso, algo como `outroNumero := numero++` não é válido.

	// Também existem outras atribuições compostas, como:
	//
	// numero *= 2
	// numero /= 2
	// numero %= 2

	// -------------------------------------------------------------------------
	// Go não possui operador ternário
	// -------------------------------------------------------------------------

	// Algumas linguagens possuem uma expressão parecida com:
	//
	// texto := numero > 5 ? "Maior que 5" : "Menor ou igual a 5"
	//
	// Essa sintaxe não existe em Go.
	// Em Go, usamos estruturas condicionais como if e else.

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
```

---

## ▶️ Executando o experimento

Dentro do diretório:

```text
5-Operadores/
```

podemos executar:

```bash
go run .
```

Ou, estando na raiz do repositório:

```bash
go run ./5-Operadores
```

A saída será:

```text
=== Operadores aritméticos ===
2 + 2 = 4
2 - 5 = -3
2 * 4 = 8
8 / 2 = 4
10 % 3 = 1
9 / 2 = 4

=== Tipos numéricos diferentes ===
int32(numero16) + numero32 = 30

=== Declaração de variáveis ===
Nome: Wesley
Sobrenome: Murat
Antes da atribuição: Go
Depois da atribuição: Golang

=== Operadores relacionais ===
1 > 2  = false
1 >= 2 = false
1 < 2  = true
1 <= 2 = true
1 == 2 = false
1 != 2 = true

=== Operadores lógicos ===
true && false = false
true || false = true
!true         = false
true != false = true

=== Operadores unários ===
Valor original: 10
-numeroPositivo: -10
!ativo: false

=== Atribuições compostas ===
Valor inicial: 10
Depois de numero += 15: 25
Depois de numero -= 10: 15
Depois de numero++: 16
Depois de numero--: 15

=== Condicional no lugar do ternário ===
É maior que 5
```

---

## 🔎 Entendendo a saída

A primeira parte demonstra os operadores aritméticos:

```text
2 + 2 = 4
2 - 5 = -3
2 * 4 = 8
8 / 2 = 4
10 % 3 = 1
```

Depois vemos:

```text
9 / 2 = 4
```

Isso acontece porque os dois operandos são inteiros.

A linha:

```text
int32(numero16) + numero32 = 30
```

mostra que a conversão explícita tornou os valores compatíveis para a soma.

Em seguida, vemos a diferença entre declaração e atribuição:

```text
Antes da atribuição: Go
Depois da atribuição: Golang
```

Os operadores relacionais produzem:

```text
false
false
true
true
false
true
```

Os operadores lógicos mostram o resultado de `&&`, `||`, `!` e da comparação `!=` entre booleanos.

Na parte de atribuições compostas, `numero` começa em `10` e passa pelas seguintes transformações:

```text
10
 ↓ +15
25
 ↓ -10
15
 ↓ ++
16
 ↓ --
15
```

Por fim, o `if` escolhe:

```text
É maior que 5
```

porque:

```text
15 > 5
  ↓
true
```

---

## 🧠 O que aprendi

Neste experimento foi possível observar que:

* operadores são símbolos utilizados para realizar operações com valores;
* os valores envolvidos em uma operação são chamados de operandos;
* `+`, `-`, `*`, `/` e `%` são operadores aritméticos;
* quando os dois operandos são inteiros, uma divisão também produz um resultado inteiro;
* `%` retorna o resto de uma divisão inteira;
* tipos numéricos diferentes, como `int16` e `int32`, não podem ser operados diretamente;
* podemos realizar uma conversão explícita para tornar valores compatíveis em uma expressão;
* a conversão utilizada em uma expressão não altera o tipo original da variável;
* `:=` é utilizado para declaração curta dentro de funções;
* `=` é utilizado para atribuição;
* `==` é utilizado para comparação de igualdade;
* `>`, `>=`, `<`, `<=`, `==` e `!=` são operadores relacionais;
* operações relacionais produzem valores do tipo `bool`;
* `&&` representa E lógico;
* `||` representa OU lógico;
* `!` representa NÃO lógico;
* `!=` entre dois booleanos pode representar o comportamento de um XOR lógico simples;
* `^` com inteiros representa XOR bit a bit, que é outro conceito;
* `&&` e `||` possuem comportamento de curto-circuito;
* operadores unários atuam sobre apenas um operando;
* `-numero` pode produzir o valor com sinal negativo;
* `!` também é um operador unário;
* atribuições compostas combinam uma operação com uma atribuição;
* `+=`, `-=`, `*=`, `/=` e `%=` são exemplos de atribuições compostas;
* `++` incrementa uma variável em `1`;
* `--` decrementa uma variável em `1`;
* `++` e `--` são instruções e não expressões em Go;
* Go utiliza apenas as formas pós-fixadas `numero++` e `numero--`;
* Go não possui operador ternário;
* `if` e `else` podem ser utilizados quando precisamos escolher entre caminhos diferentes;
* operadores possuem regras de precedência;
* parênteses podem alterar ou tornar explícita a ordem de avaliação de uma expressão.

---

> Operadores são as ferramentas que transformam valores em cálculos, comparações e decisões dentro do programa.

<p align="center">
  <img src="../docs/images/footer_05.jfif" alt="Go Soul Society">
</p>
