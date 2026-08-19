# ⚙️ 04 — Funções em Go

Em Go, **funções** permitem organizar uma sequência de instruções em blocos de código que podem ser chamados sempre que forem necessários.

Uma função pode receber valores, executar alguma operação e também devolver um ou mais resultados para quem realizou a chamada.

Neste experimento, o objetivo é entender:

* como declarar uma função utilizando `func`;
* como chamar uma função;
* o que são parâmetros e argumentos;
* como declarar parâmetros do mesmo tipo de forma agrupada;
* como retornar um valor utilizando `return`;
* como indicar o tipo de retorno de uma função;
* como armazenar o resultado de uma função em uma variável;
* como uma função pode ser armazenada em uma variável;
* o que é uma função anônima;
* como identificar o tipo `func()`;
* como uma função pode retornar vários valores;
* como receber vários retornos de uma única chamada;
* como utilizar o identificador em branco `_` para ignorar um retorno.

---

## 🧪 Estrutura do experimento

Neste exemplo, a aula está organizada da seguinte forma:

```text
go-soul-society/
│
├── 4-Funcoes/
│   │
│   ├── funcoes.md
│   └── funcoes.go
│
├── go.mod
└── README.md
```

O arquivo:

```text
4-Funcoes/funcoes.go
```

contém os exemplos utilizados para estudar funções.

---

## ⚙️ O que é uma função?

Uma função é um bloco de código que possui uma responsabilidade específica e pode ser executado quando for chamado.

Podemos imaginar:

```text
função
  ↓
agrupa uma tarefa
  ↓
pode receber valores
  ↓
executa instruções
  ↓
pode devolver resultados
```

Por exemplo:

```go
func somar(num1, num2 int) int {
    return num1 + num2
}
```

Essa função recebe dois números inteiros e devolve a soma entre eles.

Em outro ponto do programa, podemos chamá-la:

```go
resultado := somar(15, 10)
```

Podemos visualizar o fluxo assim:

```text
somar(15, 10)
      ↓
executa a função somar
      ↓
15 + 10
      ↓
25
      ↓
resultado
```

---

## 🧬 Declarando uma função com `func`

Uma função declarada em Go começa com a palavra-chave:

```go
func
```

No experimento:

```go
func somar(num1, num2 int) int {
    return num1 + num2
}
```

Podemos separar essa declaração em partes:

```text
func somar(num1, num2 int) int {
│    │     │              │
│    │     │              └── tipo retornado
│    │     └───────────────── parâmetros
│    └─────────────────────── nome da função
└──────────────────────────── declaração de função
```

O corpo da função fica entre:

```text
{
}
```

Neste caso, o corpo contém:

```go
return num1 + num2
```

---

## 🏷️ Nome da função

Depois de `func`, informamos o nome da função.

No exemplo:

```go
func somar(num1, num2 int) int {
    return num1 + num2
}
```

O nome é:

```text
somar
```

É esse nome que utilizamos para executar a função:

```go
somar(15, 10)
```

Podemos imaginar:

```text
func somar(...)
     ↓
declara a função

somar(...)
   ↓
chama a função
```

---

## 📥 Parâmetros

Uma função pode receber valores para utilizar durante sua execução.

Na declaração:

```go
func somar(num1, num2 int) int {
    return num1 + num2
}
```

os parâmetros são:

```text
num1
num2
```

Eles existem dentro da função e representam os valores recebidos quando ela é chamada.

Podemos imaginar:

```text
func somar(num1, num2 int)
           │     │
           │     └── segundo parâmetro
           └──────── primeiro parâmetro
```

Quando a chamada é:

```go
somar(15, 10)
```

os valores chegam aos parâmetros:

```text
15 → num1
10 → num2
```

Durante essa execução, é como se tivéssemos:

```text
num1 = 15
num2 = 10
```

---

## 🎯 Parâmetros e argumentos

É útil diferenciar dois nomes.

Na declaração da função:

```go
func somar(num1, num2 int) int {
    return num1 + num2
}
```

`num1` e `num2` são **parâmetros**.

Na chamada:

```go
somar(15, 10)
```

`15` e `10` são **argumentos**.

Podemos resumir:

```text
func somar(num1, num2 int)
           ↑     ↑
        parâmetros

somar(15, 10)
      ↑   ↑
   argumentos
```

Os parâmetros fazem parte da definição da função.

Os argumentos são os valores fornecidos em uma chamada específica.

---

## 🧩 Parâmetros do mesmo tipo

No experimento temos:

```go
func somar(num1, num2 int) int {
    return num1 + num2
}
```

