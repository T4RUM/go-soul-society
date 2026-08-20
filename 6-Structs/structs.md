# 🧱 06 — Structs em Go

Em Go, **structs** permitem agrupar vários valores relacionados dentro de um único tipo.

Elas são muito úteis quando diferentes informações pertencem ao mesmo conceito. Um usuário, por exemplo, pode possuir nome, idade e endereço. Em vez de manter cada informação completamente separada, podemos criar uma struct que represente esse conjunto de dados.

Neste experimento, o objetivo é entender:

* o que é uma `struct`;
* como declarar um novo tipo utilizando `type`;
* como definir campos dentro de uma struct;
* como criar uma variável de um tipo struct;
* quais valores os campos recebem quando a struct é criada sem valores explícitos;
* como acessar e alterar campos utilizando `.`;
* como criar várias instâncias do mesmo tipo;
* como imprimir structs no terminal;
* o que é um literal de struct;
* como criar uma struct informando valores pela posição dos campos;
* como criar uma struct utilizando nomes de campos;
* por que a sintaxe `user("Wesley", 28)` não cria uma struct em Go;
* como omitir campos em um literal com nomes;
* como criar uma struct para representar um endereço;
* como aninhar uma struct dentro de outra;
* como preencher uma struct aninhada campo a campo;
* como criar structs aninhadas utilizando literais;
* como acessar e alterar campos de uma struct interna;
* a diferença entre uma struct aninhada com campo nomeado e uma struct embutida;
* como a visibilidade dos campos se relaciona com letras maiúsculas e minúsculas.

---

## 🧪 Estrutura do experimento

Nesta aula, podemos organizar os arquivos da seguinte forma:

```text
go-soul-society/
│
├── 6-Structs/
│   │
│   ├── 06-structs.md
│   └── structs.go
│
├── go.mod
└── README.md
```

O arquivo:

```text
6-Structs/structs.go
```

contém os exemplos utilizados para estudar structs.

---

## 🧱 O que é uma struct?

Uma struct é um tipo que agrupa vários campos relacionados.

Podemos imaginar um usuário com duas informações:

```text
nome
idade
```

Sem uma struct, poderíamos armazenar esses valores separadamente:

```go
nome := "Wesley"
idade := 28
```

Isso funciona, mas os dois valores não estão agrupados em um único tipo que represente a ideia de um usuário.

Com uma struct, podemos criar algo conceitualmente parecido com:

```text
user
 │
 ├── name
 └── age
```

Agora `name` e `age` passam a fazer parte do mesmo tipo.

---

## 🧬 Criando um novo tipo com `type`

Para declarar um tipo próprio em Go, utilizamos a palavra-chave:

```go
type
```

No experimento:

```go
type user struct {
    name string
    age  int
}
```

Podemos separar essa declaração em partes:

```text
type user struct {
│    │    │
│    │    └── tipo estruturado
│    └─────── nome do novo tipo
└──────────── declaração de tipo
```

A partir dessa declaração, `user` passa a ser um tipo disponível naquele escopo.

---

## 🏷️ O nome do tipo

Neste exemplo:

```go
type user struct {
    name string
    age  int
}
```

criamos um tipo chamado:

```text
user
```

Depois podemos declarar variáveis desse tipo:

```go
var myProfile user
```

A leitura pode ser feita assim:

```text
myProfile
    ↓
variável do tipo user
```

---

## 🧩 Campos de uma struct

Dentro da struct declaramos os campos que pertencem àquele tipo.

No exemplo:

```go
type user struct {
    name string
    age  int
}
```

existem dois campos:

```text
name → string
age  → int
```

Podemos visualizar:

```text
user
 │
 ├── name → string
 │
 └── age  → int
```

Cada campo possui:

* um nome;
* um tipo.

---

## 📦 Declarando uma variável do tipo struct

Depois que o tipo existe, podemos criar uma variável normalmente:

```go
var myProfile user
```

Essa declaração cria uma variável chamada:

```text
myProfile
```

cujo tipo é:

```text
user
```

Podemos imaginar:

```text
myProfile
   ↓
 user
   │
   ├── name
   └── age
```

---

## 🕳️ Valores zero dos campos

Quando declaramos:

```go
var myProfile user
```

sem informar valores, a struct ainda é criada.

Cada campo recebe o **valor zero** do seu próprio tipo.

Considerando:

```go
type user struct {
    name string
    age  int
}
```

temos:

```text
name → string → ""
age  → int    → 0
```

Portanto, logo depois da declaração:

```text
myProfile
 │
 ├── name = ""
 └── age  = 0
```

Esse comportamento segue a mesma ideia dos valores zero estudados anteriormente para variáveis.

