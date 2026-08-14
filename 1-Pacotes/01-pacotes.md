# 📦 01 — Pacotes em Go

Em Go, **pacotes** são uma das principais formas de organizar e reutilizar código.

Um pacote agrupa arquivos relacionados e permite separar responsabilidades dentro de uma aplicação.

Neste experimento, o objetivo é entender:

* o que são pacotes;
* como criar e importar pacotes;
* como funcionam pacotes dentro de outros diretórios;
* identificadores exportados e não exportados;
* o papel do `package main`;
* o papel do `main.go`;
* como o `go.mod` participa da organização do projeto;
* a diferença entre `go run`, `go build` e `go install`.

---

## 🧪 Estrutura do experimento

Neste exemplo, o projeto possui a seguinte estrutura:

```text
go-soul-society/
│
├── 1-Pacotes/
│   │
│   ├── auxiliar/
│   │   └── auxiliar.go
│   │
│   ├── 01-pacotes.md
│   └── main.go
│
├── go.mod
└── README.md
```

Existem dois pacotes envolvidos:

```text
1-Pacotes/
└── main.go          → package main

1-Pacotes/auxiliar/
└── auxiliar.go      → package auxiliar
```

Cada diretório representa um pacote diferente.

---

## 📦 Declarando um pacote

Todo arquivo `.go` começa declarando a qual pacote ele pertence.

Exemplo:

```go
package main
```

ou:

```go
package auxiliar
```

No nosso exemplo, o arquivo:

```text
1-Pacotes/main.go
```

pertence ao pacote:

```go
package main
```

Enquanto:

```text
1-Pacotes/auxiliar/auxiliar.go
```

pertence ao pacote:

```go
package auxiliar
```

Normalmente, os arquivos `.go` existentes dentro de um mesmo diretório pertencem ao mesmo pacote.

---

## 🧬 Pacotes e diretórios

Em Go, existe uma relação muito forte entre **diretórios e pacotes**.

Por exemplo:

```text
1-Pacotes/
├── main.go
└── auxiliar/
    └── auxiliar.go
```

Temos dois diretórios com código Go e, consequentemente, dois pacotes:

```text
1-Pacotes         → package main
auxiliar          → package auxiliar
```

O diretório `auxiliar` está fisicamente dentro de `1-Pacotes`, mas isso não significa que ele automaticamente faça parte do pacote `main`.

Ele é um pacote independente.

> Em Go, um diretório contendo código normalmente representa um pacote.

---

## 🔬 Criando o pacote `auxiliar`

O arquivo:

```text
1-Pacotes/auxiliar/auxiliar.go
```

contém:

```go
package auxiliar

import "fmt"

func Escrever() {
    fmt.Println("Esse código veio de outro pacote!!!")
}
```

A função `Escrever` pertence ao pacote `auxiliar`.

Para utilizá-la em outro pacote, precisamos importá-la.

---

## 📥 Importando outro pacote

No `main.go`:

```go
package main

import (
    "fmt"
    "go-soul-society/1-Pacotes/auxiliar"
)

func main() {
    fmt.Println("Escrevendo do arquivo main")
    auxiliar.Escrever()
}
```

O import:

```go
"go-soul-society/1-Pacotes/auxiliar"
```

pode ser dividido mentalmente em duas partes:

```text
go-soul-society
└── 1-Pacotes/auxiliar
```

Onde:

```text
go-soul-society
```

é o nome do módulo definido no `go.mod`.

E:

```text
1-Pacotes/auxiliar
```

é o caminho até o pacote que queremos importar.

---

## 🌎 Exportado e não exportado

Go não utiliza palavras-chave como:

```text
public
private
protected
```

como acontece em algumas outras linguagens.

A visibilidade de identificadores é determinada pela **primeira letra do nome**.

Na documentação de Go, normalmente usamos os termos:

* **exportado**;
* **não exportado**.

---

## 🔓 Letra maiúscula — exportado

Quando um identificador começa com uma letra maiúscula, ele pode ser acessado por outros pacotes.

Por exemplo:

```go
func Escrever() {
    fmt.Println("Esse código veio de outro pacote!!!")
}
```

Como `Escrever` começa com:

```text
E
```

maiúsculo, a função é exportada.

Por isso podemos fazer:

```go
auxiliar.Escrever()
```

a partir do pacote `main`.

Isso se aplica a vários tipos de identificadores, como:

```go
func Escrever() {}
type Usuario struct {}
var Nome string
const Pi = 3.14
```

Todos começam com letra maiúscula e podem ser exportados pelo pacote.

---