Como `num1` e `num2` possuem o mesmo tipo, podemos escrever:

```go
num1, num2 int
```

em vez de repetir:

```go
num1 int, num2 int
```

As duas formas expressam que ambos os parâmetros são do tipo:

```text
int
```

No código da aula também aparece:

```go
func calculadora(n1, n2 int) (int, int, int, int) {
```

e:

```go
func nomeCompleto(nome, sobrenome string) (string, string) {
```

Portanto:

```text
n1, n2 int
      ↓
os dois são int

nome, sobrenome string
                ↓
os dois são string
```

---

## 📤 Retornando um valor

Uma função pode devolver um resultado para o código que realizou a chamada.

No experimento:

```go
func somar(num1, num2 int) int {
    return num1 + num2
}
```

A palavra-chave:

```go
return
```

indica o valor que será devolvido pela função.

Neste caso:

```go
return num1 + num2
```

significa aproximadamente:

```text
calcule num1 + num2
        ↓
devolva o resultado
```

Para a chamada:

```go
somar(15, 10)
```

teremos:

```text
15 + 10
   ↓
25
   ↓
return
```

---

## 🧬 Tipo de retorno

Quando uma função devolve um valor, sua declaração informa qual será o tipo retornado.

Observe:

```go
func somar(num1, num2 int) int {
    return num1 + num2
}
```

Depois da lista de parâmetros aparece:

```text
int
```

Esse `int` indica o tipo do valor retornado.

Podemos visualizar:

```text
func somar(num1, num2 int) int
                            ↑
                       tipo de retorno
```

Como a função devolve:

```go
num1 + num2
```

e os dois valores são inteiros, o retorno dessa função é um `int`.

---

## 📦 Armazenando o retorno em uma variável

Na função `main` encontramos:

```go
resultado := somar(15, 10)
```

A chamada:

```go
somar(15, 10)
```

produz um valor.

Esse valor é utilizado na declaração curta:

```go
resultado := ...
```

Podemos imaginar:

```text
somar(15, 10)
      ↓
25
      ↓
resultado := 25
```

Depois podemos utilizar a variável normalmente:

```go
fmt.Println("Soma:", resultado)
```

A saída será:

```text
Soma: 25
```

---

## 🔁 Fluxo de uma chamada de função

A sequência completa pode ser visualizada assim:

```text
main()
  ↓
resultado := somar(15, 10)
                  ↓
        entra em somar
                  ↓
          num1 = 15
          num2 = 10
                  ↓
          num1 + num2
                  ↓
               25
                  ↓
              return
                  ↓
resultado recebe 25
                  ↓
fmt.Println("Soma:", resultado)
```

A função concentra a operação em um local e a chamada utiliza o resultado produzido por ela.

---

## 📍 A função pode ser declarada depois de `main`

No arquivo da aula, a função:

```go
func somar(num1, num2 int) int {
    return num1 + num2
}
```

aparece depois de:

```go
func main() {
    // ...
}
```

Mesmo assim, dentro de `main` podemos escrever:

```go
resultado := somar(15, 10)
```

Não é necessário colocar `somar` fisicamente antes de `main` apenas para que essa chamada funcione.

Isso permite organizar as funções do arquivo de uma forma que faça sentido para o código.

No experimento, temos aproximadamente:

```text
funcoes.go
│
├── func main()
├── func somar(...)
├── var minhaFuncao = func() {...}
├── func calculadora(...)
└── func nomeCompleto(...)
```

---

## 🚫 Função sem parâmetros

Nem toda função precisa receber valores.

No experimento, a função armazenada em `minhaFuncao` é criada assim:

```go
var minhaFuncao = func() {
    fmt.Println("Essa função está armazenada em uma variável.")
}
```

Observe os parênteses:

```go
func()
```

Eles estão vazios.

Isso indica que essa função não possui parâmetros.

Podemos visualizar:

```text
func()
     ↑
nenhum parâmetro
```

A chamada também não precisa fornecer argumentos:

```go
minhaFuncao()
```

---

## 📭 Função sem retorno

A mesma função também não declara um tipo de retorno:

```go
var minhaFuncao = func() {
    fmt.Println("Essa função está armazenada em uma variável.")
}
```

Depois dos parênteses não aparece algo como:

```go
int
```

ou:

```go
string
```

Ela apenas executa:

```go
fmt.Println("Essa função está armazenada em uma variável.")
```

Podemos resumir:

```text
func()
  ↓
sem parâmetros
  ↓
sem tipo de retorno declarado
  ↓
executa instruções
```