---

## 🎯 Acessando um campo com `.`

Para acessar um campo de uma struct utilizamos:

```text
.
```

No experimento:

```go
myProfile.name = "Wesley"
myProfile.age = 28
```

Podemos ler:

```text
myProfile.name
    │       │
    │       └── campo name
    └────────── variável myProfile
```

E:

```text
myProfile.age
    │      │
    │      └── campo age
    └───────── variável myProfile
```

Depois dessas atribuições:

```text
myProfile
 │
 ├── name = "Wesley"
 └── age  = 28
```

---

## 🔄 Alterando os valores dos campos

Os campos podem receber novos valores da mesma forma que outras variáveis.

Exemplo:

```go
myProfile.age = 28
```

Depois poderíamos fazer:

```go
myProfile.age = 29
```

O campo `age` continua sendo um `int`.

Apenas seu valor foi alterado.

```text
antes
age = 28
   ↓
atribuição
   ↓
depois
age = 29
```

---

## 👥 Várias variáveis do mesmo tipo

Depois de declarar o tipo:

```go
type user struct {
    name string
    age  int
}
```

podemos criar várias variáveis desse tipo.

No código original da aula apareciam duas:

```go
var myProfile user
var otherProfile user
```

Cada uma possui seus próprios campos.

Por exemplo:

```go
myProfile.name = "Wesley"
myProfile.age = 28

otherProfile.name = "Elias"
otherProfile.age = 23
```

Podemos visualizar:

```text
myProfile
 │
 ├── name = "Wesley"
 └── age  = 28

otherProfile
 │
 ├── name = "Elias"
 └── age  = 23
```

As duas variáveis possuem o mesmo tipo, mas armazenam valores diferentes.

---

## ✏️ Ajuste no nome `outherProfile`

No código original apareceu:

```go
var outherProfile user
```

O código poderia funcionar com esse nome porque identificadores podem possuir nomes escolhidos pelo programador.

Porém, a palavra em inglês esperada seria:

```text
other
```

Por isso, no código organizado utilizaremos:

```go
var otherProfile user
```

Essa alteração não muda o comportamento do programa.

Ela apenas melhora a legibilidade do nome.

---

## 🖨️ Imprimindo uma struct

Podemos enviar uma struct diretamente para:

```go
fmt.Println()
```

Por exemplo:

```go
fmt.Println(myProfile)
```

Uma saída simples poderia aparecer como:

```text
{Wesley 28}
```

Isso mostra os valores dos campos, mas não mostra seus nomes.

---

## 🔍 Imprimindo os nomes dos campos com `%+v`

Durante os estudos, pode ser útil utilizar:

```go
fmt.Printf("%+v\n", myProfile)
```

O formato:

```text
%+v
```

faz com que os nomes dos campos também apareçam.

Por exemplo:

```text
{name:Wesley age:28}
```

No código final utilizaremos bastante essa forma porque ela facilita visualizar a estrutura dos dados.

---

## 🧪 O que é um literal de struct?

Além de declarar uma variável vazia e preencher cada campo depois, podemos criar uma struct já informando valores.

Esse formato é chamado de **literal de struct**.

Um exemplo simples é:

```go
profile := user{
    name: "Wesley",
    age:  28,
}
```

Podemos imaginar:

```text
user{...}
   ↓
cria um valor do tipo user
   ↓
profile
```

---

## ⚠️ `user("Wesley", 28)` não é a sintaxe de uma struct

No código original havia a tentativa comentada:

```go
// myProfile := user("Wesley",28)
```

Essa não é a sintaxe utilizada para criar um valor de struct em Go.

Parênteses nesse formato lembram uma chamada de função ou uma conversão de tipo:

```text
user(...)
```

Para criar um valor de struct utilizamos chaves:

```go
user{...}
```

Portanto, para uma struct com os campos:

```go
type user struct {
    name string
    age  int
}
```

uma forma válida seria:

```go
myProfile := user{"Wesley", 28}
```

ou, preferencialmente em muitos casos:

```go
myProfile := user{
    name: "Wesley",
    age:  28,
}
```

Resumo:

```text
user(...) → não cria uma struct dessa forma
user{...} → literal de struct
```

---

## 📍 Literal utilizando a posição dos campos

Uma struct pode ser criada informando apenas os valores:

```go
myProfile := user{"Wesley", 28}
```

Nesse formato, cada valor corresponde à posição de um campo na declaração.

Se temos:

```go
type user struct {
    name string
    age  int
}
```

então:

```go
user{"Wesley", 28}
```

pode ser visualizado assim:

```text
user{
    "Wesley",  → name
    28,        → age
}
```

