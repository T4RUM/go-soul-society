# 🔢 02 — Variáveis e Constantes em Go

Em Go, **variáveis** são utilizadas para armazenar valores que podem ser utilizados e alterados durante a execução do programa.

Já as **constantes** representam valores que, depois de definidos, não podem ser alterados.

Neste experimento, o objetivo é entender:

* como declarar variáveis com `var`;
* como informar explicitamente o tipo de uma variável;
* como o Go pode inferir tipos;
* como utilizar a declaração curta `:=`;
* como declarar várias variáveis em bloco;
* quais são os valores zero dos tipos básicos;
* como declarar constantes com `const`;
* a diferença entre variáveis e constantes;
* algumas convenções de nomes utilizadas em Go.

---

## 🧪 Estrutura do experimento

Neste exemplo, a aula está organizada da seguinte forma:

```text
go-soul-society/
│
├── 2-Variaveis/
│   │
│   ├── 02-variaveis.md
│   └── variaveis.go
│
├── go.mod
└── README.md
```

O arquivo:

```text
2-Variaveis/variaveis.go
```

contém os exemplos utilizados para estudar variáveis e constantes.

---

## 📦 O que é uma variável?

Uma variável é um espaço utilizado pelo programa para armazenar um valor.

Podemos imaginar assim:

```text
variável
   ↓
guarda um valor
   ↓
esse valor pode ser utilizado pelo programa
```

Por exemplo:

```go
var nome string = "Wesley"
```

Nesse caso:

```text
var
 ↓
declara uma variável

nome
 ↓
nome da variável

string
 ↓
tipo da variável

"Wesley"
 ↓
valor armazenado
```

Portanto, a variável `nome` armazena um valor do tipo `string`.

---

## 🧬 Declarando uma variável com `var`

Uma das formas de declarar variáveis em Go é utilizando:

```go
var
```

Exemplo:

```go
var nome string = "Wesley"
```

A estrutura pode ser lida como:

```text
var nome string = "Wesley"
│   │    │         │
│   │    │         └── valor
│   │    └──────────── tipo
│   └───────────────── nome
└───────────────────── declaração
```

Depois da declaração, podemos utilizar a variável normalmente:

```go
fmt.Println(nome)
```

Resultado:

```text
Wesley
```

---

## 🔤 Variável do tipo `string`

O tipo:

```go
string
```

é utilizado para armazenar textos.

Exemplo:

```go
var nome string = "Wesley"
```

Outro exemplo:

```go
var linguagem string = "Golang"
```

Strings são escritas entre aspas duplas:

```go
"texto"
```

---

## 🔢 Variável do tipo `int`

O tipo:

```go
int
```

é utilizado para representar números inteiros.

Exemplo:

```go
var idade int = 28
```

Nesse caso:

```text
idade
  ↓
28
  ↓
int
```

Como `28` não possui parte decimal, ele pode ser armazenado em uma variável inteira.

---

## ✅ Variável do tipo `bool`

O tipo:

```go
bool
```

representa valores lógicos.

Ele pode possuir apenas dois valores:

```text
true
false
```

No experimento:

```go
var estudando bool = true
```

Isso significa que a variável `estudando` possui o valor lógico:

```text
true
```

---

## 🖨️ Exibindo variáveis com `fmt.Println`

Podemos enviar variáveis para:

```go
fmt.Println()
```

Exemplo:

```go
fmt.Println(nome, idade, estudando)
```

Considerando:

```go
var nome string = "Wesley"
var idade int = 28
var estudando bool = true
```

a saída será:

```text
Wesley 28 true
```

O `fmt.Println` consegue receber vários valores separados por vírgula.

---

## ⚡ Declaração curta com `:=`

Go possui uma forma mais curta de declarar variáveis dentro de funções.

Em vez de escrever:

```go
var curso string = "Golang"
```

podemos escrever:

```go
curso := "Golang"
```

O operador:

```text
:=
```

declara a variável e permite que o Go descubra o tipo a partir do valor.

Nesse exemplo:

```go
curso := "Golang"
```

o compilador observa:

```text
"Golang"
   ↓
string
```

e entende que:

```text
curso → string
```

Não foi necessário escrever o tipo manualmente.

---

## 🧠 Inferência de tipo

Quando escrevemos:

```go
curso := "Golang"
```

