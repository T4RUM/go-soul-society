# 🧬 07 — Herança em Go: composição e embedding

Diferente de linguagens orientadas a objetos que trabalham com classes, **Go não possui herança de classes**.

Não existe em Go uma palavra-chave como:

```text
extends
inherits
super
```

Também não criamos uma `struct` que seja automaticamente uma subclasse de outra.

Porém, Go possui recursos que permitem **compor tipos a partir de outros tipos**.

Um desses recursos é o **struct embedding**, ou seja, a possibilidade de colocar uma `struct` dentro de outra utilizando um **campo embutido**.

Esse comportamento pode lembrar herança em alguns momentos, principalmente porque campos do tipo embutido podem ser acessados de forma promovida, mas é importante não confundir os conceitos:

```text
Go
 ↓
não possui herança de classes
 ↓
utiliza composição
 ↓
struct embedding é uma das formas de composição
```

Neste experimento, o objetivo é entender:

* por que não dizemos que Go possui herança de classes;
* o que significa composição em Go;
* como uma `struct` pode reutilizar outra `struct`;
* o que é um **embedded field**;
* como declarar uma struct embutida;
* a diferença entre um campo nomeado e um campo embutido;
* como criar uma variável utilizando uma struct embutida;
* por que o tipo embutido continua existindo dentro da struct externa;
* como acessar diretamente o campo embutido;
* como funciona a promoção de campos;
* por que promoção de campos não transforma `student` em `person`;
* como inicializar uma struct embutida utilizando um literal;
* como acessar os campos próprios da struct externa;
* a diferença entre **“é um”** e **“tem um”**;
* por que embedding deve ser entendido como composição, e não como herança tradicional.

---

## 🧪 Estrutura do experimento

Nesta aula, podemos organizar os arquivos da seguinte forma:

```text
go-soul-society/
│
├── 7-Heranca/
│   │
│   ├── 07-heranca.md
│   └── heranca.go
│
├── go.mod
└── README.md
```

O arquivo:

```text
7-Heranca/heranca.go
```

contém o exemplo utilizado para estudar composição e embedding de structs.

---

## 🧠 Primeiro ponto: Go não possui herança de classes

Em algumas linguagens orientadas a objetos, podemos encontrar uma relação semelhante a:

```text
Person
  ↑
  │ herança
  │
Student
```

Nesse modelo, poderíamos dizer:

```text
Student é uma Person
```

porque `Student` herdaria características de `Person`.

Em linguagens com herança de classes, normalmente existe alguma sintaxe específica para declarar essa relação.

Conceitualmente, poderíamos encontrar algo parecido com:

```text
class Student extends Person
```

Go não trabalha dessa forma.

Em Go:

```text
não existem classes
      ↓
não existe extends
      ↓
não existe uma hierarquia de classes baseada em herança
```

Por isso, quando colocamos uma `struct` dentro de outra, não devemos dizer que criamos uma subclasse.

---

## 🧩 Composição em vez de herança

Uma ideia muito importante em Go é construir tipos maiores utilizando tipos menores.

Podemos imaginar:

```text
person
  │
  ├── name
  ├── age
  ├── weight
  └── lasName

        ↓ pode ser utilizada dentro de

student
  │
  ├── person
  └── course
```

A struct `student` **contém** uma `person`.

Esse relacionamento é uma forma de composição.

Uma forma simples de pensar é:

```text
student
   ↓
tem uma person
```

e não:

```text
student
   ↓
é uma person por herança
```

Essa diferença será importante durante toda a aula.

---

## 👤 Criando a struct `person`

No código temos:

```go
type person struct {
    name    string
    age     int
    weight  int
    lasName string
}
```

Essa declaração cria um novo tipo chamado:

```text
person
```

A struct possui quatro campos:

```text
person
  │
  ├── name    → string
  ├── age     → int
  ├── weight  → int
  └── lasName → string
```

Portanto, uma variável do tipo `person` pode armazenar esses quatro valores relacionados.

---

## 🧬 Criando a struct `student`

Depois temos:

```go
type student struct {
    person
    course string
}
```

A struct `student` possui:

```text
student
  │
  ├── person
  └── course → string
```

O detalhe mais importante está nesta linha:

```go
person
```

Observe que não escrevemos um nome de campo separado seguido pelo tipo.

Por exemplo, não temos:

```go
profile person
```

Temos apenas:

```go
person
```

Esse recurso é chamado de **embedded field**, ou campo embutido.