---

## ⚠️ A ordem importa no literal por posição

Considere novamente:

```go
type user struct {
    name string
    age  int
}
```

A ordem declarada é:

```text
1 → name
2 → age
```

Por isso:

```go
user{"Wesley", 28}
```

segue essa ordem.

Esse formato fica mais frágil quando a struct cresce ou muda, porque os valores dependem diretamente da posição dos campos.

---

## ⚠️ Todos os campos precisam aparecer no literal por posição

Quando utilizamos um literal sem nomes de campos, precisamos fornecer um valor para cada campo da struct.

No código final, `user` também possuirá um endereço:

```go
type user struct {
    name    string
    age     int
    address address
}
```

Por isso, o exemplo posicional será:

```go
profileByPosition := user{
    "Carlos",
    30,
    address{},
}
```

Cada posição corresponde a:

```text
"Carlos"  → name
30        → age
address{} → address
```

---

## 🏷️ Literal utilizando nomes de campos

Outra forma de criar a struct é informar explicitamente o nome de cada campo.

Exemplo:

```go
myProfile := user{
    name: "Wesley",
    age:  28,
}
```

Agora não precisamos descobrir o significado de cada valor apenas pela posição.

A própria escrita deixa claro:

```text
name → "Wesley"
age  → 28
```

---

## ⚠️ A sintaxe `user(name:"Wesley")` também não é válida

No código original apareceu:

```go
// myProfile := user(name:"Wesley")
```

A intenção estava correta: criar uma struct informando um campo pelo nome.

Porém, a sintaxe precisa utilizar chaves:

```go
myProfile := user{
    name: "Wesley",
}
```

Portanto:

```text
user(name: "Wesley")
       ↓
não é a sintaxe de literal de struct
```

Enquanto:

```text
user{name: "Wesley"}
       ↓
literal de struct com campo nomeado
```

---

## 🕳️ Podemos omitir campos quando usamos nomes

Uma vantagem importante do literal com nomes é poder informar apenas os campos que desejamos inicializar.

Exemplo:

```go
profileByFields := user{
    name: "Mariana",
}
```

Se `user` possui:

```go
type user struct {
    name    string
    age     int
    address address
}
```

então os campos não informados recebem seus valores zero.

Podemos imaginar:

```text
name    = "Mariana"
age     = 0
address = valor zero de address
```

---

## 🔀 Com nomes, a ordem deixa de ser o significado principal

Quando utilizamos campos nomeados, podemos escrever:

```go
user{
    age:  28,
    name: "Wesley",
}
```

mesmo que a declaração seja:

```go
type user struct {
    name string
    age  int
}
```

Isso acontece porque cada valor está associado explicitamente ao nome do campo.

Por legibilidade, ainda podemos manter uma ordem organizada, mas o significado não depende mais da posição.

---

## 🧭 Comparando as duas formas

Considere uma struct simples:

```go
type user struct {
    name string
    age  int
}
```

Temos:

| Forma | Exemplo | Característica |
| --- | --- | --- |
| Por posição | `user{"Wesley", 28}` | depende da ordem e precisa preencher todos os campos |
| Com nomes | `user{name: "Wesley", age: 28}` | deixa explícito qual campo recebe cada valor |
| Com nomes e campos omitidos | `user{name: "Wesley"}` | campos restantes recebem o valor zero |

Para exemplos maiores, a forma com nomes costuma comunicar melhor a intenção do código.

---

## 🏠 Criando uma struct para endereço

Na aula também foi criada uma segunda struct para representar um endereço.

O código original possuía:

```go
type address struct {
    name string
    city string
}
```

No código organizado, utilizaremos:

```go
type address struct {
    street string
    city   string
}
```

O campo:

```text
street
```

comunica melhor que o valor representa o nome da rua ou logradouro.

---

## ✏️ Ajuste de `myAdress` para `myAddress`

No código original apareceu:

```go
var myAdress address
```

A grafia em inglês de endereço é:

```text
address
```

com duas letras `d`.

Por isso, utilizaremos:

```go
var myAddress address
```

Novamente, é uma melhoria de nome e legibilidade.

---

## 🧱 Criando um endereço campo a campo

Podemos declarar:

```go
var myAddress address
```

Inicialmente:

```text
myAddress
 │
 ├── street = ""
 └── city   = ""
```

Depois:

```go
myAddress.street = "Massachusetts Ave NW"
myAddress.city = "Washington"
```

A struct passa a representar:

```text
myAddress
 │
 ├── street = "Massachusetts Ave NW"
 └── city   = "Washington"
```

---

# 🪆 Como aninhar structs