---

## 🧠 Funções também são valores

Um dos pontos importantes deste experimento é que uma função também pode ser armazenada em uma variável.

No código:

```go
var minhaFuncao = func() {
    fmt.Println("Essa função está armazenada em uma variável.")
}
```

Temos:

```text
var minhaFuncao = func() { ... }
    │              │
    │              └── valor que é uma função
    └───────────────── variável
```

A variável:

```text
minhaFuncao
```

passa a guardar um valor que pode ser chamado como função.

Por isso podemos escrever:

```go
minhaFuncao()
```

---

## 🕶️ Função anônima

A expressão:

```go
func() {
    fmt.Println("Essa função está armazenada em uma variável.")
}
```

cria uma função sem declarar um nome diretamente depois de `func`.

Compare com uma função nomeada:

```go
func somar(num1, num2 int) int {
    return num1 + num2
}
```

Nela existe um nome logo depois de `func`:

```text
somar
```

Já na função anônima temos:

```go
func() {
    // ...
}
```

Podemos visualizar:

```text
função nomeada
     ↓
func somar(...) {...}
     ↑
   possui nome

função anônima
     ↓
func() {...}
     ↑
sem nome nessa expressão
```

No experimento, essa função anônima é armazenada na variável `minhaFuncao`.

---

## 📦 `var minhaFuncao = func() {...}`

A declaração completa é:

```go
var minhaFuncao = func() {
    fmt.Println("Essa função está armazenada em uma variável.")
}
```

Podemos acompanhar da direita para a esquerda:

```text
func() { ... }
      ↓
cria um valor de função
      ↓
minhaFuncao
      ↓
variável que guarda essa função
```

Depois:

```go
minhaFuncao()
```

executa o valor de função armazenado nessa variável.

A saída será:

```text
Essa função está armazenada em uma variável.
```

---

## 🧬 O tipo `func()`

No comentário do experimento aparece uma observação importante:

```go
// Mesmo estando em uma variável, o tipo de minhaFuncao continua sendo func().
```

A função armazenada não recebe parâmetros e não declara retornos.

Por isso, podemos representar seu tipo como:

```go
func()
```

Podemos imaginar:

```text
minhaFuncao
    ↓
guarda uma função
    ↓
tipo: func()
```

O fato de uma função estar armazenada em uma variável não faz com que ela deixe de ser uma função.

---

## 🔬 Função nomeada vs função armazenada em variável

Neste experimento aparecem as duas formas.

Função nomeada:

```go
func somar(num1, num2 int) int {
    return num1 + num2
}
```

Função anônima armazenada em variável:

```go
var minhaFuncao = func() {
    fmt.Println("Essa função está armazenada em uma variável.")
}
```

Podemos comparar:

| Característica | `somar` | `minhaFuncao` |
| --- | --- | --- |
| Possui uma declaração `func nome(...)` | Sim | Não |
| O valor da função está em uma variável | Não | Sim |
| Recebe parâmetros neste exemplo | Sim | Não |
| Retorna valor neste exemplo | Sim | Não |
| Pode ser chamada com `()` | Sim | Sim |

Chamadas:

```go
somar(15, 10)
minhaFuncao()
```

Apesar de terem sido criadas de formas diferentes, as duas podem ser executadas através de uma chamada de função.

---

## 📤📤 Múltiplos retornos

Go permite que uma função devolva mais de um valor em uma única chamada.

No experimento:

```go
func calculadora(n1, n2 int) (int, int, int, int) {
    soma := n1 + n2
    subtracao := n1 - n2
    multiplicacao := n1 * n2
    divisao := n1 / n2

    return soma, subtracao, multiplicacao, divisao
}
```

Depois da lista de parâmetros encontramos:

```go
(int, int, int, int)
```

Isso indica que a função retorna quatro valores.

Podemos visualizar:

```text
func calculadora(n1, n2 int) (int, int, int, int)
                              │    │    │    │
                              │    │    │    └── quarto retorno
                              │    │    └─────── terceiro retorno
                              │    └──────────── segundo retorno
                              └───────────────── primeiro retorno
```

---

## 🧮 A função `calculadora`

A função recebe:

```go
n1, n2 int
```

Dentro dela são calculados quatro valores:

```go
soma := n1 + n2
subtracao := n1 - n2
multiplicacao := n1 * n2
divisao := n1 / n2
```

Depois todos são devolvidos:

```go
return soma, subtracao, multiplicacao, divisao
```

Podemos imaginar:

```text
calculadora(10, 5)
        ↓
 n1 = 10
 n2 = 5
        ↓
┌───────────────────┐
│ soma          = 15 │
│ subtracao      = 5 │
│ multiplicacao = 50 │
│ divisao        = 2 │
└───────────────────┘
        ↓
return 15, 5, 50, 2
```

---

## 📦 Recebendo vários retornos

Na função `main`, a chamada é:

```go
soma, subtracao, multiplicacao, divisao := calculadora(10, 5)
```

A função retorna quatro valores e a chamada recebe os quatro:

```text
calculadora(10, 5)
        ↓
15, 5, 50, 2
│   │   │   │
│   │   │   └── divisao
│   │   └────── multiplicacao
│   └────────── subtracao
└────────────── soma
```

Depois podemos utilizar cada variável separadamente:

```go
fmt.Println("Soma:", soma)
fmt.Println("Subtração:", subtracao)
fmt.Println("Multiplicação:", multiplicacao)
fmt.Println("Divisão:", divisao)
```

---

## 🔢 A ordem dos retornos importa

A função devolve:

```go
return soma, subtracao, multiplicacao, divisao
```

A chamada recebe:

```go
soma, subtracao, multiplicacao, divisao := calculadora(10, 5)
```

Os valores são associados pela posição.

Podemos representar:

```text
return soma, subtracao, multiplicacao, divisao
       │       │             │            │
       │       │             │            └── 4º
       │       │             └─────────────── 3º
       │       └───────────────────────────── 2º
       └───────────────────────────────────── 1º

       ↓       ↓             ↓            ↓

     soma  subtracao  multiplicacao   divisao
```

Por isso, a sequência utilizada no `return` e a sequência utilizada para receber os valores precisam representar a mesma ordem esperada pelo programa.

---

## ⚡ Múltiplos retornos com `:=`

Na aula de variáveis vimos que:

```go
:=
```

pode declarar variáveis dentro de funções.

Aqui ele aparece declarando várias variáveis ao mesmo tempo:

```go
soma, subtracao, multiplicacao, divisao := calculadora(10, 5)
```

Podemos imaginar:

```text
calculadora(10, 5)
        ↓
quatro valores retornados
        ↓
:=
        ↓
quatro variáveis recebem os resultados
```

Essa combinação é bastante útil quando uma função possui múltiplos retornos.

---

## 📝 Retornando mais de uma `string`

Múltiplos retornos não aparecem apenas na função `calculadora`.

O experimento também possui:

```go
func nomeCompleto(nome, sobrenome string) (string, string) {
    return nome, sobrenome
}
```

A função recebe duas strings:

```text
nome
sobrenome
```

E retorna duas strings:

```go
(string, string)
```

Podemos visualizar:

```text
nomeCompleto("Wesley", "Murat")
              ↓          ↓
            nome     sobrenome
              ↓          ↓
         return nome, sobrenome
              ↓          ↓
          "Wesley"    "Murat"
```

---

## 🕳️ Ignorando um retorno com `_`

Na chamada de `nomeCompleto`, o código utiliza:

```go
nome, _ := nomeCompleto("Wesley", "Murat")
```

A função retorna dois valores:

```text
"Wesley"
"Murat"
```

Porém, neste ponto do programa queremos utilizar apenas o primeiro.

Por isso temos:

```text
nome ← "Wesley"
_    ← "Murat"
```

O identificador:

```go
_
```

é utilizado para indicar que aquele valor será ignorado.

Depois o programa utiliza apenas:

```go
fmt.Println("Nome:", nome)
```

---

## 🧠 O identificador em branco

O `_` é conhecido como **identificador em branco**.

No experimento:

```go
nome, _ := nomeCompleto("Wesley", "Murat")
```

podemos interpretar:

```text
primeiro retorno
      ↓
    nome

segundo retorno
      ↓
      _
      ↓
   ignorado
```

Isso é útil quando uma função retorna vários valores, mas uma chamada específica não precisa utilizar todos eles.

---

## ⚠️ `_` não é uma variável comum

Observe:

```go
nome, _ := nomeCompleto("Wesley", "Murat")
```

O objetivo de `_` é descartar aquele valor.

Não estamos criando uma variável para consultar depois.

Uma forma simples de lembrar:

```text
nome
 ↓
valor que queremos utilizar

_
 ↓
valor que queremos ignorar
```

Neste exemplo, o programa precisa apenas do nome e não utiliza o sobrenome retornado.

---

## 🔬 As funções do experimento

O arquivo possui quatro comportamentos relacionados a funções.

### `somar`