---

## 📦 O que realmente acontece em `person`

É importante entender corretamente a sintaxe:

```go
type student struct {
    person
    course string
}
```

Na linha:

```go
person
```

`person` **é o tipo que está sendo embutido**.

O que não informamos é um nome de campo explícito separado.

Em um campo comum, temos:

```text
nome do campo   tipo
      ↓          ↓
    profile    person
```

Por exemplo:

```go
type student struct {
    profile person
    course  string
}
```

No embedding, escrevemos:

```go
type student struct {
    person
    course string
}
```

Podemos imaginar:

```text
person
  ↓
tipo embutido
  ↓
também origina o nome do campo embutido
```

Por isso, dizer que “não passamos o tipo” não seria a descrição correta.

O tipo está presente.

O que foi omitido é o **nome explícito de um campo separado**.

---

## 🪆 Relembrando a aula anterior

Na aula de structs vimos uma forma de aninhamento com campo nomeado:

```go
type user struct {
    name    string
    address address
}
```

Nesse caso:

```text
address address
│       │
│       └── tipo
└────────── nome do campo
```

E o acesso ocorre assim:

```go
profile.address.city
```

Agora estamos estudando outra forma:

```go
type student struct {
    person
    course string
}
```

Aqui `person` é um **campo embutido**.

Essa é justamente a diferença que havia sido introduzida na aula anterior.

---

## 🔬 Campo nomeado vs campo embutido

Podemos comparar as duas formas.

### Campo nomeado

```go
type student struct {
    profile person
    course  string
}
```

Estrutura:

```text
student
  │
  ├── profile
  │     ↓
  │   person
  │
  └── course
```

Para acessar o nome:

```go
student1.profile.name
```

### Campo embutido

```go
type student struct {
    person
    course string
}
```

Estrutura:

```text
student
  │
  ├── person
  │     ├── name
  │     ├── age
  │     ├── weight
  │     └── lasName
  │
  └── course
```

Podemos acessar explicitamente:

```go
student1.person.name
```

Mas, por causa da promoção de campos, também podemos utilizar:

```go
student1.name
```

em situações em que não existe ambiguidade.

---

## ⬆️ Promoção de campos

Quando uma struct é embutida em outra, os campos acessíveis da struct interna podem ser **promovidos**.

Considere:

```go
type person struct {
    name string
    age  int
}

type student struct {
    person
    course string
}
```

Se tivermos:

```go
student1 := student{
    person: person{
        name: "Jack",
        age:  20,
    },
    course: "Data Science",
}
```

podemos acessar o nome de forma explícita:

```go
student1.person.name
```

Mas também podemos escrever:

```go
student1.name
```

Podemos visualizar:

```text
student1
   │
   ├── person
   │     └── name = "Jack"
   │
   └── course
```

O caminho completo continua existindo:

```text
student1.person.name
```

Porém o embedding permite uma forma promovida:

```text
student1.name
```

Isso é uma conveniência oferecida pela linguagem.

---

## ⚠️ Promoção não é herança

Esse é um dos pontos mais importantes da aula.

O fato de podermos escrever:

```go
student1.name
```

não significa que `student` herdou de `person`.

O campo ainda pertence ao valor `person` embutido dentro de `student`.

Podemos pensar:

```text
student1.name
      ↓
forma promovida de acesso
      ↓
student1.person.name
```

O embedding facilita o acesso.

Ele não cria uma relação de herança de classes.

---

## 🧠 `student` continua sendo um tipo diferente de `person`

Temos dois tipos declarados:

```go
type person struct {
    // ...
}

type student struct {
    person
    course string
}
```

Para o Go:

```text
person
   ≠
student
```

São tipos diferentes.

Por exemplo, isto não é válido:

```go
var p person

student1 := student{}

p = student1
```

A atribuição acima não funciona simplesmente porque `student` contém uma `person`.

O Go não considera `student` automaticamente um subtipo de `person`.

Se quisermos obter a parte `person`, podemos acessar o campo embutido:

```go
p = student1.person
```

Agora temos:

```text
student1
   │
   └── person
          ↓
          p
```

Isso deixa ainda mais claro que estamos trabalhando com composição.

---

## 🔀 “É um” vs “tem um”

Uma maneira útil de diferenciar os conceitos é pensar nas frases:

### Herança tradicional

```text
Student é uma Person
```

Isso representa uma relação de subtipo comum em modelos baseados em herança.