Uma struct pode possuir um campo cujo tipo também é uma struct.

Esse é o conceito principal de **aninhamento de structs** utilizado nesta aula.

Temos primeiro:

```go
type address struct {
    street string
    city   string
}
```

Depois podemos utilizar `address` como o tipo de um campo dentro de `user`:

```go
type user struct {
    name    string
    age     int
    address address
}
```

Agora a struct `user` contém outra struct.

---

## 🧬 Entendendo `address address`

Esta linha pode parecer estranha no começo:

```go
address address
```

Mas ela segue a mesma regra de qualquer outro campo.

Por exemplo:

```go
name string
```

significa:

```text
campo: name
tipo:  string
```

Da mesma forma:

```go
address address
```

significa:

```text
campo: address
tipo:  address
```

Podemos visualizar:

```text
user
 │
 ├── name    → string
 ├── age     → int
 └── address → address
                  │
                  ├── street → string
                  └── city   → string
```

---

## 🧩 A struct externa e a struct interna

Quando escrevemos:

```go
type user struct {
    name    string
    age     int
    address address
}
```

podemos pensar em duas camadas:

```text
user                 ← struct externa
 │
 ├── name
 ├── age
 └── address          ← campo
      │
      ├── street      ← campo da struct interna
      └── city        ← campo da struct interna
```

O valor armazenado em `user.address` é um valor do tipo `address`.

---

## 🕳️ Valor zero de uma struct aninhada

Se declararmos:

```go
var profile user
```

sem fornecer nenhum valor, `profile.address` também existe.

Ele apenas contém os valores zero dos seus campos.

Considerando:

```go
type address struct {
    street string
    city   string
}
```

teremos:

```text
profile
 │
 ├── name = ""
 ├── age  = 0
 └── address
      │
      ├── street = ""
      └── city   = ""
```

Não precisamos criar `address` separadamente para que esses campos existam dentro da variável.

---

## 🎯 Acessando campos da struct interna

Para chegar ao campo `city`, primeiro acessamos o campo `address` de `user`.

Depois acessamos `city` dentro de `address`.

```go
profile.address.city
```

Podemos decompor:

```text
profile.address.city
   │       │      │
   │       │      └── campo city da struct address
   │       └───────── campo address da struct user
   └───────────────── variável profile
```

---

## 🧱 Preenchendo uma struct aninhada campo a campo

Podemos fazer:

```go
var profileWithAddress user

profileWithAddress.name = "Wesley"
profileWithAddress.age = 28
profileWithAddress.address.street = "Massachusetts Ave NW"
profileWithAddress.address.city = "Washington"
```

O caminho:

```go
profileWithAddress.address.street
```

significa:

```text
profileWithAddress
       ↓
address
       ↓
street
```

Depois das atribuições:

```text
profileWithAddress
 │
 ├── name = "Wesley"
 ├── age  = 28
 └── address
      │
      ├── street = "Massachusetts Ave NW"
      └── city   = "Washington"
```

---

## 🖨️ Lendo apenas um campo da struct interna

Depois de preencher:

```go
profileWithAddress.address.city = "Washington"
```

podemos acessar apenas a cidade:

```go
fmt.Println("Cidade:", profileWithAddress.address.city)
```

Resultado:

```text
Cidade: Washington
```

---

## 🪆 Criando structs aninhadas com literal

Também podemos criar a struct externa e a interna ao mesmo tempo.

Exemplo:

```go
completeProfile := user{
    name: "Elias",
    age:  23,
    address: address{
        street: "Avenida Brasil",
        city:   "Washington",
    },
}
```

Observe que existem dois literais:

```text
user{ ... }
     │
     └── contém address{ ... }
```

Podemos visualizar:

```text
user{
    name: "Elias",
    age: 23,
    address: address{
        street: "Avenida Brasil",
        city: "Washington",
    },
}
```

---

## 🔬 Entendendo o literal aninhado passo a passo

Primeiro temos o valor externo:

```go
user{
}
```

Depois preenchemos campos simples:

```go
user{
    name: "Elias",
    age:  23,
}
```

O campo `address` espera um valor do tipo:

```text
address
```

Então podemos fornecer:

```go
address{
    street: "Avenida Brasil",
    city:   "Washington",
}
```

Juntando tudo:

```go
user{
    name: "Elias",
    age:  23,
    address: address{
        street: "Avenida Brasil",
        city:   "Washington",
    },
}
```

---

## 🧭 Fluxo do valor aninhado

Podemos imaginar o processo assim:

```text
address{
    street: "Avenida Brasil",
    city: "Washington",
}
        ↓
cria um valor do tipo address
        ↓
esse valor é colocado no campo user.address
        ↓
completeProfile
```