```go
func somar(num1, num2 int) int {
    return num1 + num2
}
```

Demonstra:

```text
parâmetros
    +
retorno único
```

### `minhaFuncao`

```go
var minhaFuncao = func() {
    fmt.Println("Essa função está armazenada em uma variável.")
}
```

Demonstra:

```text
função anônima
      +
função armazenada em variável
```

### `calculadora`

```go
func calculadora(n1, n2 int) (int, int, int, int) {
    soma := n1 + n2
    subtracao := n1 - n2
    multiplicacao := n1 * n2
    divisao := n1 / n2

    return soma, subtracao, multiplicacao, divisao
}
```

Demonstra:

```text
parâmetros
    +
múltiplos retornos
```

### `nomeCompleto`

```go
func nomeCompleto(nome, sobrenome string) (string, string) {
    return nome, sobrenome
}
```

Demonstra:

```text
dois parâmetros string
        +
dois retornos string
        +
um retorno pode ser ignorado com _
```

---

## 🔄 Chamadas realizadas em `main`

Dentro da função principal, temos quatro chamadas importantes:

```go
resultado := somar(15, 10)
```

```go
minhaFuncao()
```

```go
soma, subtracao, multiplicacao, divisao := calculadora(10, 5)
```

```go
nome, _ := nomeCompleto("Wesley", "Murat")
```

Cada chamada demonstra uma possibilidade diferente:

| Chamada | Conceito principal |
| --- | --- |
| `somar(15, 10)` | parâmetros e retorno único |
| `minhaFuncao()` | função armazenada em variável |
| `calculadora(10, 5)` | múltiplos retornos |
| `nomeCompleto("Wesley", "Murat")` | descarte de retorno com `_` |

---

## 🧭 Fluxo geral do programa

Podemos visualizar o experimento inteiro desta forma:

```text
main()
 │
 ├── somar(15, 10)
 │      ↓
 │     25
 │      ↓
 │   resultado
 │
 ├── minhaFuncao()
 │      ↓
 │   imprime uma mensagem
 │
 ├── calculadora(10, 5)
 │      ↓
 │   15, 5, 50, 2
 │      ↓
 │   quatro variáveis
 │
 └── nomeCompleto("Wesley", "Murat")
        ↓
     "Wesley", "Murat"
        ↓        ↓
      nome       _
                 ↓
              ignorado
```

O arquivo utiliza funções para separar pequenas responsabilidades e demonstrar diferentes formas de entrada e saída de valores.

---

## 🧪 Resumo das formas estudadas

| Forma | Exemplo | Resultado |
| --- | --- | --- |
| Função com retorno único | `func somar(...) int` | retorna um `int` |
| Função sem parâmetros | `func()` | não recebe argumentos |
| Função sem retorno | `func() { ... }` | apenas executa instruções |
| Função em variável | `var minhaFuncao = func() {...}` | variável guarda uma função |
| Múltiplos retornos | `(int, int, int, int)` | devolve quatro valores |
| Descarte de retorno | `nome, _ := ...` | ignora um dos valores |

---

## 🧠 O que aprendi

Neste experimento foi possível observar que:

* funções agrupam instruções que podem ser executadas através de chamadas;
* uma função declarada utiliza a palavra-chave `func`;
* funções podem receber parâmetros;
* os valores enviados em uma chamada são os argumentos;
* parâmetros consecutivos do mesmo tipo podem compartilhar a declaração do tipo, como em `num1, num2 int`;
* uma função pode devolver um valor utilizando `return`;
* o tipo de retorno aparece depois da lista de parâmetros;
* o resultado de uma função pode ser armazenado em uma variável;
* uma função pode ser declarada depois de `main` e ainda ser chamada pelo programa;
* uma função pode não receber parâmetros;
* uma função pode não retornar valores;
* funções também podem ser tratadas como valores;
* uma função anônima pode ser armazenada em uma variável;
* no experimento, `minhaFuncao` armazena um valor do tipo `func()`;
* uma função pode retornar vários valores em uma única chamada;
* múltiplos retornos podem ser recebidos em várias variáveis ao mesmo tempo;
* a posição dos valores retornados determina qual variável recebe cada resultado;
* o identificador em branco `_` permite ignorar um valor que não será utilizado;
* `_` descarta o valor e não funciona como uma variável comum que consultamos depois.

---

> Funções ajudam a transformar uma sequência de instruções em pequenas responsabilidades que podem receber dados, executar uma tarefa e devolver resultados.

<p align="center">
  <img src="../docs/images/footer_04.jfif" alt="Go Soul Society">
</p>