### Composição com embedding

```text
Student tem uma Person embutida
```

No código desta aula, a segunda leitura é a mais adequada.

```text
student
   ↓
contém
   ↓
person
```

Por isso, o embedding não deve ser explicado simplesmente como “herança do Go”.

Ele pode lembrar herança em alguns usos, mas o modelo da linguagem é diferente.

---

## 🧪 Criando `person1`

Dentro de `main`, criamos:

```go
person1 := person{
    name:    "John",
    age:     20,
    weight:  40,
    lasName: "Jim",
}
```

Aqui estamos utilizando um literal da struct:

```text
person
```

com nomes de campos.

Podemos visualizar:

```text
person1
  │
  ├── name    = "John"
  ├── age     = 20
  ├── weight  = 40
  └── lasName = "Jim"
```

Essa variável possui diretamente o tipo:

```text
person
```

---

## 🎓 Criando `person2` como `student`

Depois temos:

```go
person2 := student{
    person: person{
        name:    "Jack",
        age:     20,
        weight:  40,
        lasName: "Jim",
    },
    course: "Harvard Data Science",
}
```

Apesar do nome da variável ser:

```text
person2
```

o tipo dela é:

```text
student
```

Isso acontece porque o valor utilizado na declaração é:

```go
student{ ... }
```

Podemos imaginar:

```text
person2
   ↓
student
   │
   ├── person
   │     ├── name    = "Jack"
   │     ├── age     = 20
   │     ├── weight  = 40
   │     └── lasName = "Jim"
   │
   └── course = "Harvard Data Science"
```

O nome da variável não determina seu tipo.

O valor utilizado na declaração é que permite ao Go inferir o tipo.

---

## 🧬 Por que usamos `person: person{...}`?

A declaração de `student` é:

```go
type student struct {
    person
    course string
}
```

Mesmo que `person` seja um campo embutido, ele ainda pode ser inicializado em um literal utilizando seu nome implícito:

```go
person: person{
    // campos
},
```

Podemos separar:

```text
person: person{...}
│       │
│       └── cria um valor do tipo person
│
└────────── inicializa o campo embutido person
```

Portanto:

```go
student{
    person: person{
        name: "Jack",
    },
}
```

significa aproximadamente:

```text
crie um student
      ↓
preencha seu campo embutido person
      ↓
com um novo valor person
```

---

## ⚠️ Campos promovidos não são chaves diretas do literal externo

Depois que o valor existe, podemos acessar um campo promovido:

```go
person2.name
```

Porém, durante a criação do literal de `student`, não escrevemos:

```go
student{
    name:   "Jack",
    course: "Harvard Data Science",
}
```

apenas porque `name` pode ser promovido.

A forma utilizada no experimento é:

```go
student{
    person: person{
        name: "Jack",
    },
    course: "Harvard Data Science",
}
```

Isso torna explícito que `name` pertence à `person` embutida.

Podemos imaginar:

```text
literal de student
      ↓
inicializa person
      ↓
literal de person
      ↓
inicializa name
```

---

## 🎯 Acessando o campo próprio de `student`

No final do programa temos:

```go
fmt.Println(person2.course)
```

O campo:

```text
course
```

foi declarado diretamente em:

```go
type student struct {
    person
    course string
}
```

Por isso o acesso é:

```go
person2.course
```

Podemos visualizar:

```text
person2
  │
  ├── person
  │
  └── course = "Harvard Data Science"
           ↑
           │
     person2.course
```

---

## 👤 Acessando explicitamente a `person` embutida

Mesmo existindo promoção de campos, a struct embutida continua acessível diretamente.

Podemos fazer:

```go
fmt.Println(person2.person)
```

Ou acessar um campo pelo caminho completo:

```go
fmt.Println(person2.person.name)
```

O caminho é:

```text
person2
   ↓
person
   ↓
name
```

Portanto:

```go
person2.person.name
```

deixa explícito de onde o campo veio.

---

## ⬆️ Acessando um campo promovido

Como `person` foi embutida, também podemos escrever:

```go
fmt.Println(person2.name)
```

Nesse caso, o Go encontra o campo:

```text
name
```

na `person` embutida.

Conceitualmente:

```text
person2.name
     ↓
campo promovido
     ↓
person2.person.name
```

Isso deixa o código mais curto.

Mas novamente:

> acesso promovido não significa herança.

A `person` continua sendo um valor contido dentro de `student`.

---