## 🔒 Letra minúscula: não exportado

Se alterarmos:

```go
func Escrever() {
}
```

para:

```go
func escrever() {
}
```

a função deixa de ser exportada.

Nesse caso, ela pode ser utilizada dentro do próprio pacote `auxiliar`, mas outro pacote não poderá fazer:

```go
auxiliar.escrever()
```

Portanto:

```text
Escrever()
↑
Maiúscula
↑
Exportado
```

Enquanto:

```text
escrever()
↑
Minúscula
↑
Não exportado
```

Uma forma simples de lembrar:

```text
Maiúscula → outros pacotes podem acessar
Minúscula → somente o próprio pacote pode acessar
```

> A visibilidade está relacionada ao pacote, e não ao arquivo.

Isso significa que vários arquivos pertencentes ao mesmo pacote conseguem compartilhar identificadores não exportados entre si.

---

## 🚀 `package main`

O pacote:

```go
package main
```

possui um significado especial.

Ele indica que aquele pacote representa um **programa executável**.

Para que esse programa possa iniciar sua execução, precisa existir uma função:

```go
func main() {
}
```

Por exemplo:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

A execução começa pela função:

```go
main()
```

Portanto:

```text
package main
     ↓
pacote executável

func main()
     ↓
ponto de entrada do programa
```

---

## 📄 O `main.go`

O nome:

```text
main.go
```

é uma convenção muito comum para o arquivo que contém:

```go
func main()
```

Porém, o nome do arquivo **não é o que transforma o programa em executável**.

O importante é existir:

```go
package main
```

e:

```go
func main() {
}
```

Por exemplo, tecnicamente poderíamos ter:

```text
programa.go
```

contendo:

```go
package main

func main() {
}
```

e o programa continuaria funcionando.

O nome `main.go` é utilizado principalmente por organização e convenção.

---

## 🧪 Vários `main.go` no mesmo projeto

Este repositório é organizado por experimentos.

Por isso, cada aula pode possuir seu próprio programa:

```text
go-soul-society/
│
├── 1-Pacotes/
│   └── main.go
│
├── 2-Variaveis/
│   └── main.go
│
├── 3-Tipos/
│   └── main.go
│
└── go.mod
```

Isso é válido porque cada `main.go` está em um diretório diferente.

Cada diretório representa um pacote independente.

Assim, podemos ter:

```text
1-Pacotes      → package main
2-Variaveis    → package main
3-Tipos        → package main
```

sem conflito entre eles.

Cada pasta representa um executável diferente.

---

## 🧩 O arquivo `go.mod`

Na raiz do projeto existe:

```text
go.mod
```

Ele define o **módulo Go**.

Neste projeto:

```go
module go-soul-society

go ...
```

A primeira linha:

```go
module go-soul-society
```

define o caminho base do módulo.

Podemos imaginar assim:

```text
go-soul-society
│
├── 1-Pacotes
│   └── auxiliar
│
├── 2-Variaveis
└── ...
```

Por isso conseguimos importar:

```go
import "go-soul-society/1-Pacotes/auxiliar"
```

O caminho começa pelo nome definido no:

```text
go.mod
```

e depois segue a estrutura de diretórios.

---

## 🧠 Módulo != pacote

É importante não confundir os dois conceitos.

Um **módulo** pode conter vários pacotes.

Neste projeto:

```text
go-soul-society              ← módulo
│
├── 1-Pacotes                ← pacote
│   └── auxiliar             ← outro pacote
│
├── 2-Variaveis              ← outro pacote
│
└── 3-Tipos                  ← outro pacote
```

Portanto:

```text
go.mod
   ↓
define o módulo

diretórios com código Go
   ↓
definem os pacotes
```

Uma forma simples de lembrar:

```text
Módulo
└── Pacote
    └── Arquivos .go
```

---

## ⚗️ `go run`

O comando:

```bash
go run .
```

compila e executa o pacote atual.

O:

```text
.
```

significa:

```text
diretório atual
```

Por exemplo:

```bash
cd 1-Pacotes
go run .
```

O Go encontra o `package main`, compila o programa e executa a função:

```go
func main()
```

O executável criado durante esse processo é temporário.

O comando é muito útil durante o desenvolvimento.

---

## ⚠️ `go run main.go`

Também é possível executar:

```bash
go run main.go
```

Nesse caso estamos informando explicitamente um arquivo.

Porém, conforme o programa cresce e o mesmo pacote passa a possuir vários arquivos, normalmente é mais interessante executar o pacote inteiro:

```bash
go run .
```

Por exemplo:

```text
meu-programa/
├── main.go
├── usuario.go
└── banco.go
```