---

## 🏠 Reutilizando uma variável `address`

Não somos obrigados a criar a struct interna diretamente dentro do literal.

Também podemos criar um endereço antes:

```go
myAddress := address{
    street: "Massachusetts Ave NW",
    city:   "Washington",
}
```

E depois utilizá-lo:

```go
profile := user{
    name:    "Wesley",
    age:     28,
    address: myAddress,
}
```

Nesse caso:

```text
myAddress
    ↓
valor do tipo address
    ↓
profile.address
```

As duas formas são válidas.

---

## 🔄 Alterando um campo da struct interna

Depois de criar:

```go
completeProfile := user{
    name: "Elias",
    age:  23,
    address: address{
        street: "Avenida Brasil",
        city:   "Washington",
    },
}
```

podemos alterar somente a cidade:

```go
completeProfile.address.city = "Votorantim"
```

Antes:

```text
completeProfile.address.city
              ↓
          "Washington"
```

Depois:

```text
completeProfile.address.city
              ↓
          "Votorantim"
```

Não foi necessário substituir toda a struct `address`.

---

## 🧠 O operador `.` pode formar um caminho

Em uma struct simples:

```go
profile.name
```

existe um acesso.

Em uma struct aninhada:

```go
profile.address.city
```

existem acessos encadeados.

Podemos ler da esquerda para a direita:

```text
profile
   ↓
address
   ↓
city
```

Esse padrão aparecerá com frequência quando trabalharmos com estruturas de dados mais complexas.

---

## 🪆 Podemos ter mais de um nível de aninhamento

O conceito não é limitado a duas structs.

Por exemplo, futuramente poderíamos ter algo como:

```go
type country struct {
    name string
}

type address struct {
    street  string
    city    string
    country country
}

type user struct {
    name    string
    address address
}
```

O acesso poderia chegar a:

```go
profile.address.country.name
```

Visualmente:

```text
user
 └── address
      └── country
           └── name
```

Não precisamos utilizar tantos níveis sem necessidade, mas o mecanismo é o mesmo.

---

## 🧩 Aninhamento ajuda a representar relações

Podemos pensar em structs como pequenos blocos de dados.

```text
address
 ├── street
 └── city
```

Depois outro tipo pode utilizar esse bloco:

```text
user
 ├── name
 ├── age
 └── address
      ├── street
      └── city
```

Em vez de colocar todos os campos diretamente em `user`:

```go
type user struct {
    name   string
    age    int
    street string
    city   string
}
```

podemos separar a responsabilidade do endereço:

```go
type address struct {
    street string
    city   string
}

type user struct {
    name    string
    age     int
    address address
}
```

Essa organização deixa explícito que `street` e `city` pertencem ao conceito de endereço.

---

## 🧱 Struct aninhada não é a mesma coisa que embedding

Existe outro recurso de Go chamado **embedding** de structs.

Nesta aula, nosso foco principal é o aninhamento com um campo nomeado:

```go
type user struct {
    name    string
    address address
}
```

Aqui existe um campo chamado:

```text
address
```

Por isso acessamos:

```go
profile.address.city
```

---

## 🧬 Exemplo de struct embutida

Go também permite escrever:

```go
type user struct {
    name string
    address
}
```

Observe que agora aparece apenas:

```go
address
```

sem um nome de campo separado.

Esse recurso é chamado de **embedded field** ou campo embutido.

Com embedding, alguns campos podem ser promovidos e acessados de forma mais curta, por exemplo:

```go
profile.city
```

em determinados contextos.

Esse comportamento é diferente do aninhamento nomeado utilizado no experimento principal.

Por enquanto, uma forma simples de lembrar é:

```text
address address
      ↓
campo nomeado
      ↓
profile.address.city
```

Enquanto:

```text
address
   ↓
campo embutido
   ↓
conceito adicional da linguagem
```

O código principal desta aula utiliza o campo nomeado porque torna a estrutura explícita e é mais simples para estudar o aninhamento.

---

## 📍 Onde declarar os tipos?

No código original da aula, os tipos foram declarados dentro de `main`:

```go
func main() {
    type user struct {
        name string
        age  int
    }
}
```

Isso é permitido em Go.

Nesse caso, o tipo é local à função e só pode ser utilizado dentro daquele escopo.

---

## 📦 Tipos declarados no nível do pacote

No código organizado, utilizaremos:

```go
type address struct {
    street string
    city   string
}

type user struct {
    name    string
    age     int
    address address
}

func main() {
    // ...
}
```

Agora os tipos são declarados no nível do pacote.

Podemos visualizar:

```text
package main
 │
 ├── type address
 ├── type user
 └── func main
```

Essa organização facilita a leitura e permite que outras funções do mesmo pacote utilizem esses tipos.

> A declaração dentro de `main` não estava errada. A mudança é uma melhoria de organização para o experimento completo.

---

## 🌎 Campos com letra minúscula

No experimento utilizamos:

```go
type user struct {
    name string
    age  int
}
```

Os campos começam com letras minúsculas:

```text
name
age
```

Como estudamos na aula sobre pacotes, identificadores iniciados com letra minúscula não são exportados para outros pacotes.

Isso não causa problema aqui porque o código utiliza esses campos dentro do próprio:

```text
package main
```

---

## 🔓 Como seriam campos exportados?

Se futuramente quisermos acessar esses campos a partir de outro pacote, poderíamos declarar:

```go
type User struct {
    Name    string
    Age     int
    Address Address
}
```

E:

```go
type Address struct {
    Street string
    City   string
}
```

As letras maiúsculas indicariam identificadores exportados.

Nesta aula manteremos os nomes minúsculos para preservar a simplicidade do experimento dentro do mesmo pacote.

---

## 🧩 Struct é um tipo; variável é um valor desse tipo

É importante separar duas ideias.

Quando escrevemos:

```go
type user struct {
    name string
    age  int
}
```

estamos declarando um **tipo**.

Quando escrevemos:

```go
var myProfile user
```

estamos criando uma **variável** daquele tipo.

Podemos resumir:

```text
type user struct { ... }
        ↓
define o formato

var myProfile user
        ↓
cria uma variável seguindo esse formato
```

---

## 🏗️ Uma struct funciona como um formato de dados

Podemos imaginar a declaração:

```go
type user struct {
    name    string
    age     int
    address address
}
```

como um formato:

```text
user
 ├── name: string
 ├── age: int
 └── address: address
```

Depois podemos criar vários valores seguindo esse mesmo formato:

```text
myProfile
otherProfile
profileByPosition
profileByFields
profileWithAddress
completeProfile
```

Todos podem possuir o tipo `user`.

---

## ⚠️ Uma struct não é um `map`

Uma struct possui campos definidos no momento em que o tipo é declarado.

Por exemplo:

```go
type user struct {
    name string
    age  int
}
```

Não podemos simplesmente inventar depois um campo que não existe:

```go
myProfile.email = "wesley@example.com"
```

se `email` não fizer parte da declaração de `user`.

Para adicionar esse campo ao formato, precisaríamos alterar o tipo:

```go
type user struct {
    name  string
    age   int
    email string
}
```

Maps serão estudados separadamente e possuem outro comportamento.

---

## 🔬 Código original e código organizado

O código original foi importante para introduzir:

* declaração de `struct`;
* criação de variáveis do tipo struct;
* acesso aos campos;
* criação de uma segunda struct para endereço.

No código organizado, foram realizadas algumas melhorias:

```text
outherProfile → otherProfile
myAdress      → myAddress
address.name  → address.street
```

Também:

```text
tipos dentro de main
        ↓
tipos no nível do pacote
```

E as tentativas:

```go
user("Wesley", 28)
user(name: "Wesley")
```

foram substituídas pelas formas corretas:

```go
user{"Wesley", 28, address{}}
```

ou:

```go
user{
    name: "Wesley",
}
```

Finalmente, o campo:

```go
address address
```

foi adicionado a `user` para completar o conteúdo sobre structs aninhadas.

---

## 🧪 Código completo do experimento

Arquivo:

```text
6-Structs/structs.go
```