o Go realiza a **inferência de tipo**.

Isso significa que o compilador determina o tipo da variável utilizando o valor atribuído.

Outros exemplos:

```go
nome := "Wesley"
idade := 28
estudando := true
altura := 1.75
```

Podemos imaginar:

```text
"Wesley" → string
28       → int
true     → bool
1.75     → float64
```

O tipo continua existindo.

A diferença é que não precisamos escrevê-lo explicitamente.

---

## ⚠️ `:=` é utilizado dentro de funções

A declaração curta:

```go
:=
```

é utilizada dentro de funções.

Por exemplo:

```go
func main() {
    curso := "Golang"
    fmt.Println(curso)
}
```

No nível do pacote, utilizamos declarações como:

```go
var linguagem = "Golang"
```

ou:

```go
const linguagem = "Golang"
```

Uma forma simples de lembrar:

```text
dentro de função
      ↓
:= pode ser utilizado

fora de função
      ↓
use var ou const
```

---

## 🔬 `var` e `:=`

As duas formas abaixo criam uma variável:

```go
var curso string = "Golang"
```

e:

```go
curso := "Golang"
```

Mas existem algumas diferenças importantes.

| Forma | Exemplo                       | Tipo explícito | Pode ser usada fora de funções |
| ----- | ----------------------------- | -------------- | ------------------------------ |
| `var` | `var curso string = "Golang"` | Pode ser       | Sim                            |
| `:=`  | `curso := "Golang"`           | Não            | Não                            |

Durante o desenvolvimento, a declaração curta aparece com bastante frequência quando o tipo já está claro pelo valor.

---

## 🧩 Declarando várias variáveis em bloco

Go permite agrupar várias declarações utilizando:

```go
var (
)
```

No experimento:

```go
var (
    documento_1 string = "Documento"
    documento_2 string = "Document"
)
```

Isso evita repetir:

```go
var
```

em todas as linhas.

Sem agrupamento, teríamos:

```go
var documento_1 string = "Documento"
var documento_2 string = "Document"
```

Com agrupamento:

```go
var (
    documento_1 string = "Documento"
    documento_2 string = "Document"
)
```

As duas formas são válidas.

O bloco apenas ajuda na organização quando várias variáveis relacionadas são declaradas juntas.

---

## 🏷️ Nomes de variáveis

Os nomes:

```go
documento_1
documento_2
```

funcionam em Go.

Porém, no estilo mais comum da linguagem, identificadores compostos costumam utilizar letras maiúsculas para separar palavras, em vez de `_`.

Por exemplo:

```go
documentoPrincipal
documentoTraduzido
```

No caso de nomes numerados, também podemos escrever:

```go
documento1
documento2
```

Por isso, no código organizado ao final deste experimento, utilizaremos:

```go
documento1
documento2
```

> O código original está correto. Essa alteração é apenas uma melhoria de estilo e legibilidade.

---

## 🔄 Alterando o valor de uma variável

Uma variável pode receber um novo valor depois de ter sido declarada.

Exemplo:

```go
var linguagem string = "Go"

linguagem = "Golang"
```

Depois da segunda atribuição:

```text
linguagem
    ↓
"Golang"
```

Como a variável já existe, utilizamos apenas:

```go
=
```

e não:

```go
:=
```

Neste caso:

```go
linguagem = "Golang"
```

significa:

```text
atribua um novo valor à variável existente
```

---

## 🕳️ Valores zero

Em Go, uma variável declarada sem receber um valor inicial recebe automaticamente o **valor zero** do seu tipo.

Exemplo:

```go
var nome string
var idade int
var estudando bool
var numero float64
```

Os valores serão:

| Tipo      | Valor zero |
| --------- | ---------- |
| `string`  | `""`       |
| `int`     | `0`        |
| `bool`    | `false`    |
| `float64` | `0`        |

Exemplo:

```go
var idade int

fmt.Println(idade)
```

Resultado:

```text
0
```

Isso significa que variáveis em Go não ficam com um valor indefinido quando são declaradas dessa forma.

---

## 🔒 Constantes

Uma constante é declarada utilizando:

```go
const
```

No experimento:

```go
const pi float64 = 3.1415926
```

Podemos dividir essa declaração assim:

```text
const pi float64 = 3.1415926
│     │  │         │
│     │  │         └── valor
│     │  └──────────── tipo
│     └─────────────── nome
└───────────────────── constante
```

A principal diferença é que o valor de uma constante não pode ser alterado depois da declaração.

---

## 🚫 Uma constante não pode ser reatribuída

Depois de escrever:

```go
const pi float64 = 3.1415926
```

não podemos fazer:

```go
pi = 3.14
```

Isso gera erro de compilação.

Podemos imaginar:

```text
variável
   ↓
pode receber outro valor

constante
   ↓
permanece com o mesmo valor
```

---

## ⚡ Constantes também podem ter o tipo inferido

Também é possível declarar:

```go
const pi = 3.1415926
```

Nesse caso, não informamos explicitamente:

```go
float64
```

No código da aula foi utilizada uma constante com tipo explícito:

```go
const pi float64 = 3.1415926
```

As duas formas são possíveis, mas possuem diferenças de tipagem que podem ser estudadas com mais profundidade posteriormente.

---

## ⚠️ Constantes não utilizam `:=`

Isto é válido:

```go
const pi = 3.1415926
```

Isto também é válido:

```go
const pi float64 = 3.1415926
```

Mas isto não declara uma constante:

```go
pi := 3.1415926
```

Nesse caso, `pi` seria uma **variável**.

Uma forma simples de lembrar:

```text
var ou :=
   ↓
variável

const
   ↓
constante
```

---

## 🧪 Variável vs constante

| Característica          | Variável      | Constante |
| ----------------------- | ------------- | --------- |
| Palavra-chave           | `var` ou `:=` | `const`   |
| Pode mudar de valor     | Sim           | Não       |
| Pode ter tipo explícito | Sim           | Sim       |
| Pode ter tipo inferido  | Sim           | Sim       |
| `:=` pode ser utilizado | Sim           | Não       |

Exemplo de variável:

```go
idade := 28
idade = 29
```

Exemplo de constante:

```go
const pi = 3.1415926
```

---

## 🧬 Tipos utilizados no experimento

Nesta aula apareceram quatro tipos básicos.

### `string`

Utilizado para textos:

```go
var nome string = "Wesley"
```

### `int`

Utilizado para números inteiros:

```go
var idade int = 28
```

### `bool`

Utilizado para valores lógicos:

```go
var estudando bool = true
```

### `float64`

Utilizado para números de ponto flutuante:

```go
const pi float64 = 3.1415926
```

Resumo:

```text
string  → texto
int     → número inteiro
bool    → verdadeiro ou falso
float64 → número com parte decimal
```

---

## ⚠️ Variáveis declaradas precisam ser utilizadas

Go é bastante rigoroso com código que não está sendo utilizado.

Dentro de uma função, se declararmos uma variável e nunca a utilizarmos:

```go
func main() {
    nome := "Wesley"
}
```

o compilador informa que a variável foi declarada e não utilizada.

Por isso, nos exemplos da aula utilizamos:

```go
fmt.Println()
```

para exibir os valores armazenados.

Essa característica ajuda a evitar código desnecessário dentro do programa.

---

## 🧠 O que aprendi

Neste experimento foi possível observar que:

* variáveis armazenam valores utilizados durante a execução do programa;
* `var` pode ser utilizado para declarar variáveis;
* o tipo de uma variável pode ser informado explicitamente;
* `string` representa textos;
* `int` representa números inteiros;
* `bool` representa valores lógicos;
* `float64` pode representar números com parte decimal;
* `:=` permite declarar variáveis de forma curta dentro de funções;
* com `:=`, o Go consegue inferir o tipo através do valor;
* várias variáveis podem ser agrupadas em um bloco `var`;
* variáveis declaradas sem valor recebem o valor zero do seu tipo;
* variáveis podem receber novos valores depois da declaração;
* constantes são declaradas com `const`;
* constantes não podem receber um novo valor depois de declaradas;
* `:=` cria uma variável, não uma constante;
* nomes com `_` são válidos, mas o estilo comum em Go prefere nomes como `documento1` ou `documentoPrincipal`;
* variáveis locais declaradas precisam ser utilizadas.

---

> Variáveis guardam o que pode mudar. Constantes dão nome ao que deve permanecer igual.

<p align="center">
  <img src="../docs/images/footer_02.jfif" alt="Go Soul Society">
</p>
