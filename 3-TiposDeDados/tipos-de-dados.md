# 🧬 03 — Tipos de Dados em Go

Em Go, **tipos de dados** determinam quais valores uma variável pode armazenar e quais operações podem ser realizadas com esses valores.

Ao declarar uma variável, o tipo informa ao compilador se aquele valor representa, por exemplo, um número inteiro, um número decimal, um texto, um valor lógico ou até mesmo um erro.

Neste experimento, o objetivo é entender:

* os principais tipos de números inteiros;
* a diferença entre `int8`, `int16`, `int32` e `int64`;
* como funciona o tipo `int`;
* o que são números inteiros sem sinal;
* a diferença entre `int` e `uint`;
* os aliases `rune` e `byte`;
* os tipos `float32` e `float64`;
* como Go representa textos com `string`;
* a diferença entre aspas simples e aspas duplas;
* como funciona o tipo `bool`;
* o que representa o tipo embutido `error`;
* o significado de `nil` em uma variável do tipo `error`;
* quais tipos são inferidos quando utilizamos `:=`.

---

## 🧪 Estrutura do experimento

Neste exemplo, a aula está organizada da seguinte forma:

```text
go-soul-society/
│
├── 3-Tipos-de-Dados/
│   │
│   ├── 03-tipos-de-dados.md
│   └── tipos-de-dados.go
│
├── go.mod
└── README.md
```

O arquivo:

```text
3-Tipos-de-Dados/tipos-de-dados.go
```

contém os exemplos utilizados para explorar alguns dos principais tipos disponíveis em Go.

---

## 🧬 O que é um tipo de dado?

Quando criamos uma variável, o Go precisa saber que tipo de valor será armazenado nela.

Por exemplo:

```go
var idade int = 28
```

Nesse caso:

```text
idade
  ↓
guarda o valor 28
  ↓
o valor pertence ao tipo int
```

Outro exemplo:

```go
var nome string = "Wesley"
```

Agora temos:

```text
nome
 ↓
guarda "Wesley"
 ↓
o valor pertence ao tipo string
```

O tipo ajuda o compilador a entender:

* quais valores podem ser armazenados;
* quanto espaço é necessário para representar determinados valores;
* quais operações fazem sentido;
* quando duas operações entre valores são incompatíveis.

Podemos imaginar:

```text
valor
  ↓
possui um tipo
  ↓
o tipo define como esse valor pode ser utilizado
```

---

## 🔢 Números inteiros

Números inteiros são valores que não possuem parte decimal.

Exemplos:

```text
-10
0
28
1000
10000000000
```

Go possui diferentes tipos para representar números inteiros com sinal:

```text
int8
int16
int32
int64
```

O número no nome do tipo indica quantos **bits** são utilizados para representar aquele valor.

Podemos visualizar assim:

```text
int8  → 8 bits
int16 → 16 bits
int32 → 32 bits
int64 → 64 bits
```

Quanto maior a quantidade de bits, maior pode ser o intervalo de valores representados.

---

## 📏 Intervalos dos inteiros com sinal

Os tipos inteiros possuem limites diferentes.

| Tipo    |                Menor valor |               Maior valor |
| ------- | -------------------------: | ------------------------: |
| `int8`  |                       -128 |                       127 |
| `int16` |                    -32.768 |                    32.767 |
| `int32` |             -2.147.483.648 |             2.147.483.647 |
| `int64` | -9.223.372.036.854.775.808 | 9.223.372.036.854.775.807 |

Por exemplo:

```go
var numero int8 = 100
```

funciona porque:

```text
100
 ↓
está entre -128 e 127
```

Porém:

```go
var numero int8 = 10000000000
```

não funciona.

O valor é muito maior do que aquilo que um `int8` consegue representar.

---

## ⚠️ Quando o valor não cabe no tipo

No experimento, queremos armazenar:

```text
10000000000
```

Esse valor não cabe em um:

```go
int8
```

Por isso utilizamos:

```go
var numero int64 = 10000000000
```