Se todos pertencem ao:

```go
package main
```

podemos simplesmente executar:

```bash
go run .
```

Assim o Go considera todos os arquivos relevantes daquele pacote.

---

## 📍 Executando outro pacote sem entrar no diretório

Também podemos executar um pacote informando seu caminho.

Estando na raiz:

```text
go-soul-society/
```

podemos executar:

```bash
go run ./1-Pacotes
```

O:

```text
./
```

indica um caminho relativo ao diretório atual.

Portanto:

```bash
go run ./1-Pacotes
```

significa aproximadamente:

```text
compile e execute o pacote localizado em ./1-Pacotes
```

---

## 🔨 `go build`

Enquanto:

```bash
go run .
```

compila e imediatamente executa o programa, o comando:

```bash
go build .
```

serve para **compilar** o programa.

Quando executado em um `package main`, ele pode gerar um executável.

No Windows, por exemplo, podemos obter algo como:

```text
1-Pacotes.exe
```

Depois podemos executar esse arquivo sem precisar utilizar novamente:

```bash
go run .
```

Resumo:

```text
go run
   ↓
compila
   ↓
executa
   ↓
não mantém o executável como resultado principal
```

Enquanto:

```text
go build
   ↓
compila
   ↓
gera o executável
```

---

## 📦 `go install`

O comando:

```bash
go install .
```

também compila o programa, mas possui outro objetivo.

Em vez de simplesmente deixar o executável no diretório atual, ele instala o binário no diretório de binários configurado pelo ambiente Go.

Normalmente esse local está relacionado ao:

```text
GOBIN
```

ou, quando ele não está definido, ao diretório de binários associado ao ambiente Go.

Depois de instalado e estando esse diretório no `PATH`, o programa pode ser executado diretamente pelo nome.

Portanto:

```text
go build
   ↓
gera o executável para uso local
```

Enquanto:

```text
go install
   ↓
compila
   ↓
instala o executável no diretório de binários do Go
```

O `go install` também é muito utilizado para instalar ferramentas escritas em Go.

Um formato comum para ferramentas externas é:

```bash
go install caminho/do/pacote@versao
```

Por exemplo:

```text
go install exemplo/ferramenta@latest
```

O `@latest` indica que queremos instalar a versão mais recente disponível daquele programa.

---

## 🧪 Comparando os comandos

| Comando              | Função                                        |
| -------------------- | --------------------------------------------- |
| `go run .`           | Compila temporariamente e executa o pacote    |
| `go run main.go`     | Compila e executa os arquivos informados      |
| `go run ./1-Pacotes` | Executa um pacote através do caminho          |
| `go build .`         | Compila o pacote                              |
| `go install .`       | Compila e instala o executável no ambiente Go |

Durante o desenvolvimento, provavelmente o comando utilizado com maior frequência será:

```bash
go run .
```

---

## 🔬 Experimento

Arquivo:

```text
1-Pacotes/auxiliar/auxiliar.go
```

```go
package auxiliar

import "fmt"

func Escrever() {
    fmt.Println("Esse código veio de outro pacote!!!")
}
```

Arquivo:

```text
1-Pacotes/main.go
```

```go
package main

import (
    "fmt"

    "go-soul-society/1-Pacotes/auxiliar"
)

func main() {
    fmt.Println("Escrevendo do arquivo main")
    auxiliar.Escrever()
}
```

Executando:

```bash
go run ./1-Pacotes
```

Resultado:

```text
Escrevendo do arquivo main
Esse código veio de outro pacote!!!
```

O programa principal conseguiu chamar uma função definida em outro pacote.

---

## 🧠 O que aprendi

Neste experimento foi possível observar que:

* arquivos Go pertencem a pacotes;
* normalmente cada diretório representa um pacote;
* um diretório dentro de outro pode representar outro pacote independente;
* pacotes podem importar e utilizar outros pacotes;
* identificadores iniciados com letra maiúscula são exportados;
* identificadores iniciados com letra minúscula não são exportados;
* `package main` representa um programa executável;
* `func main()` representa o ponto inicial desse programa;
* `main.go` é uma convenção de organização, não uma exigência de nome;
* um projeto pode possuir vários `package main` em diretórios diferentes;
* `go.mod` define o módulo e serve como base para os caminhos de importação;
* um módulo pode possuir vários pacotes;
* `go run` compila e executa;
* `go build` compila o programa;
* `go install` compila e instala um executável no ambiente Go.

---
<p align="center">
  <img src="../docs/images/footer.jfif" alt="Go Soul Society">
</p>