## 🕳️ Valor zero de uma struct embutida

Também podemos declarar:

```go
var student1 student
```

Sem fornecer valores.

Como `student` é uma struct, todos os seus campos recebem valores zero.

Temos:

```text
student1
  │
  ├── person
  │     ├── name    = ""
  │     ├── age     = 0
  │     ├── weight  = 0
  │     └── lasName = ""
  │
  └── course = ""
```

A `person` embutida também recebe seu valor zero.

Isso segue o mesmo comportamento estudado anteriormente para structs.

---

## 🧱 Embedding ainda é composição

Podemos representar o relacionamento principal da aula assim:

```text
student
  │
  ├── contém → person
  │             │
  │             ├── name
  │             ├── age
  │             ├── weight
  │             └── lasName
  │
  └── course
```

O tipo `student` foi construído utilizando outro tipo.

Isso é composição.

O embedding adiciona algumas conveniências, como a promoção de campos, mas não altera essa natureza.

---

## 🧬 Comparando com uma struct aninhada nomeada

Se tivéssemos:

```go
type student struct {
    profile person
    course  string
}
```

o acesso seria:

```go
student1.profile.name
```

Agora, com:

```go
type student struct {
    person
    course string
}
```

podemos utilizar:

```go
student1.person.name
```

e também:

```go
student1.name
```

Resumo:

| Forma | Declaração | Acesso |
| --- | --- | --- |
| Campo nomeado | `profile person` | `student1.profile.name` |
| Campo embutido | `person` | `student1.person.name` ou `student1.name` |

A principal diferença nesta aula está na promoção oferecida pelo campo embutido.

---

## 🚫 O que embedding não faz

Embedding não transforma automaticamente uma struct em outra.

Considere:

```go
type person struct {
    name string
}

type student struct {
    person
    course string
}
```

Não devemos concluir:

```text
student == person
```

Nem:

```text
student é subtipo automático de person
```

O correto é pensar:

```text
student
   ↓
possui um campo embutido
   ↓
person
```

Por isso, embedding não substitui o conceito de herança de classes.

É outro mecanismo.

---

## 🧭 O que o embedding oferece

Dentro do escopo desta aula, o embedding nos oferece principalmente:

```text
reutilização de uma struct
        ↓
composição de tipos
        ↓
promoção de campos
        ↓
acesso mais conveniente
```

Por exemplo:

```go
person2.name
```

pode chegar ao campo:

```go
person2.person.name
```

Isso reduz a quantidade de caminhos explícitos que precisamos escrever em algumas situações.

---

## 🧠 Por que isso pode parecer herança?

Pode surgir a impressão de herança porque `student` passa a permitir acessos como:

```go
person2.name
person2.age
person2.weight
person2.lasName
```

Mesmo que esses campos tenham sido declarados em:

```go
type person struct {
    // ...
}
```

Visualmente parece que `student` “ganhou” esses campos.

Mas o que realmente aconteceu foi:

```text
person foi embutida
      ↓
seus campos continuam nela
      ↓
alguns deles podem ser promovidos
      ↓
student permite um acesso mais curto
```

Essa distinção evita transportar diretamente para Go o modelo mental de classes e subclasses.

---

## 🧩 Composição favorece tipos menores

O exemplo da aula separa duas responsabilidades.

A struct:

```go
person
```

representa os dados de uma pessoa.

A struct:

```go
student
```

adiciona uma informação específica:

```text
course
```

Podemos visualizar:

```text
dados gerais de pessoa
        ↓
      person

dados específicos de estudante
        ↓
      course

person + course
      ↓
    student
```

Em vez de duplicar todos os campos dentro de `student`, reutilizamos o tipo `person`.

---

## 🔬 Se não utilizássemos composição

Poderíamos declarar:

```go
type student struct {
    name    string
    age     int
    weight  int
    lasName string
    course  string
}
```

Isso funcionaria.

Mas agora os campos de pessoa estariam repetidos.

Teríamos:

```text
person
  ├── name
  ├── age
  ├── weight
  └── lasName

student
  ├── name
  ├── age
  ├── weight
  ├── lasName
  └── course
```

Com composição:

```text
person
  ├── name
  ├── age
  ├── weight
  └── lasName

student
  ├── person
  └── course
```

A estrutura deixa explícito que parte dos dados de `student` vem de um conceito reutilizável: `person`.

---

## 📛 O nome da variável não cria herança

No código temos:

```go
person1 := person{ ... }
person2 := student{ ... }
```

Os nomes:

```text
person1
person2
```

podem dar a impressão de que os dois valores são do mesmo tipo.

Mas não são.

Temos:

```text
person1 → person
person2 → student
```

O tipo vem da expressão utilizada no lado direito:

```go
person{ ... }
student{ ... }
```

Portanto:

```text
nome da variável
      ≠
tipo da variável
```

---

## ✏️ Observação sobre `lasName`

No código da aula aparece o campo:

```go
lasName string
```

Esse nome foi mantido neste material exatamente como aparece no experimento.

Se a intenção for representar “sobrenome” em inglês, um nome mais comum seria:

```go
lastName string
```

Essa alteração seria apenas uma correção de nomenclatura e não muda o conceito de embedding estudado nesta aula.

---

## 🌎 Campos com letra minúscula

Os tipos e campos do experimento começam com letra minúscula:

```go
person
student
name
age
weight
lasName
course
```

Como vimos anteriormente, identificadores iniciados com letra minúscula não são exportados para outros pacotes.

Isso não impede o uso no exemplo porque tudo está no mesmo pacote:

```go
package main
```

Se os tipos precisassem ser utilizados por outros pacotes, poderíamos estudar versões exportadas em outro momento.

---

## 🧪 Código completo do experimento

Arquivo:

```text
7-Heranca/heranca.go
```

```go
package main

import "fmt"

type person struct {
    name    string
    age     int
    weight  int
    lasName string
}

type student struct {
    person
    course string
}

func main() {
    person1 := person{
        name:    "John",
        age:     20,
        weight:  40,
        lasName: "Jim",
    }

    person2 := student{
        person: person{
            name:    "Jack",
            age:     20,
            weight:  40,
            lasName: "Jim",
        },
        course: "Harvard Data Science",
    }

    fmt.Println(person1)
    fmt.Println(person2.course)
}
```

O ponto central do código está em:

```go
type student struct {
    person
    course string
}
```

A linha:

```go
person
```

representa um campo embutido.

---

## ▶️ Executando o experimento

Dentro do diretório da aula, podemos executar:

```bash
go run .
```

Ou, informando o arquivo diretamente:

```bash
go run heranca.go
```

---

## 🖨️ Resultado do experimento

Com o código da aula, a saída será:

```text
{John 20 40 Jim}
Harvard Data Science
```

A primeira linha vem de:

```go
fmt.Println(person1)
```

Como utilizamos `fmt.Println` diretamente com uma struct, os valores aparecem na ordem dos campos:

```text
John
20
40
Jim
```

A segunda linha vem de:

```go
fmt.Println(person2.course)
```

e imprime:

```text
Harvard Data Science
```

---

## 🔍 Observando `person1`

Temos:

```go
fmt.Println(person1)
```

A variável contém:

```text
person1
  │
  ├── name    = "John"
  ├── age     = 20
  ├── weight  = 40
  └── lasName = "Jim"
```

Por isso o `Println` mostra:

```text
{John 20 40 Jim}
```

---

## 🔍 Observando `person2`

A variável:

```go
person2
```

é um `student`.

Sua estrutura pode ser visualizada assim:

```text
person2
  │
  ├── person
  │     ├── name    = "Jack"
  │     ├── age     = 20
  │     ├── weight  = 40
  │     └── lasName = "Jim"
  │
  └── course = "Harvard Data Science"
```

No código original imprimimos apenas:

```go
person2.course
```

Por isso vemos:

```text
Harvard Data Science
```

---

## 🧪 Experimentando a promoção de campos

Para observar o embedding com mais clareza, poderíamos acrescentar temporariamente:

```go
fmt.Println(person2.name)
fmt.Println(person2.age)
fmt.Println(person2.person.name)
```

A saída seria:

```text
Jack
20
Jack
```

Os dois acessos:

```go
person2.name
```

e:

```go
person2.person.name
```

chegam ao mesmo campo neste exemplo.

A diferença é que:

```text
person2.name
```

utiliza promoção.

Enquanto:

```text
person2.person.name
```

explicita o caminho completo.

---

## 🔬 Experimentando o tipo embutido diretamente

Também poderíamos escrever:

```go
fmt.Println(person2.person)
```

Isso mostraria o valor da `person` que está dentro do `student`.

Esse experimento ajuda a perceber:

```text
student
   ↓
não virou person
   ↓
ele contém uma person
```