```go
package main

import "fmt"

// address representa um endereço que poderá ser utilizado dentro de outra struct.
type address struct {
    street string
    city   string
}

// user agrupa os dados que representam um usuário.
// O campo address demonstra como uma struct pode conter outra struct.
type user struct {
    name    string
    age     int
    address address
}

func main() {
    // =========================================================================
    // DECLARANDO E PREENCHENDO UMA STRUCT CAMPO A CAMPO
    // =========================================================================

    // Ao declarar uma variável do tipo user sem informar valores,
    // todos os campos começam com o valor zero de seus respectivos tipos.
    var myProfile user

    myProfile.name = "Wesley"
    myProfile.age = 28

    fmt.Println("=== Struct preenchida campo a campo ===")
    fmt.Printf("%+v\n", myProfile)

    // Cada variável do tipo user possui sua própria cópia dos campos.
    var otherProfile user
    otherProfile.name = "Elias"
    otherProfile.age = 23

    fmt.Println("\n=== Outra variável do mesmo tipo ===")
    fmt.Printf("%+v\n", otherProfile)

    // =========================================================================
    // LITERAIS DE STRUCT
    // =========================================================================

    // Uma struct pode ser criada informando os valores pela posição dos campos.
    // Nesse formato, a ordem precisa ser exatamente a mesma da declaração da struct.
    profileByPosition := user{
        "Carlos",
        30,
        address{},
    }

    fmt.Println("\n=== Literal por posição ===")
    fmt.Printf("%+v\n", profileByPosition)

    // A forma com nomes de campos costuma ser mais legível.
    // Campos não informados recebem o valor zero de seu tipo.
    profileByFields := user{
        name: "Mariana",
    }

    fmt.Println("\n=== Literal com nomes de campos ===")
    fmt.Printf("%+v\n", profileByFields)

    // =========================================================================
    // STRUCT ADDRESS
    // =========================================================================

    var myAddress address
    myAddress.street = "Massachusetts Ave NW"
    myAddress.city = "Washington"

    fmt.Println("\n=== Struct de endereço ===")
    fmt.Printf("%+v\n", myAddress)

    // =========================================================================
    // COMO ANINHAR STRUCTS
    // =========================================================================

    // Como user possui um campo chamado address do tipo address,
    // podemos preencher os campos da struct interna usando acesso encadeado.
    var profileWithAddress user
    profileWithAddress.name = "Wesley"
    profileWithAddress.age = 28
    profileWithAddress.address.street = "Massachusetts Ave NW"
    profileWithAddress.address.city = "Washington"

    fmt.Println("\n=== Struct aninhada preenchida campo a campo ===")
    fmt.Printf("%+v\n", profileWithAddress)
    fmt.Println("Cidade:", profileWithAddress.address.city)

    // Também podemos construir a struct externa e a struct interna
    // ao mesmo tempo utilizando literais com nomes de campos.
    completeProfile := user{
        name: "Elias",
        age:  23,
        address: address{
            street: "Avenida Brasil",
            city:   "Washington",
        },
    }

    fmt.Println("\n=== Struct aninhada criada com literal ===")
    fmt.Printf("%+v\n", completeProfile)
    fmt.Println("Rua:", completeProfile.address.street)
    fmt.Println("Cidade:", completeProfile.address.city)

    // Para alterar um campo da struct interna, usamos o mesmo caminho encadeado.
    completeProfile.address.city = "Votorantim"

    fmt.Println("\n=== Alterando um campo da struct interna ===")
    fmt.Println("Nova cidade:", completeProfile.address.city)
}
```

---

## ▶️ Executando o experimento

Dentro do diretório da aula, podemos executar:

```bash
go run .
```

Ou, informando o arquivo diretamente:

```bash
go run structs.go
```

---

## 🖨️ Resultado do código organizado

Uma execução do programa produz:

```text
=== Struct preenchida campo a campo ===
{name:Wesley age:28 address:{street: city:}}

=== Outra variável do mesmo tipo ===
{name:Elias age:23 address:{street: city:}}

=== Literal por posição ===
{name:Carlos age:30 address:{street: city:}}

=== Literal com nomes de campos ===
{name:Mariana age:0 address:{street: city:}}

=== Struct de endereço ===
{street:Massachusetts Ave NW city:Washington}

=== Struct aninhada preenchida campo a campo ===
{name:Wesley age:28 address:{street:Massachusetts Ave NW city:Washington}}
Cidade: Washington

=== Struct aninhada criada com literal ===
{name:Elias age:23 address:{street:Avenida Brasil city:Washington}}
Rua: Avenida Brasil
Cidade: Washington

=== Alterando um campo da struct interna ===
Nova cidade: Votorantim
```

---

## 🔍 Observando os valores zero na saída

Na primeira saída temos:

```text
{name:Wesley age:28 address:{street: city:}}
```

O endereço ainda não foi preenchido.

Por isso:

```text
street = ""
city   = ""
```

O `fmt.Printf` mostra os nomes dos campos, mas strings vazias aparecem sem conteúdo depois dos dois pontos:

```text
street:
city:
```

Isso não significa que os campos não existem.

Eles existem e possuem o valor zero de `string`.

---

## 🔍 Observando campos omitidos no literal

No exemplo:

```go
profileByFields := user{
    name: "Mariana",
}
```

não informamos:

```text
age
address
```

Por isso a saída contém:

```text
{name:Mariana age:0 address:{street: city:}}
```

Podemos decompor:

```text
name = "Mariana"
age = 0
address.street = ""
address.city = ""
```

---

## 🧭 Fluxo geral do experimento

Podemos visualizar a aula inteira desta forma:

```text
type address struct
        │
        ├── street
        └── city

        ↓ utilizado por

type user struct
        │
        ├── name
        ├── age
        └── address
             │
             ├── street
             └── city

        ↓ cria valores

myProfile
otherProfile
profileByPosition
profileByFields
profileWithAddress
completeProfile
```

---

## 🧪 Resumo das formas estudadas

| Conceito | Exemplo | Significado |
| --- | --- | --- |
| Declarar um tipo struct | `type user struct { ... }` | cria um novo tipo estruturado |
| Criar uma variável vazia | `var profile user` | campos recebem valores zero |
| Acessar campo | `profile.name` | acessa um campo da struct |
| Alterar campo | `profile.age = 28` | atribui um novo valor ao campo |
| Literal por posição | `user{"Wesley", 28, address{}}` | valores seguem a ordem dos campos |
| Literal com nomes | `user{name: "Wesley"}` | associa o valor ao nome do campo |
| Omitir campo | `user{name: "Wesley"}` | campos restantes recebem valor zero |
| Struct aninhada | `address address` | um campo possui outro tipo struct |
| Acesso aninhado | `profile.address.city` | percorre a struct externa e a interna |
| Literal aninhado | `address: address{...}` | cria a struct interna dentro da externa |
| Alteração aninhada | `profile.address.city = "Votorantim"` | altera um campo da struct interna |
| Campo embutido | `address` | recurso diferente do campo nomeado |

---

## ⚠️ Erros de sintaxe importantes desta aula

### Tentativa 1

```go
user("Wesley", 28)
```

Forma correta com literal:

```go
user{"Wesley", 28, address{}}
```

ou:

```go
user{
    name: "Wesley",
    age:  28,
}
```

### Tentativa 2

```go
user(name: "Wesley")
```

Forma correta:

```go
user{
    name: "Wesley",
}
```

Uma forma simples de lembrar:

```text
struct literal
     ↓
usa { }
```

---

## 🧠 O que aprendi

Neste experimento foi possível observar que:

* structs agrupam campos relacionados dentro de um único tipo;
* `type` pode ser utilizado para declarar um novo tipo;
* `struct` define um tipo estruturado composto por campos;
* cada campo possui nome e tipo;
* uma variável de tipo struct pode ser declarada com `var`;
* quando uma struct é criada sem valores explícitos, seus campos recebem os valores zero de seus tipos;
* campos são acessados utilizando `.`;
* campos podem receber novos valores através de atribuição;
* várias variáveis podem utilizar o mesmo tipo struct e armazenar valores diferentes;
* `fmt.Printf("%+v", valor)` ajuda a visualizar os nomes dos campos durante os estudos;
* uma struct pode ser criada utilizando um literal;
* literais de struct utilizam chaves `{}`;
* `user("Wesley", 28)` não representa um literal de struct;
* um literal sem nomes associa valores pela posição dos campos;
* quando utilizamos valores por posição, a ordem importa;
* no literal por posição, precisamos fornecer um valor para todos os campos;
* um literal com nomes deixa explícito qual campo recebe cada valor;
* em literais com nomes, campos podem ser omitidos e receberão seus valores zero;
* `user(name: "Wesley")` não é uma sintaxe válida para nomear campos;
* a forma válida é `user{name: "Wesley"}`;
* podemos criar tipos diferentes para representar conceitos diferentes, como `user` e `address`;
* uma struct pode possuir outra struct como um de seus campos;
* `address address` significa campo chamado `address` cujo tipo é `address`;
* structs aninhadas podem ser preenchidas campo a campo;
* campos aninhados podem ser acessados com caminhos como `profile.address.city`;
* uma struct interna também pode ser criada com um literal dentro do literal da struct externa;
* campos de uma struct interna podem ser alterados sem substituir toda a struct externa;
* structs podem possuir vários níveis de aninhamento quando o modelo de dados exigir;
* aninhamento com campo nomeado e embedding são conceitos relacionados, mas diferentes;
* tipos podem ser declarados dentro de uma função, porém declará-los no nível do pacote pode melhorar a organização e permitir reutilização por outras funções do mesmo pacote;
* campos iniciados com letra minúscula não são exportados para outros pacotes;
* campos iniciados com letra maiúscula podem ser exportados;
* o nome de um campo deve comunicar claramente o dado representado, por isso `street` é mais específico que `name` dentro de `address`.

---

> Structs nos permitem transformar informações soltas em tipos que representam conceitos do nosso programa — e, ao aninhar structs, podemos construir modelos maiores a partir de estruturas menores e bem definidas.

<p align="center">
  <img src="../docs/images/footer_06.jfif" alt="Go Soul Society">
</p>