Agora o valor está dentro do intervalo suportado por `int64`.

Podemos imaginar:

```text
10000000000
     ↓
grande demais para int8
     ↓
cabe em int64
```

O compilador consegue detectar quando uma constante numérica utilizada na declaração ultrapassa o limite do tipo escolhido.

---

## 🔢 O tipo `int`

Além dos tipos com tamanho explícito, Go possui:

```go
int
```

Exemplo:

```go
var numeroDois int = 1000000000
```

Diferente de:

```text
int8
int16
int32
int64
```

o tamanho de `int` depende da arquitetura para a qual o programa está sendo compilado.

Na prática, ele terá tamanho de:

```text
32 bits
```

ou:

```text
64 bits
```

Por isso, não devemos pensar em `int` como sendo simplesmente outro nome para `int64`.

São tipos diferentes.

Uma forma simples de lembrar:

```text
int8 / int16 / int32 / int64
            ↓
tamanho explicitamente definido

int
 ↓
tamanho natural da arquitetura
```

Para números inteiros comuns, `int` é utilizado com bastante frequência.

Quando o tamanho exato importa, tipos como `int32` e `int64` podem ser escolhidos explicitamente.

---

## ⚠️ Código dependente da arquitetura

Um valor como:

```go
var numeroDois int = 100000000000
```

funciona quando `int` possui 64 bits, mas não cabe em um `int` de 32 bits.

Para que o exemplo seja portátil entre arquiteturas diferentes, podemos utilizar um valor menor:

```go
var numeroDois int = 1000000000
```

Se realmente precisarmos garantir que um inteiro grande seja suportado, podemos escolher explicitamente:

```go
var numeroDois int64 = 100000000000
```

Assim, o tamanho esperado pelo programa fica claro.

---

## ⚡ Inferência de tipo com números inteiros

Na aula anterior vimos que:

```go
:=
```

permite declarar uma variável utilizando inferência de tipo.

No experimento:

```go
numeroTres := 10000000
```

Como estamos utilizando um literal inteiro sem informar o tipo explicitamente, o tipo padrão escolhido será:

```text
int
```

Podemos imaginar:

```text
10000000
   ↓
literal inteiro
   ↓
tipo padrão: int
```

Portanto:

```go
numeroTres := 10000000
```

é semelhante, em relação ao tipo resultante, a:

```go
var numeroTres int = 10000000
```

---

## ➕➖ Inteiros com sinal

Os tipos:

```text
int
int8
int16
int32
int64
```

representam números inteiros **com sinal**.

Isso significa que podem representar valores:

```text
negativos
zero
positivos
```

Exemplo:

```go
var temperatura int = -5
```

Nesse caso, o sinal:

```text
-
```

faz parte do valor representado.

---

## 🚫 Inteiros sem sinal

Go também possui tipos inteiros **sem sinal**:

```text
uint
uint8
uint16
uint32
uint64
```

A palavra:

```text
uint
```

vem de:

```text
unsigned integer
```

ou seja:

```text
inteiro sem sinal
```

Esses tipos representam apenas valores iguais ou maiores que zero.

Exemplo:

```go
var numeroQuatro uint = 1000
```

Isso funciona.

Porém, isto não:

```go
var numeroQuatro uint = -1000
```

Um `uint` não representa números negativos.

Podemos resumir:

```text
int
 ↓
negativo, zero ou positivo

uint
  ↓
zero ou positivo
```

---

## 📏 Inteiros sem sinal e capacidade

Como os tipos sem sinal não precisam representar valores negativos, seu intervalo começa em zero.

Por exemplo:

```text
uint8
```

pode representar:

```text
0 até 255
```

Enquanto:

```text
int8
```

representa:

```text
-128 até 127
```

Isso não significa que `uint` deve ser utilizado sempre que sabemos que um valor será positivo.

A escolha do tipo depende do significado do dado e do contexto do programa.

---

## 🧬 `rune`

Go possui o identificador:

```go
rune
```

`rune` é um **alias de `int32`**.

Isso significa que:

```text
rune == int32
```

em relação ao tipo representado.

No experimento:

```go
var numeroCinco rune = 123456
```

O valor é armazenado utilizando a mesma representação de um `int32`.

Porém, o nome `rune` normalmente é utilizado quando queremos expressar a ideia de um **code point Unicode**, ou seja, um valor que identifica um caractere no padrão Unicode.

Exemplo:

```go
var letra rune = 'B'
```

Aqui:

```text
'B'
 ↓
rune
 ↓
code point Unicode
```

O nome `rune` ajuda a comunicar a intenção do código.

---

## 🔤 Aspas simples representam uma `rune`

No experimento original apareceu:

```go
char := 'B'
```

Em Go, aspas simples:

```text
' '
```

representam um literal de `rune`.

Portanto:

```go
caractere := 'B'
```

faz com que `caractere` tenha como tipo padrão:

```text
rune
```

Como `rune` é um alias para:

```text
int32
```

ao fazer:

```go
fmt.Println(caractere)
```

o resultado é o valor numérico do code point:

```text
66
```

Podemos exibir o caractere utilizando:

```go
fmt.Printf("%c\n", caractere)
```

Resultado:

```text
B
```

Portanto:

```text
fmt.Println(caractere) → 66
fmt.Printf("%c", caractere) → B
```

O caractere `B` também existe na tabela ASCII com o valor `66`, mas a forma mais correta de pensar em uma `rune` em Go é como um **code point Unicode**, e não apenas como um caractere ASCII.

---

## 🧱 `byte`

Outro alias importante em Go é:

```go
byte
```

`byte` é um alias para:

```go
uint8
```

Portanto:

```text
byte == uint8
```

No experimento:

```go
var numeroSeis byte = 8
```

Isso é equivalente, em relação ao tipo, a:

```go
var numeroSeis uint8 = 8
```

O nome `byte` costuma ser utilizado quando o valor representa um byte de dados.

Por exemplo, futuramente será comum encontrar:

```go
[]byte
```

que representa uma sequência de bytes.

Resumo:

```text
rune → alias de int32
byte → alias de uint8
```

---

## 🔬 `rune` e `byte`

Apesar de ambos serem aliases de tipos inteiros, normalmente comunicam intenções diferentes.

```text
rune
 ↓
code point Unicode
 ↓
int32

byte
 ↓
byte de dados
 ↓
uint8
```

Isso é um exemplo de como o nome utilizado no código também pode ajudar a explicar o significado daquele valor.

---

## 🌊 Números de ponto flutuante

Para representar números que possuem parte decimal, Go possui:

```text
float32
float64
```

Exemplos:

```go
var numeroSete float32 = 0.1
var numeroOito float64 = 0.22222222
```

Podemos visualizar:

```text
float32 → ponto flutuante de 32 bits
float64 → ponto flutuante de 64 bits
```

Normalmente, `float64` oferece maior precisão do que `float32`.

---

## 🎯 `float32` e `float64`

No experimento:

```go
var numeroSete float32 = 0.1
```

e:

```go
var numeroOito float64 = 0.22222222
```

As duas variáveis armazenam números de ponto flutuante, mas utilizam precisões diferentes.

Uma forma simples de pensar neste momento é:

```text
float32
   ↓
menor precisão

float64
   ↓
maior precisão
```

Isso não significa que `float64` consiga representar perfeitamente qualquer número decimal.

Números de ponto flutuante são armazenados em formato binário e alguns valores decimais possuem apenas uma representação aproximada.

Esse detalhe se torna especialmente importante em situações que exigem precisão específica, como determinados cálculos financeiros.

---

## ⚡ Inferência de tipo com números decimais

No experimento:

```go
numeroNove := 12313.54
```

Como o valor possui parte decimal e nenhum tipo foi informado explicitamente, o tipo padrão inferido é:

```text
float64
```

Portanto:

```text
12313.54
    ↓
literal de ponto flutuante
    ↓
tipo padrão: float64
```

Isso significa que:

```go
numeroNove := 12313.54
```

resulta em uma variável do tipo:

```go
float64
```

---

## 📝 Strings

O tipo:

```go
string
```

é utilizado para representar textos.

No experimento:

```go
var nome string = "Wesley"
```

Outro exemplo:

```go
curso := "Golang"
```

Nos dois casos temos valores do tipo:

```text
string
```

Strings normalmente são escritas utilizando aspas duplas:

```go
"texto"
```

Exemplo:

```go
mensagem := "Estudando Go"
```

---

## 🔤 Aspas duplas vs aspas simples

Essa diferença é importante em Go.

Aspas duplas:

```go
"B"
```

representam uma:

```text
string
```

Aspas simples:

```go
'B'
```

representam uma:

```text
rune
```

Portanto:

```go
texto := "B"
caractere := 'B'
```

não criam valores do mesmo tipo.

Podemos visualizar:

```text
"B"
 ↓
string

'B'
 ↓
rune
 ↓
int32
```

Essa diferença explica por que:

```go
fmt.Println("B")
```

imprime:

```text
B
```

enquanto:

```go
fmt.Println('B')
```

imprime:

```text
66
```

---

## 🌎 Strings e Unicode

Strings em Go armazenam uma sequência de bytes.

Na prática, textos em Go são frequentemente codificados utilizando UTF-8.

Isso será importante quando estudarmos operações mais avançadas com strings, porque:

```text
byte
```

e:

```text
rune
```

não representam exatamente a mesma ideia.

Por enquanto, podemos manter a seguinte visão:

```text
string
  ↓
texto

byte
  ↓
unidade de dados de 8 bits

rune
  ↓
code point Unicode
```

---

## ✅ Booleanos

O tipo:

```go
bool
```

representa valores lógicos.

Existem apenas dois valores possíveis:

```text
true
false
```

No experimento:

```go
var booleanoUm bool = false
```

Podemos imaginar:

```text
booleanoUm
    ↓
false
    ↓
bool
```

Valores booleanos aparecem com muita frequência em condições.

Por exemplo, futuramente poderemos ter algo como:

```go
estaLogado := true
```

ou:

```go
temPermissao := false
```

Esses valores poderão ser utilizados em estruturas como:

```text
if
```

que será estudada em outro experimento.

---

## ⚠️ O tipo `error`

Go possui um tipo embutido chamado:

```go
error
```

Ele é extremamente comum na linguagem.

No experimento:

```go
var erro error
```

Existe um detalhe importante:

> `error` não é um tipo numérico ou textual básico como `int`, `string` ou `bool`.

`error` é uma **interface embutida da linguagem** utilizada para representar condições de erro.

Sua definição conceitualmente exige que um valor forneça uma mensagem através de um método chamado:

```go
Error() string
```

Não precisamos dominar interfaces ainda para utilizar erros.

Por enquanto, podemos pensar:

```text
error
  ↓
forma padrão de representar erros em Go
```

---

## 🕳️ Valor zero de `error`

Quando declaramos:

```go
var erro error
```

sem atribuir nenhum erro, seu valor é:

```go
nil
```

Ao executar:

```go
fmt.Println(erro)
```

vemos:

```text
<nil>
```

Isso indica que, naquele momento, a variável não contém um valor de erro concreto.

Podemos imaginar:

```text
erro
 ↓
nil
 ↓
nenhum erro armazenado
```

Em código Go, é muito comum encontrar verificações parecidas com:

```go
if err != nil {
    // ocorreu algum erro
}
```

Esse padrão será explorado com mais profundidade futuramente.

---

## 🧠 `nil` não é uma string

Quando:

```go
fmt.Println(erro)
```

mostra:

```text
<nil>
```

isso não significa que a variável contém o texto:

```go
"<nil>"
```

`nil` representa a ausência de um valor para determinados tipos.

Ele pode aparecer, por exemplo, com interfaces, ponteiros, slices, maps, channels e funções.

Nem todo tipo em Go pode receber `nil`.

Por exemplo, isto não é válido:

```go
var numero int = nil
```

Um `int` possui como valor zero:

```text
0
```

e não `nil`.

---

## 🚨 Criando um erro com `errors.New`

Para criar um valor de erro simples, podemos utilizar o pacote padrão:

```go
errors
```

Por isso o experimento importa:

```go
import "errors"
```

Depois podemos fazer:

```go
var erroInterno error = errors.New("erro interno")
```

Agora:

```text
erroInterno
     ↓
contém um erro
     ↓
"erro interno"
```

Ao executar:

```go
fmt.Println(erroInterno)
```

a saída será:

```text
erro interno
```

---

## 📦 Importando `errors`

O início do programa fica:

```go
import (
    "errors"
    "fmt"
)
```

O pacote:

```text
fmt
```

é utilizado para exibir valores.

O pacote:

```text
errors
```

é utilizado neste exemplo para criar um novo erro.

Portanto:

```text
fmt
 ↓
saída no terminal

errors
  ↓
criação de um valor de erro
```

---

## 🔍 Tipos inferidos com `:=`

Ao utilizar:

```go
:=
```

o compilador escolhe o tipo a partir do valor informado.

Nos exemplos desta aula:

```go
numeroTres := 10000000
numeroNove := 12313.54
curso := "Golang"
caractere := 'B'
```

temos:

| Declaração               | Tipo resultante  |
| ------------------------ | ---------------- |
| `numeroTres := 10000000` | `int`            |
| `numeroNove := 12313.54` | `float64`        |
| `curso := "Golang"`      | `string`         |
| `caractere := 'B'`       | `rune` (`int32`) |

Podemos visualizar:

```text
10000000  → int
12313.54  → float64
"Golang"  → string
'B'       → rune
```

A inferência evita que precisemos escrever o tipo quando ele já pode ser determinado pelo valor.

---

## 🔬 Descobrindo o tipo de uma variável

Durante os estudos, uma forma útil de confirmar o tipo de um valor é utilizar:

```go
fmt.Printf("%T\n", valor)
```

Por exemplo:

```go
numero := 10
texto := "Go"
decimal := 1.5
caractere := 'G'

fmt.Printf("%T\n", numero)
fmt.Printf("%T\n", texto)
fmt.Printf("%T\n", decimal)
fmt.Printf("%T\n", caractere)
```

A saída será semelhante a:

```text
int
string
float64
int32
```

Como `rune` é um alias para `int32`, `%T` apresenta:

```text
int32
```

para esse valor.

Esse recurso pode ser útil para experimentar a inferência de tipos.

---

## 🧩 Tipos diferentes continuam sendo diferentes

Mesmo quando dois tipos conseguem representar valores parecidos, Go normalmente não realiza conversões numéricas implícitas entre eles.

Por exemplo:

```go
var numeroA int32 = 10
var numeroB int64 = 20
```

Não podemos simplesmente tratar os dois como se fossem exatamente o mesmo tipo.

Quando uma conversão é necessária, ela deve ser explícita.

Exemplo:

```go
resultado := int64(numeroA) + numeroB
```

Nesse caso:

```go
int64(numeroA)
```

converte o valor de `numeroA` para `int64`.

Essa característica ajuda a deixar claro no código quando estamos mudando a representação de um valor.

---

## 🕳️ Relembrando os valores zero

Na aula anterior vimos que variáveis declaradas sem valor recebem automaticamente o **valor zero** de seu tipo.

Alguns exemplos relacionados a esta aula:

| Tipo      | Valor zero |
| --------- | ---------- |
| `int`     | `0`        |
| `int8`    | `0`        |
| `int64`   | `0`        |
| `uint`    | `0`        |
| `float32` | `0`        |
| `float64` | `0`        |
| `string`  | `""`       |
| `bool`    | `false`    |
| `error`   | `nil`      |

Por exemplo:

```go
var numero int
var texto string
var ativo bool
var erro error
```

Os valores iniciais serão:

```text
numero → 0
texto  → ""
ativo  → false
erro   → nil
```