A `person` pode ser acessada como uma parte concreta do valor `student`.

---

## 🧭 Fluxo geral do programa

Podemos visualizar o experimento desta forma:

```text
type person struct
      │
      ├── name
      ├── age
      ├── weight
      └── lasName
             │
             │ embutida em
             ↓
type student struct
      │
      ├── person
      └── course

             ↓

main()

  ├── person1 := person{...}
  │        ↓
  │     tipo person
  │        ↓
  │     fmt.Println(person1)
  │
  └── person2 := student{...}
           ↓
        tipo student
           │
           ├── person: person{...}
           └── course
                  ↓
        fmt.Println(person2.course)
```

O fluxo reforça que `person` e `student` continuam sendo tipos diferentes.

---

## 🧪 Resumo das formas estudadas

| Conceito | Exemplo | Significado |
| --- | --- | --- |
| Struct base do exemplo | `type person struct { ... }` | representa os dados de uma pessoa |
| Struct composta | `type student struct { ... }` | representa um estudante |
| Campo embutido | `person` | incorpora um valor `person` dentro de `student` sem nome de campo explícito separado |
| Campo próprio | `course string` | pertence diretamente a `student` |
| Inicializar campo embutido | `person: person{...}` | cria e atribui a parte `person` do `student` |
| Acesso explícito | `person2.person.name` | percorre o campo embutido diretamente |
| Campo promovido | `person2.name` | acesso abreviado a um campo da struct embutida |
| Acesso ao campo externo | `person2.course` | acessa um campo declarado diretamente em `student` |
| Composição | `student` contém `person` | constrói um tipo utilizando outro tipo |
| Herança de classes | não existe | Go não utiliza hierarquia de classes com `extends` |

---

## ⚠️ Ideias que não devemos confundir

### 1. Embedding não significa `extends`

Isto:

```go
type student struct {
    person
}
```

não equivale a dizer literalmente:

```text
student extends person
```

Go não possui essa relação de classes.

---

### 2. `student` não vira `person`

Mesmo contendo:

```go
person
```

os tipos continuam diferentes:

```text
student ≠ person
```

---

### 3. O tipo não foi omitido

Na linha:

```go
person
```

o tipo é justamente:

```text
person
```

O que foi omitido é um **nome de campo explícito separado**.

---

### 4. Campo promovido continua pertencendo à struct embutida

Quando escrevemos:

```go
person2.name
```

o campo `name` continua vindo de:

```go
person2.person
```

---

### 5. O nome `person2` não define o tipo da variável

Temos:

```go
person2 := student{ ... }
```

Logo:

```text
person2 → student
```

e não:

```text
person2 → person
```

---

## 🧠 O que aprendi

Neste experimento foi possível observar que:

* Go não possui herança de classes;
* Go não utiliza palavras-chave como `extends` para criar subclasses;
* structs podem ser combinadas utilizando composição;
* uma struct pode ser embutida dentro de outra;
* um campo escrito apenas como `person` dentro de `student` é um campo embutido;
* na declaração `person`, o tipo não foi omitido — `person` é o próprio tipo embutido;
* o que não aparece é um nome de campo explícito separado do tipo;
* um campo embutido também pode ser acessado diretamente pelo nome do tipo, como `person2.person`;
* os campos da struct embutida podem ser promovidos;
* `person2.name` pode acessar o mesmo campo que `person2.person.name` quando não existe ambiguidade;
* promoção de campos é uma conveniência da linguagem e não representa herança de classes;
* `student` e `person` continuam sendo tipos diferentes;
* um valor `student` não pode ser utilizado automaticamente como um valor `person`;
* podemos acessar a parte `person` de um `student` com `student1.person`;
* um literal de `student` inicializa a struct embutida utilizando `person: person{...}`;
* campos promovidos não se tornam chaves diretas do literal da struct externa;
* `course` pertence diretamente à struct `student`;
* o relacionamento do exemplo é melhor entendido como “`student` tem uma `person`”;
* embedding permite reutilizar estruturas sem duplicar todos os seus campos;
* composição e embedding podem produzir uma aparência semelhante à herança em alguns acessos, mas são conceitos diferentes.

---

> Em Go, não herdamos uma classe para criar outra. Construímos tipos por composição, e o struct embedding permite reutilizar e promover campos sem transformar um tipo em subtipo do outro.

<p align="center">
  <img src="../docs/images/footer_07.jfif" alt="Go Soul Society">
</p>