---

## 🧪 Experimento

Arquivo:

```text
3-Tipos-de-Dados/tipos-de-dados.go
```

```go
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
```

---

## ▶️ Executando o experimento

Dentro do diretório:

```text
3-Tipos-de-Dados/
```

podemos executar:

```bash
go run .
```

Ou, estando na raiz do repositório:

```bash
go run ./3-Tipos-de-Dados
```

A saída será semelhante a:

```text
10000000000
1000000000
10000000
1000
123456
8
0.1
0.22222222
12313.54
Wesley
Golang
66
B
false
<nil>
erro interno
```

Algumas linhas merecem atenção especial:

```text
66
B
```

As duas foram produzidas a partir da mesma `rune`:

```go
caractere := 'B'
```

Com:

```go
fmt.Println(caractere)
```

vemos seu valor numérico.

Com:

```go
fmt.Printf("%c\n", caractere)
```

vemos o caractere representado por aquele code point.

Também podemos observar:

```text
<nil>
```

como o valor zero da variável:

```go
var erro error
```

---

## 🧪 Resumo dos tipos estudados

| Tipo      | Utilização neste experimento                                 |
| --------- | ------------------------------------------------------------ |
| `int8`    | Inteiro com sinal de 8 bits                                  |
| `int16`   | Inteiro com sinal de 16 bits                                 |
| `int32`   | Inteiro com sinal de 32 bits                                 |
| `int64`   | Inteiro com sinal de 64 bits                                 |
| `int`     | Inteiro com tamanho natural da arquitetura                   |
| `uint`    | Inteiro sem sinal                                            |
| `rune`    | Alias de `int32`, normalmente usado para code points Unicode |
| `byte`    | Alias de `uint8`                                             |
| `float32` | Número de ponto flutuante de 32 bits                         |
| `float64` | Número de ponto flutuante de 64 bits                         |
| `string`  | Texto                                                        |
| `bool`    | Valor lógico: `true` ou `false`                              |
| `error`   | Interface embutida utilizada para representar erros          |

---

## 🧠 O que aprendi

Neste experimento foi possível observar que:

* Go possui diferentes tipos para representar números inteiros;
* `int8`, `int16`, `int32` e `int64` possuem tamanhos e limites diferentes;
* um valor precisa caber no intervalo suportado pelo tipo escolhido;
* `int` possui tamanho dependente da arquitetura e não deve ser tratado simplesmente como um alias de `int64`;
* ao utilizar `:=` com um literal inteiro, o tipo padrão normalmente será `int`;
* `uint` representa números inteiros sem sinal e não aceita valores negativos;
* existem também `uint8`, `uint16`, `uint32` e `uint64`;
* `rune` é um alias de `int32`;
* `rune` normalmente representa um code point Unicode;
* `byte` é um alias de `uint8`;
* `float32` e `float64` representam números de ponto flutuante;
* `float64` oferece maior precisão do que `float32`;
* um literal decimal utilizado com `:=` recebe `float64` como tipo padrão;
* `string` representa textos;
* aspas duplas criam strings;
* aspas simples representam uma `rune`;
* ao imprimir uma `rune` com `fmt.Println`, vemos seu valor numérico;
* `%c` pode ser utilizado com `fmt.Printf` para exibir o caractere representado por uma `rune`;
* `bool` pode possuir os valores `true` ou `false`;
* `error` é uma interface embutida e muito utilizada em Go;
* uma variável do tipo `error` sem valor possui valor zero `nil`;
* `errors.New` pode ser utilizado para criar um erro simples;
* `nil` não é o valor zero de todos os tipos;
* tipos numéricos diferentes não são convertidos automaticamente entre si;
* `fmt.Printf("%T", valor)` pode ajudar a descobrir o tipo de um valor durante os estudos.

---

> Entender os tipos é entender quais valores o programa consegue representar e como o Go espera que cada um deles seja utilizado.

<p align="center">
  <img src="../docs/images/footer_03.jfif" alt="Go Soul Society">
</p>
